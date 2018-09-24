//go:generate goagen bootstrap -d github.com/jossemargt/cms-sao/design

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/jossemargt/cms-sao/app"
	"github.com/ufoscout/go-up"
)

func main() {
	// Resolve configurations
	var up go_up.GoUp
	{
		up, _ = go_up.NewGoUp().
			// Read default config file if any
			AddFile("./config.properties", true).
			// Read environment variables with prefix SAO_
			AddReader(go_up.NewEnvReader("SAO_", true, true)).
			Build();
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

	// cmsURL := up.GetStringOrDefault("cms.http.url", "http://localhost:8080/")

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
	// Mount "entry" controller
	c2 := NewEntryController(service)
	app.MountEntryController(service, c2)
	// Mount "result" controller
	c3 := NewResultController(service)
	app.MountResultController(service, c3)
	// Mount "scores" controller
	c4 := NewScoresController(service)
	app.MountScoresController(service, c4)

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
