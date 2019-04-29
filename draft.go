package main

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/goadesign/goa"
	"github.com/jossemargt/cms-sao/app"
	"github.com/jossemargt/cms-sao/model"
	"github.com/jossemargt/cms-sao/storage"
	"github.com/pkg/errors"
)

// DraftController implements the Entry Draft resource.
type DraftController struct {
	*goa.Controller
	entryRepo storage.EntryRepository
}

// NewDraftController creates an Entry Draft controller.
func NewDraftController(service *goa.Service, repository storage.EntryRepository) *DraftController {
	return &DraftController{
		Controller: service.NewController("DraftController"),
		entryRepo:  repository,
	}
}

// Get a single Entry Draft that corresponds to the given ID as query parameter
func (c *DraftController) Get(ctx *app.GetDraftContext) error {
	entry, err := c.entryRepo.FindByID(ctx.DraftID)
	if err != nil {
		if errors.Cause(sql.ErrNoRows) != nil {
			return ctx.NotFound()
		}

		return errors.Wrap(err, "Un-expected error")
	}

	res := entryToDraftFullMedia(entry)
	return ctx.OKFull(res)
}

// Show a set of Entry Drafts that comply with the given query parameters
func (c *DraftController) Show(ctx *app.ShowDraftContext) error {
	dto := storage.EntryDTO{
		ContestID:   ctx.Contest,
		ContestSlug: ctx.ContestSlug,
		TaskID:      ctx.Task,
		TaskSlug:    ctx.TaskSlug,
		DTO: storage.DTO{
			Page:     ctx.Page,
			PageSize: ctx.PageSize,
			Order:    strings.ToUpper(ctx.Sort),
		},
	}

	entries, err := c.entryRepo.FindBy(dto)

	res := app.ComJossemargtSaoDraftCollection{}
	if err != nil {
		if errors.Cause(sql.ErrNoRows) != nil {
			return ctx.OK(res)
		}

		return errors.Wrap(err, "Un-expected error")
	}

	for _, entry := range entries {
		res = append(res, entryModelToDraftMedia(&entry))
	}

	return ctx.OK(res)
}

func entryToDraftFullMedia(entry *model.Entry) *app.ComJossemargtSaoDraftFull {
	media := app.ComJossemargtSaoDraftFull{
		ID:          entry.ID,
		ContestID:   entry.ContestID,
		ContestSlug: entry.ContestSlug,
		TaskID:      entry.TaskID,
		TaskSlug:    entry.TaskSlug,
		Language:    entry.Language,
		Href:        fmt.Sprintf("%s%d", app.DraftHref(), entry.ID),
	}

	links := new(app.ComJossemargtSaoDraftLinks)
	if entry.DatasetID != 0 {
		links.Result = &app.ComJossemargtSaoDraftResultLink{
			ID:   entry.ResultID(),
			Href: fmt.Sprintf("%s%s", app.DraftresultHref(), entry.ResultID()),
		}

		media.Links = links
	}

	return &media
}

func entryModelToDraftMedia(entry *model.Entry) *app.ComJossemargtSaoDraft {
	media := app.ComJossemargtSaoDraft{
		ID:          entry.ID,
		ContestSlug: entry.ContestSlug,
		TaskSlug:    entry.TaskSlug,
		Href:        fmt.Sprintf("%s%d", app.DraftHref(), entry.ID),
	}

	return &media
}
