package main

import (
	"github.com/goadesign/goa"
	"github.com/jossemargt/cms-sao/app"
)

// DraftSubmitTrxController implements the DraftSubmitTrx resource.
type DraftSubmitTrxController struct {
	*goa.Controller
}

// NewDraftSubmitTrxController creates a DraftSubmitTrx controller.
func NewDraftSubmitTrxController(service *goa.Service) *DraftSubmitTrxController {
	return &DraftSubmitTrxController{Controller: service.NewController("DraftSubmitTrxController")}
}

// Get runs the get action.
func (c *DraftSubmitTrxController) Get(ctx *app.GetDraftSubmitTrxContext) error {
	// DraftSubmitTrxController_Get: start_implement

	// Put your logic here

	res := &app.ComJossemargtSaoDraftSubmitTransactionFull{}
	return ctx.OKFull(res)
	// DraftSubmitTrxController_Get: end_implement
}
