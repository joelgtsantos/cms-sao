package main

import (
	"github.com/goadesign/goa"
	"github.com/jossemargt/cms-sao/app"
)

// ResultController implements the result resource.
type ResultController struct {
	*goa.Controller
}

// NewResultController creates a result controller.
func NewResultController(service *goa.Service) *ResultController {
	return &ResultController{Controller: service.NewController("ResultController")}
}

// Get runs the get action.
func (c *ResultController) Get(ctx *app.GetResultContext) error {
	// ResultController_Get: start_implement

	// Put your logic here

	res := &app.ComJossemargtSaoResultFull{}
	return ctx.OKFull(res)
	// ResultController_Get: end_implement
}

// Show runs the show action.
func (c *ResultController) Show(ctx *app.ShowResultContext) error {
	// ResultController_Show: start_implement

	// Put your logic here

	res := app.ComJossemargtSaoResultCollection{}
	return ctx.OK(res)
	// ResultController_Show: end_implement
}
