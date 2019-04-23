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

// ResultController implements the result resource.
type ResultController struct {
	*goa.Controller
	resultRepo storage.ResultRepository
}

// NewResultController creates a result controller.
func NewResultController(service *goa.Service, repository storage.ResultRepository) *ResultController {
	return &ResultController{
		Controller: service.NewController("EntryController"),
		resultRepo: repository,
	}
}

// Get a single Result that corresponds to the given ID as query parameter
func (c *ResultController) Get(ctx *app.GetResultContext) error {
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

	res := resultModelToMediaFull(result)
	return ctx.OKFull(res)
}

// Show runs the show action.
func (c *ResultController) Show(ctx *app.ShowResultContext) error {
	dto := storage.ResultDTO{
		ContestID:   ctx.Contest,
		ContestSlug: ctx.ContestSlug,
		TaskID:      ctx.Task,
		TaskSlug:    ctx.TaskSlug,
		UserID:      ctx.User,
		EntryID:     ctx.Entry,
		MaxScore:    ctx.Max,
		DTO: storage.DTO{
			Page:     ctx.Page,
			PageSize: ctx.PageSize,
			Order:    strings.ToUpper(ctx.Sort),
		},
	}

	results, err := c.resultRepo.FindBy(dto)

	res := app.ComJossemargtSaoResultCollection{}
	if err != nil {
		if errors.Cause(sql.ErrNoRows) != nil {
			return ctx.OK(res)
		}

		return errors.Wrap(err, "Un-expected error")
	}

	for _, result := range results {
		res = append(res, resultModelToMedia(&result, ctx.View))
	}

	return ctx.OK(res)
}

func resultModelToMedia(result *model.Result, view string) *app.ComJossemargtSaoResult {
	id := fmt.Sprintf("%d-%d", result.EntryID, result.DatasetID)

	media := app.ComJossemargtSaoResult{
		ID:   id,
		Href: fmt.Sprintf("%s%s", app.ResultHref(), id),
		Score: &app.ScoreResult{
			ContestValue: float64(result.Scoring.ContestScore),
			TaskValue:    float64(result.Scoring.TaskScore),
		},
		Evaluation: &app.EvaluationResult{
			Tries: result.Evaluation.Tries,
		},
	}

	if result.Evaluation.Done {
		media.Evaluation.Status = "ok"
	} else {
		media.Evaluation.Status = "unprocessed"
	}

	if view == "score" {
		media.Evaluation = new(app.EvaluationResult)
	}

	return &media
}

func resultModelToMediaFull(result *model.Result) *app.ComJossemargtSaoResultFull {
	id := fmt.Sprintf("%d-%d", result.EntryID, result.DatasetID)
	media := app.ComJossemargtSaoResultFull{
		ID:   id,
		Href: fmt.Sprintf("%s%s", app.ResultHref(), id),
		Evaluation: &app.EvaluationResult{
			Tries: result.Evaluation.Tries,
		},
		Score: &app.ScoreResult{
			ContestValue: float64(result.Scoring.ContestScore),
			TaskValue:    float64(result.Scoring.TaskScore),
		},
	}

	if result.Compilation.Status != nil {
		media.Compilation = &app.CompilationResult{
			Status:        *result.Compilation.Status,
			Tries:         result.Compilation.Tries,
			Stdout:        result.Compilation.Stdout,
			Stderr:        result.Compilation.Stderr,
			Time:          float64(result.Compilation.Time),
			WallClockTime: float64(result.Compilation.WallClockTime),
			Memory:        result.Compilation.Memory,
		}
	}

	if result.Evaluation.Done {
		media.Evaluation.Status = "ok"
	} else {
		media.Evaluation.Status = "unprocessed"
	}

	return &media
}
