package main

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"testing"

	"github.com/goadesign/goa"
	goatesthelper "github.com/jossemargt/cms-sao/app/test"
	"github.com/pkg/errors"

	"github.com/jossemargt/cms-sao/app"
	"github.com/jossemargt/cms-sao/model"
	"github.com/jossemargt/cms-sao/storage"
)

func TestResultController_Get(t *testing.T) {
	scenarios := []struct {
		name             string
		mockRepo         storage.ResultRepository
		expectedResource *app.ComJossemargtSaoResultFull
		goaFnWrapper     func(*testing.T, context.Context, *goa.Service, app.ResultController) *app.ComJossemargtSaoResultFull
	}{
		{
			name: "Get Result of an Entry that only had been compiled",
			mockRepo: &mockResultRepository{
				result: &model.Result{
					EntryID:   5,
					DatasetID: 7,
					Compilation: model.Compilation{
						Status:        "ok",
						Tries:         1,
						Memory:        125,
						WallClockTime: 2,
						Time:          1,
					},
					Evaluation: model.Evaluation{Done: false},
				},
			},
			expectedResource: &app.ComJossemargtSaoResultFull{
				ID: "5-7",
				Compilation: &app.CompilationResult{
					Status:        "ok",
					Tries:         1,
					Memory:        125,
					WallClockTime: 2,
					Time:          1,
				},
				Evaluation: &app.EvaluationResult{
					Status: "unprocessed",
				},
				Score: &app.ScoreResult{},
				Href:  fmt.Sprintf("%s%s", app.ResultHref(), "5-7"),
			},
			goaFnWrapper: func(t *testing.T, c context.Context, s *goa.Service, ctrl app.ResultController) *app.ComJossemargtSaoResultFull {
				// The controller should call to Ok fn
				_, resource := goatesthelper.GetResultOKFull(t, c, s, ctrl, "5-7")
				return resource
			},
		},
		{
			name: "No Result found",
			mockRepo: &mockResultRepository{
				err: errors.Wrap(sql.ErrNoRows, "Test, no rows found"),
			},
			goaFnWrapper: func(t *testing.T, c context.Context, s *goa.Service, ctrl app.ResultController) *app.ComJossemargtSaoResultFull {
				// The controller should call to NotFound fn
				goatesthelper.GetResultNotFound(t, c, s, ctrl, "0-0")
				return nil
			},
		},
		{
			name:     "Bad request on wrong ID pattern",
			mockRepo: nil,
			goaFnWrapper: func(t *testing.T, c context.Context, s *goa.Service, ctrl app.ResultController) *app.ComJossemargtSaoResultFull {
				// The controller should call to NotFound fn
				goatesthelper.GetResultBadRequest(t, c, s, ctrl, "a-1")
				return nil
			},
		},
		{
			name:     "Should get Bad request on wrong ID",
			mockRepo: nil,
			goaFnWrapper: func(t *testing.T, c context.Context, s *goa.Service, ctrl app.ResultController) *app.ComJossemargtSaoResultFull {
				// The controller should call to NotFound fn
				goatesthelper.GetResultBadRequest(t, c, s, ctrl, "doesntfollowatall")
				return nil
			},
		},
	}

	var (
		service = goa.New("entry-test")
		ctx     = context.Background()
	)

	for _, tt := range scenarios {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := NewResultController(service, tt.mockRepo)
			resource := tt.goaFnWrapper(t, ctx, service, ctrl)

			if !reflect.DeepEqual(tt.expectedResource, resource) {
				t.Errorf("Unexpected response body, expeted: %#v got: %#v", tt.expectedResource, resource)
			}
		})
	}
}

func TestResultController_Show(t *testing.T) {
	scenarios := []struct {
		name             string
		mockRepo         storage.ResultRepository
		expectedResource app.ComJossemargtSaoResultCollection
		goaFnWrapper     func(*testing.T, context.Context, *goa.Service, app.ResultController) app.ComJossemargtSaoResultCollection
	}{
		{
			name: "Get a single Result that match search criteria",
			mockRepo: &mockResultRepository{
				results: []model.Result{
					model.Result{
						EntryID:   5,
						DatasetID: 7,
						Compilation: model.Compilation{
							Status:        "ok",
							Tries:         1,
							Memory:        125,
							WallClockTime: 2,
							Time:          1,
						},
						Evaluation: model.Evaluation{
							Done:  false,
							Tries: 1,
						},
					},
				},
			},
			expectedResource: []*app.ComJossemargtSaoResult{
				&app.ComJossemargtSaoResult{
					ID: "5-7",
					Evaluation: &app.EvaluationResult{
						Status: "unprocessed",
						Tries:  1,
					},
					Score: &app.ScoreResult{},
					Href:  fmt.Sprintf("%s%s", app.ResultHref(), "5-7"),
				},
			},
			goaFnWrapper: func(t *testing.T, c context.Context, s *goa.Service, ctrl app.ResultController) app.ComJossemargtSaoResultCollection {
				// The controller should call to Ok fn
				_, resource := goatesthelper.ShowResultOK(t, c, s, ctrl, 1, "", 1, false,
					1, 10, "desc", 1, "", 1, "default")
				return resource
			},
		},
		{
			name: "No Entries found",
			mockRepo: &mockResultRepository{
				err: errors.Wrap(sql.ErrNoRows, "Test, no rows found"),
			},
			expectedResource: []*app.ComJossemargtSaoResult{},
			goaFnWrapper: func(t *testing.T, c context.Context, s *goa.Service, ctrl app.ResultController) app.ComJossemargtSaoResultCollection {
				// The controller should call to Ok fn
				_, resource := goatesthelper.ShowResultOK(t, c, s, ctrl, 1, "", 1, false,
					1, 10, "desc", 1, "", 1, "default")
				return resource
			},
		},
	}

	var (
		service = goa.New("entry-test")
		ctx     = context.Background()
	)

	for _, tt := range scenarios {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := NewResultController(service, tt.mockRepo)
			resource := tt.goaFnWrapper(t, ctx, service, ctrl)

			if !reflect.DeepEqual(tt.expectedResource, resource) {
				t.Errorf("Unexpected response body, expected: %#v got: %#v", tt.expectedResource, resource)
			}
		})
	}
}

type mockResultRepository struct {
	result  *model.Result
	results []model.Result
	err     error
}

func (m *mockResultRepository) FindByID(_, _ int) (*model.Result, error) {
	if m.err != nil {
		return nil, m.err
	}

	return m.result, nil
}

func (m *mockResultRepository) FindBy(_ storage.ResultDTO) ([]model.Result, error) {
	if m.err != nil {
		return nil, m.err
	}

	return m.results, nil
}
