//go:generate go run gen/main.go

package main

import (
	"context"
	"fmt"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	goup "github.com/ufoscout/go-up"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/jossemargt/cms-sao/app"
	"github.com/jossemargt/cms-sao/storage"
)

func main() {
	appCtx, appCtxCancel := context.WithCancel(context.Background())
	defer appCtxCancel()

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

	// Create service
	service := goa.New("SAO v1")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(up.GetBoolOrDefault("server.log.request", false)))
	service.Use(middleware.ErrorHandler(service, up.GetBoolOrDefault("server.response.tracedump", false)))
	service.Use(middleware.Recover())

	var dbConn *sqlx.DB
	{
		dbhost := up.GetStringOrDefault("cms.datasource.host", "localhost")
		dbport := up.GetStringOrDefault("cms.datasource.port", "5432")
		dbname := up.GetStringOrDefault("cms.datasource.name", "cmsdb")
		dbuser := up.GetStringOrDefault("cms.datasource.username", "cmsuser")
		dbpassword := up.GetStringOrDefault("cms.datasource.password", "")
		dbsslmode := up.GetStringOrDefault("cms.datasource.sslmode", "require")

		var err error
		dbConn, err = sqlx.Connect("postgres",
			fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
				dbuser,
				dbpassword,
				dbname,
				dbhost,
				dbport,
				dbsslmode,
			),
		)

		if err != nil {
			service.LogError("startup", "err", err)
			return
		}
	}

	var mongoDB *mongo.Database
	{
		dbhost := up.GetStringOrDefault("documentsource.host", "localhost")
		dbport := up.GetIntOrDefault("documentsource.port", 27017)
		dbname := up.GetStringOrDefault("documentsource.name", "cmsdb")
		dbuser := up.GetStringOrDefault("documentsource.username", "cmsuser")
		dbpassword := up.GetStringOrDefault("documentsource.password", "")

		mongoClient, err := mongo.Connect(appCtx,
			options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d/%s",
				dbuser,
				dbpassword,
				dbhost,
				dbport,
				dbname,
			)),
		)

		if err != nil {
			service.LogError("startup", "err", err)
			return
		}

		if err = mongoClient.Ping(appCtx, nil); err != nil {
			service.LogError("startup", "err", err)
			return
		}

		mongoDB = mongoClient.Database(dbname)
	}

	// Create resource repositories
	entryRepository := storage.NewEntryRepository(dbConn)
	resultRepository := storage.NewResultRepository(dbConn)
	draftRepository := storage.NewEntryDraftRepository(dbConn)
	draftresultRepository := storage.NewDraftResultRepository(dbConn)
	entrySubmitTrxRepository := storage.NewEntrySubmitTrxRepository(mongoDB)
	draftSubmitTrxRepository := storage.NewDraftSubmitTrxRepository(mongoDB)
	nesoQueue := storage.NewQueueWriter(mongoDB)

	// Mount "actions" controller
	c := NewActionsController(service, entrySubmitTrxRepository, draftSubmitTrxRepository, nesoQueue)
	app.MountActionsController(service, c)
	// Mount "draft" controller
	c2 := NewDraftController(service, draftRepository)
	app.MountDraftController(service, c2)
	// Mount "draft-result" controller
	c3 := NewDraftResultController(service, draftresultRepository)
	app.MountDraftresultController(service, c3)
	// Mount "entry" controller
	c4 := NewEntryController(service, entryRepository)
	app.MountEntryController(service, c4)
	// Mount "result" controller
	c5 := NewResultController(service, resultRepository)
	app.MountResultController(service, c5)
	// Mount "submit entry transaction" controller
	c6 := NewEntrySubmitTrxController(service, entrySubmitTrxRepository)
	app.MountEntrySubmitTrxController(service, c6)
	// Mount "submit entry draft transaction" controller
	c7 := NewDraftSubmitTrxController(service, draftSubmitTrxRepository)
	app.MountDraftSubmitTrxController(service, c7)

	// Start service
	serverAddr := fmt.Sprintf("%s:%s", "", up.GetStringOrDefault("server.port", "8000"))
	if err := service.ListenAndServe(serverAddr); err != nil {
		service.LogError("startup", "err", err)
	}
}
