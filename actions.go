package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/goadesign/goa"
	"github.com/jossemargt/cms-sao/app"
	"github.com/jossemargt/cms-sao/model"
	"github.com/jossemargt/cms-sao/storage"
	"github.com/pkg/errors"
)

// ActionsController implements the actions resource.
type ActionsController struct {
	*goa.Controller
	nesoQueue    storage.QueueWriter
	entryTrxRepo storage.EntrySubmitTrxRepository
	draftTrxRepo storage.EntrySubmitTrxRepository
}

// NewActionsController creates a actions controller.
func NewActionsController(service *goa.Service, entryTrxStore storage.EntrySubmitTrxRepository,
	draftTrxStore storage.EntrySubmitTrxRepository, nesoQueue storage.QueueWriter) *ActionsController {

	return &ActionsController{
		Controller:   service.NewController("ActionsController"),
		nesoQueue:    nesoQueue,
		entryTrxRepo: entryTrxStore,
		draftTrxRepo: draftTrxStore,
	}
}

// SubmitEntry runs the submitEntry action.
func (c *ActionsController) SubmitEntry(ctx *app.SubmitEntryActionsContext) error {
	n := time.Now()

	id, err := c.entryTrxRepo.Save(ctx, &model.EntrySubmitTrx{
		CreatedAt: n,
		UpdatedAt: n,
		Status:    "unprocessed",
	})

	if err != nil {
		return errors.Wrap(err, "Unable to create Entry transaction")
	}

	nesoMsg := &model.NesoMessage{
		Kind: model.NesoMessageEntryKind,
		Auth: model.NesoMessageAuth{
			Cookies: serializeCookies(ctx.Cookies()),
		},
		Transaction: model.NesoMessageTrx{
			ID: id,
		},
		EntryPayload: struct {
			ContestSlug string               `json:"contestSlug"`
			TaskSlug    string               `json:"taskSlug"`
			Token       bool                 `json:"token"`
			Sources     []*model.EntrySource `json:"sources"`
		}{
			ContestSlug: ctx.Payload.ContestSlug,
			TaskSlug:    ctx.Payload.TaskSlug,
			Token:       ctx.Payload.Token,
			Sources:     sourceToModel(ctx.Payload.Sources),
		},
	}

	if err = c.nesoQueue.Write(ctx, nesoMsg); err != nil {
		return errors.Wrap(err, "Unable submit entry in process queue")
	}

	return ctx.CreatedFull(&app.ComJossemargtSaoEntrySubmitTransactionFull{
		ID:        id,
		Status:    "unprocessed",
		UpdatedAt: &n,
		CreatedAt: &n,
		Href:      fmt.Sprintf("%s%s", app.EntrySubmitTrxHref(), id),
	})
}

// SubmitEntryDraft runs the submitEntryDraft action.
func (c *ActionsController) SubmitEntryDraft(ctx *app.SubmitEntryDraftActionsContext) error {
	n := time.Now()

	id, err := c.draftTrxRepo.Save(ctx, &model.EntrySubmitTrx{
		CreatedAt: n,
		UpdatedAt: n,
		Status:    "unprocessed",
	})

	if err != nil {
		return errors.Wrap(err, "Unable to create Draft transaction")
	}

	nesoMsg := &model.NesoMessage{
		Kind: model.NesoMessageDraftKind,
		Auth: model.NesoMessageAuth{
			Cookies: serializeCookies(ctx.Cookies()),
		},
		Transaction: model.NesoMessageTrx{
			ID: id,
		},
		EntryPayload: struct {
			ContestSlug string               `json:"contestSlug"`
			TaskSlug    string               `json:"taskSlug"`
			Token       bool                 `json:"token"`
			Sources     []*model.EntrySource `json:"sources"`
		}{
			ContestSlug: ctx.Payload.ContestSlug,
			TaskSlug:    ctx.Payload.TaskSlug,
			Token:       ctx.Payload.Token,
			Sources:     sourceToModel(ctx.Payload.Sources),
		},
	}

	if err = c.nesoQueue.Write(ctx, nesoMsg); err != nil {
		return errors.Wrap(err, "Unable submit draft in process queue")
	}

	return ctx.CreatedFull(&app.ComJossemargtSaoDraftSubmitTransactionFull{
		ID:        id,
		Status:    "unprocessed",
		UpdatedAt: &n,
		CreatedAt: &n,
		Href:      fmt.Sprintf("%s%s", app.DraftSubmitTrxHref(), id),
	})
}

func serializeCookies(cookies []*http.Cookie) []string {
	cks := make([]string, len(cookies))

	for i, c := range cookies {
		cks[i] = fmt.Sprintf("%s=%s", c.Name, c.Value)
	}

	return cks
}

func sourceToModel(sources []*app.EntrySource) []*model.EntrySource {
	srcs := make([]*model.EntrySource, len(sources))

	for i, src := range sources {
		srcs[i] = &model.EntrySource{
			FileID:   src.Fileid,
			Filename: src.Filename,
			Language: src.Language,
			Content:  src.Content,
		}
	}

	return srcs
}

// SummarizeScore runs the summarizeScore action.
func (c *ActionsController) SummarizeScore(ctx *app.SummarizeScoreActionsContext) error {
	// ActionsController_SummarizeScore: start_implement

	// Put your logic here

	res := app.ComJossemargtSaoScoreSumCollection{}
	return ctx.OK(res)
	// ActionsController_SummarizeScore: end_implement
}
