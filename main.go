//go:generate goagen bootstrap -d github.com/jossemargt/cms-sao/design

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	goup "github.com/ufoscout/go-up"

	"github.com/jossemargt/cms-sao/app"
	"github.com/jossemargt/cms-sao/storage"
)

func main() {
	// Resolve configurations
	var up goup.GoUp
	{
		up, _ = goup.NewGoUp().
			// Read default config file if any
			AddFile("./config.properties", true).
			// Read environment variables with prefix SAO_
			AddReader(goup.NewEnvReader("SAO_", true, true)).
			Build()
	}

	var httpClient *httpRetryableClient
	{
		httpClient = new(httpRetryableClient)
		retryable := retryablehttp.NewClient()

		// Max retries
		retryable.RetryMax = up.GetIntOrDefault("cms.http.retry.limit", 5)

		// Minimal retry wait (in seconds)
		minWait := up.GetIntOrDefault("cms.http.retry.waitmin", 1)
		retryable.RetryWaitMin = time.Duration(minWait) * time.Second

		// Max retry wait
		maxWait := up.GetIntOrDefault("cms.http.retry.waitmax", 30)
		retryable.RetryWaitMax = time.Duration(maxWait) * time.Second

		httpClient.retryable = retryable
	}

	var dbConn *sqlx.DB
	{
		dbhost := up.GetStringOrDefault("cms.datasource.host", "localhost")
		dbport := up.GetStringOrDefault("cms.datasource.port", "5432")
		dbname := up.GetStringOrDefault("cms.datasource.name", "cmsdb")
		dbuser := up.GetStringOrDefault("cms.datasource.username", "cmsuser")
		dbpassword := up.GetStringOrDefault("cms.datasource.password", "")
		dbsslmode := up.GetStringOrDefault("cms.datasource.sslmode", "require")

		dbConn = sqlx.MustConnect("postgres",
			fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
				dbuser,
				dbpassword,
				dbname,
				dbhost,
				dbport,
				dbsslmode,
			),
		)
	}

	// Create resource repositories
	entryRepository := storage.NewEntryRepository(dbConn)

	// Create service
	service := goa.New("SAO v1")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(up.GetBoolOrDefault("server.log.request", false)))
	service.Use(middleware.ErrorHandler(service, up.GetBoolOrDefault("server.response.tracedump", false)))
	service.Use(middleware.Recover())

	// Mount "actions" controller
	c := NewActionsController(service)
	app.MountActionsController(service, c)
	// Mount "draft" controller
	c2 := NewDraftController(service)
	app.MountDraftController(service, c2)
	// Mount "draft-result" controller
	c3 := NewDraftResultController(service)
	app.MountDraftresultController(service, c3)
	// Mount "entry" controller
	c4 := NewEntryController(service, entryRepository)
	app.MountEntryController(service, c4)
	// Mount "result" controller
	c5 := NewResultController(service)
	app.MountResultController(service, c5)
	// Mount "scores" controller
	c6 := NewScoresController(service)
	app.MountScoresController(service, c6)

	// Start service
	serverAddr := fmt.Sprintf("%s:%s", "", up.GetStringOrDefault("server.port", "8080"))
	if err := service.ListenAndServe(serverAddr); err != nil {
		service.LogError("startup", "err", err)
	}

}

type httpRetryableClient struct {
	retryable *retryablehttp.Client
}

func (c *httpRetryableClient) Do(req *http.Request) (*http.Response, error) {
	request, err := retryablehttp.NewRequest(req.Method, req.URL.String(), req.Body)

	if err != nil {
		return nil, err
	}

	return c.retryable.Do(request)
}
