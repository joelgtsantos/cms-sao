package main

import (
	"fmt"

	"github.com/goadesign/goa"
	"github.com/jossemargt/cms-sao/app"
	"github.com/jossemargt/cms-sao/storage"
)

// EntrySubmitTrxController implements the EntrySubmitTrx resource.
type EntrySubmitTrxController struct {
	*goa.Controller
	repository storage.EntrySubmitTrxRepository
}

// NewEntrySubmitTrxController creates a EntrySubmitTrx controller.
func NewEntrySubmitTrxController(service *goa.Service, store storage.EntrySubmitTrxRepository) *EntrySubmitTrxController {
	return &EntrySubmitTrxController{
		Controller: service.NewController("EntrySubmitTrxController"),
		repository: store,
	}
}

// Get runs the get action.
func (c *EntrySubmitTrxController) Get(ctx *app.GetEntrySubmitTrxContext) error {
	entryTrx, err := c.repository.FindByID(ctx, ctx.TrxID)

	if err != nil {
		return ctx.NotFound()
	}

	res := &app.ComJossemargtSaoEntrySubmitTransactionFull{
		ID:        entryTrx.ID,
		Status:    entryTrx.Status,
		UpdatedAt: &entryTrx.UpdatedAt,
		CreatedAt: &entryTrx.CreatedAt,
		Href:      fmt.Sprintf("%s%s", app.EntrySubmitTrxHref(), entryTrx.ID),
		Links:     new(app.ComJossemargtSaoEntrySubmitTransactionLinks),
	}

	if entryTrx.EntryID != 0 {
		res.Links.Entry = &app.ComJossemargtSaoEntryLink{
			ID:   entryTrx.EntryID,
			Href: fmt.Sprintf("%s%d", app.EntryHref(), entryTrx.EntryID),
		}
	}

	if entryTrx.ResultID != "" {
		res.Links.Result = &app.ComJossemargtSaoResultLink{
			ID:   entryTrx.ResultID,
			Href: fmt.Sprintf("%s%s", app.ResultHref(), entryTrx.ResultID),
		}
	}

	if entryTrx.EntryID == 0 && entryTrx.ResultID == "" {
		res.Links = nil
	}

	return ctx.OKFull(res)
}

// Show runs the show action.
func (c *EntrySubmitTrxController) Show(ctx *app.ShowEntrySubmitTrxContext) error {
	return ctx.NotImplemented()
}
