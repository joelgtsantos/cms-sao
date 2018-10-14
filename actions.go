package main

import (
	"github.com/goadesign/goa"
	"github.com/jossemargt/cms-sao/app"
)

// ActionsController implements the actions resource.
type ActionsController struct {
	*goa.Controller
}

// NewActionsController creates a actions controller.
func NewActionsController(service *goa.Service) *ActionsController {
	return &ActionsController{Controller: service.NewController("ActionsController")}
}

// SubmitEntry runs the submitEntry action.
func (c *ActionsController) SubmitEntry(ctx *app.SubmitEntryActionsContext) error {
	// ActionsController_SubmitEntry: start_implement

	// Put your logic here

	return nil
	// ActionsController_SubmitEntry: end_implement
}

// SummarizeScore runs the summarizeScore action.
func (c *ActionsController) SummarizeScore(ctx *app.SummarizeScoreActionsContext) error {
	// ActionsController_SummarizeScore: start_implement

	// Put your logic here

	res := app.ComJossemargtSaoScoresumCollection{}
	return ctx.OK(res)
	// ActionsController_SummarizeScore: end_implement
}
