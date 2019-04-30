package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/goadesign/goa"
	"github.com/jossemargt/cms-sao/app"
	"github.com/jossemargt/cms-sao/model"
	"github.com/jossemargt/cms-sao/storage"
	"github.com/pkg/errors"
)

// DraftResultController implements the DraftResult resource.
type DraftResultController struct {
	*goa.Controller
	resultRepo storage.DraftResultRepository
}

// NewDraftResultController creates a Draftresult controller.
func NewDraftResultController(service *goa.Service, repository storage.DraftResultRepository) *DraftResultController {
	return &DraftResultController{
		Controller: service.NewController("DraftResultController"),
		resultRepo: repository,
	}
}

// Get runs the get action.
func (c *DraftResultController) Get(ctx *app.GetDraftresultContext) error {
	rawIDs := strings.Split(ctx.ResultID, "-")

	if len(rawIDs) != 2 {
		return ctx.BadRequest(fmt.Errorf("wrong Result ID"))
	}

	IDs := make([]int, 0, 2)

	for _, idString := range rawIDs {
		id, err := strconv.ParseInt(idString, 10, 0)
		if err != nil {
			return ctx.BadRequest(fmt.Errorf("wrong Result ID"))
		}
		IDs = append(IDs, int(id))
	}

	result, err := c.resultRepo.FindByID(IDs[0], IDs[1])
	if err != nil {
		if errors.Cause(sql.ErrNoRows) != nil {
			return ctx.NotFound()
		}

		return errors.Wrap(err, "Un-expected error")
	}

	res := draftresultModelToMediaFull(result)
	return ctx.OKFull(res)
}

// Show runs the show action.
func (c *DraftResultController) Show(ctx *app.ShowDraftresultContext) error {
	dto := storage.DraftResultDTO{
		ContestID:   ctx.Contest,
		ContestSlug: ctx.ContestSlug,
		TaskID:      ctx.Task,
		TaskSlug:    ctx.TaskSlug,
		UserID:      ctx.User,
		DraftID:     ctx.Entry,
		DTO: storage.DTO{
			Page:     ctx.Page,
			PageSize: ctx.PageSize,
			Order:    strings.ToUpper(ctx.Sort),
		},
	}

	results, err := c.resultRepo.FindBy(dto)

	res := app.ComJossemargtSaoDraftResultCollection{}
	if err != nil {
		if errors.Cause(sql.ErrNoRows) != nil {
			return ctx.OK(res)
		}

		return errors.Wrap(err, "Un-expected error")
	}

	for _, result := range results {
		res = append(res, draftresultModelToMedia(&result))
	}

	return ctx.OK(res)
}

func draftresultModelToMediaFull(result *model.DraftResult) *app.ComJossemargtSaoDraftResultFull {
	id := fmt.Sprintf("%d-%d", result.EntryID, result.DatasetID)
	media := app.ComJossemargtSaoDraftResultFull{
		ID:   id,
		Href: fmt.Sprintf("%s%s", app.DraftresultHref(), id),
		Evaluation: &app.EvaluationResult{
			Status: "ok",
			Tries:  result.Evaluation.Tries,
		},
		Execution: &app.ExecutionResult{
			Time:          float64(result.Execution.Time),
			Memory:        result.Execution.Memory,
			WallClockTime: float64(result.Execution.WallClockTime),
			Output:        string(result.Output),
		},
		Compilation: &app.CompilationResult{
			Status:        result.Compilation.Status,
			Tries:         result.Compilation.Tries,
			Stdout:        result.Compilation.Stdout,
			Stderr:        result.Compilation.Stderr,
			Time:          float64(result.Compilation.Time),
			WallClockTime: float64(result.Compilation.WallClockTime),
			Memory:        result.Compilation.Memory,
		},
	}

	if !result.Evaluation.Done {
		media.Evaluation.Status = "unprocessed"
	}

	return &media
}

func draftresultModelToMedia(result *model.DraftResult) *app.ComJossemargtSaoDraftResult {
	id := fmt.Sprintf("%d-%d", result.EntryID, result.DatasetID)

	media := app.ComJossemargtSaoDraftResult{
		ID:   id,
		Href: fmt.Sprintf("%s%s", app.DraftresultHref(), id),
		Execution: &app.ExecutionResult{
			Time:          float64(result.Execution.Time),
			Memory:        result.Execution.Memory,
			WallClockTime: float64(result.Execution.WallClockTime),
		},
	}

	return &media
}
