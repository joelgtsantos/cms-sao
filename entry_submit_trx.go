package main

import (
	"github.com/goadesign/goa"
	"github.com/jossemargt/cms-sao/app"
)

// EntrySubmitTrxController implements the EntrySubmitTrx resource.
type EntrySubmitTrxController struct {
	*goa.Controller
}

// NewEntrySubmitTrxController creates a EntrySubmitTrx controller.
func NewEntrySubmitTrxController(service *goa.Service) *EntrySubmitTrxController {
	return &EntrySubmitTrxController{Controller: service.NewController("EntrySubmitTrxController")}
}

// Get runs the get action.
func (c *EntrySubmitTrxController) Get(ctx *app.GetEntrySubmitTrxContext) error {
	// EntrySubmitTrxController_Get: start_implement

	// Put your logic here

	res := &app.ComJossemargtSaoEntrySubmitTransactionFull{}
	return ctx.OKFull(res)
	// EntrySubmitTrxController_Get: end_implement
}
