package main

import (
	"github.com/goadesign/goa"
	"github.com/jossemargt/cms-sao/app"
)

// DraftresultController implements the draftresult resource.
type DraftresultController struct {
	*goa.Controller
}

// NewDraftresultController creates a draftresult controller.
func NewDraftresultController(service *goa.Service) *DraftresultController {
	return &DraftresultController{Controller: service.NewController("DraftresultController")}
}

// Get runs the get action.
func (c *DraftresultController) Get(ctx *app.GetDraftresultContext) error {
	// DraftresultController_Get: start_implement

	// Put your logic here

	res := &app.ComJossemargtSaoResultFull{}
	return ctx.OKFull(res)
	// DraftresultController_Get: end_implement
}

// Show runs the show action.
func (c *DraftresultController) Show(ctx *app.ShowDraftresultContext) error {
	// DraftresultController_Show: start_implement

	// Put your logic here

	res := app.ComJossemargtSaoResultCollection{}
	return ctx.OK(res)
	// DraftresultController_Show: end_implement
}
