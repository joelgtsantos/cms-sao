package main

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/goadesign/goa"
	"github.com/pkg/errors"

	"github.com/jossemargt/cms-sao/app"
	"github.com/jossemargt/cms-sao/model"
	"github.com/jossemargt/cms-sao/storage"
)

// EntryController implements the Entry resource.
type EntryController struct {
	*goa.Controller
	entryRepo storage.EntryRepository
}

// NewEntryController creates an Entry controller.
func NewEntryController(service *goa.Service, repository storage.EntryRepository) *EntryController {
	return &EntryController{
		Controller: service.NewController("EntryController"),
		entryRepo:  repository,
	}
}

// Get a single Entry that corresponds to the given ID as query parameter
func (c *EntryController) Get(ctx *app.GetEntryContext) error {
	entry, err := c.entryRepo.FindByID(ctx.EntryID)
	if err != nil {
		if errors.Cause(sql.ErrNoRows) != nil {
			return ctx.NotFound()
		}

		return errors.Wrap(err, "Un-expected error")
	}

	res := entryModelToMediaFull(entry)
	return ctx.OKFull(res)
}

// Show a set of Entries that comply with the given query parameters
func (c *EntryController) Show(ctx *app.ShowEntryContext) error {
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

	res := app.ComJossemargtSaoEntryCollection{}
	if err != nil {
		if errors.Cause(sql.ErrNoRows) != nil {
			return ctx.OK(res)
		}

		return errors.Wrap(err, "Un-expected error")
	}

	for _, entry := range entries {
		res = append(res, entryModelToMedia(&entry))
	}

	return ctx.OK(res)
}

func entryModelToMedia(entry *model.Entry) *app.ComJossemargtSaoEntry {
	media := app.ComJossemargtSaoEntry{
		ID:          entry.ID,
		ContestSlug: entry.ContestSlug,
		TaskSlug:    entry.TaskSlug,
		Ranked:      true,
		Href:        fmt.Sprintf("%s%d", app.EntryHref(), entry.ID),
	}

	return &media
}

func entryModelToMediaFull(entry *model.Entry) *app.ComJossemargtSaoEntryFull {
	media := app.ComJossemargtSaoEntryFull{
		ID:          entry.ID,
		ContestID:   entry.ContestID,
		ContestSlug: entry.ContestSlug,
		TaskID:      entry.TaskID,
		TaskSlug:    entry.TaskSlug,
		Ranked:      true,
		Href:        fmt.Sprintf("%s%d", app.EntryHref(), entry.ID),
	}

	links := new(app.ComJossemargtSaoEntryLinks)

	if entry.DatasetID != 0 {
		links.Result = &app.ComJossemargtSaoResultLink{
			ID:   entry.ResultID(),
			Href: fmt.Sprintf("%s%s", app.ResultHref(), entry.ResultID()),
		}
	}

	if entry.DatasetID != 0 {
		links.Score = &app.ComJossemargtSaoScoreLink{
			ID:   entry.ResultID(),
			Href: fmt.Sprintf("%s%s", app.ScoresHref(), entry.ResultID()),
		}
	}

	if links.Result != nil || links.Score != nil {
		media.Links = links
	}

	return &media
}
