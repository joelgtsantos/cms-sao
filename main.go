//go:generate goagen bootstrap -d github.com/jossemargt/cms-sao/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/jossemargt/cms-sao/app"
	"github.com/ufoscout/go-up"
	"fmt"
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

	// Create service
	service := goa.New("SAO v1")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(up.GetBoolOrDefault("server.log.request", false)))
	service.Use(middleware.ErrorHandler(service, up.GetBoolOrDefault("server.response.tracedump", false)))
	service.Use(middleware.Recover())

	// Mount "entry" controller
	c := NewEntryController(service)
	app.MountEntryController(service, c)
	// Mount "result" controller
	c2 := NewResultController(service)
	app.MountResultController(service, c2)
	// Mount "scores" controller
	c3 := NewScoresController(service)
	app.MountScoresController(service, c3)

	// Start service
	serverAddr := fmt.Sprintf("%s:%s","", up.GetStringOrDefault("server.port", "8080"))
	if err := service.ListenAndServe(serverAddr); err != nil {
		service.LogError("startup", "err", err)
	}

}
