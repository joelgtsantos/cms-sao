//go:generate goagen bootstrap -d github.com/jossemargt/cms-sao/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/jossemargt/cms-sao/app"
)

func main() {
	// Create service
	service := goa.New("SAO v1")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
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
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
