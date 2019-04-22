package main

import (
	"github.com/goadesign/goa"
	"github.com/jossemargt/cms-sao/app"
)

// DraftResultController implements the draftresult resource.
type DraftResultController struct {
	*goa.Controller
}

// NewDraftResultController creates a draftresult controller.
func NewDraftResultController(service *goa.Service) *DraftResultController {
	return &DraftResultController{Controller: service.NewController("DraftResultController")}
}

// Get runs the get action.
func (c *DraftResultController) Get(ctx *app.GetDraftresultContext) error {
	// DraftresultController_Get: start_implement

	// Put your logic here

	res := &app.ComJossemargtSaoDraftResultFull{}
	return ctx.OKFull(res)
	// DraftresultController_Get: end_implement
}

// Show runs the show action.
func (c *DraftResultController) Show(ctx *app.ShowDraftresultContext) error {
	// DraftresultController_Show: start_implement

	// Put your logic here

	res := app.ComJossemargtSaoDraftResultCollection{}
	return ctx.OK(res)
	// DraftresultController_Show: end_implement
}
