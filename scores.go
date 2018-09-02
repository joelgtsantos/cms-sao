package main

import (
	"github.com/goadesign/goa"
	"github.com/jossemargt/cms-sao/app"
)

// ScoresController implements the scores resource.
type ScoresController struct {
	*goa.Controller
}

// NewScoresController creates a scores controller.
func NewScoresController(service *goa.Service) *ScoresController {
	return &ScoresController{Controller: service.NewController("ScoresController")}
}

// Get runs the get action.
func (c *ScoresController) Get(ctx *app.GetScoresContext) error {
	// ScoresController_Get: start_implement

	// Put your logic here

	res := &app.ComJossemargtSaoScoreFull{}
	return ctx.OKFull(res)
	// ScoresController_Get: end_implement
}

// Show runs the show action.
func (c *ScoresController) Show(ctx *app.ShowScoresContext) error {
	// ScoresController_Show: start_implement

	// Put your logic here

	res := app.ComJossemargtSaoScoreCollection{}
	return ctx.OK(res)
	// ScoresController_Show: end_implement
}

// Summarize runs the summarize action.
func (c *ScoresController) Summarize(ctx *app.SummarizeScoresContext) error {
	// ScoresController_Summarize: start_implement

	// Put your logic here

	res := app.ComJossemargtSaoScoresumCollection{}
	return ctx.OK(res)
	// ScoresController_Summarize: end_implement
}
