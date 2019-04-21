package main

import (
	"github.com/goadesign/goa"
	"github.com/jossemargt/cms-sao/app"
)

// DraftController implements the draft resource.
type DraftController struct {
	*goa.Controller
}

// NewDraftController creates a draft controller.
func NewDraftController(service *goa.Service) *DraftController {
	return &DraftController{Controller: service.NewController("DraftController")}
}

// Get runs the get action.
func (c *DraftController) Get(ctx *app.GetDraftContext) error {
	// DraftController_Get: start_implement

	// Put your logic here

	res := &app.ComJossemargtSaoEntryFull{}
	return ctx.OKFull(res)
	// DraftController_Get: end_implement
}

// Show runs the show action.
func (c *DraftController) Show(ctx *app.ShowDraftContext) error {
	// DraftController_Show: start_implement

	// Put your logic here

	res := app.ComJossemargtSaoEntryCollection{}
	return ctx.OK(res)
	// DraftController_Show: end_implement
}
