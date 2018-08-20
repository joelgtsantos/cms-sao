package main

import (
	"github.com/goadesign/goa"
	"github.com/jossemargt/cms-sao/app"
)

// EntryController implements the entry resource.
type EntryController struct {
	*goa.Controller
}

// NewEntryController creates a entry controller.
func NewEntryController(service *goa.Service) *EntryController {
	return &EntryController{Controller: service.NewController("EntryController")}
}

// Create runs the create action.
func (c *EntryController) Create(ctx *app.CreateEntryContext) error {
	// EntryController_Create: start_implement

	// Put your logic here

	return nil
	// EntryController_Create: end_implement
}

// Get runs the get action.
func (c *EntryController) Get(ctx *app.GetEntryContext) error {
	// EntryController_Get: start_implement

	// Put your logic here

	res := &app.ComJossemargtSaoEntryFull{}
	return ctx.OKFull(res)
	// EntryController_Get: end_implement
}

// Show runs the show action.
func (c *EntryController) Show(ctx *app.ShowEntryContext) error {
	// EntryController_Show: start_implement

	// Put your logic here

	res := app.ComJossemargtSaoEntryCollection{}
	return ctx.OK(res)
	// EntryController_Show: end_implement
}
