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

func TestDraftResultController_Get(t *testing.T) {
	okString := "ok"

	scenarios := []struct {
		name             string
		mockRepo         storage.DraftResultRepository
		expectedResource *app.ComJossemargtSaoDraftResultFull
		goaFnWrapper     func(*testing.T, context.Context, *goa.Service, app.DraftresultController) *app.ComJossemargtSaoDraftResultFull
	}{
		{
			name: "Get Result of an Entry that only had been compiled",
			mockRepo: &mockDraftResultRepository{
				result: &model.DraftResult{
					EntryID:   5,
					DatasetID: 7,
					Compilation: model.Compilation{
						Status:        &okString,
						Tries:         1,
						Memory:        125,
						WallClockTime: 2,
						Time:          1,
					},
					Evaluation: model.Evaluation{Done: false},
				},
			},
			expectedResource: &app.ComJossemargtSaoDraftResultFull{
				ID: "5-7",
				Compilation: &app.CompilationResult{
					Status:        okString,
					Tries:         1,
					Memory:        125,
					WallClockTime: 2,
					Time:          1,
				},
				Evaluation: &app.EvaluationResult{
					Status: "unprocessed",
				},
				Execution: &app.ExecutionResult{},
				Href:      fmt.Sprintf("%s%s", app.DraftresultHref(), "5-7"),
			},
			goaFnWrapper: func(t *testing.T, c context.Context, s *goa.Service, ctrl app.DraftresultController) *app.ComJossemargtSaoDraftResultFull {
				// The controller should call to Ok fn
				_, resource := goatesthelper.GetDraftresultOKFull(t, c, s, ctrl, "5-7")
				return resource
			},
		},
		{
			name: "No Result found",
			mockRepo: &mockDraftResultRepository{
				err: errors.Wrap(sql.ErrNoRows, "Test, no rows found"),
			},
			goaFnWrapper: func(t *testing.T, c context.Context, s *goa.Service, ctrl app.DraftresultController) *app.ComJossemargtSaoDraftResultFull {
				// The controller should call to NotFound fn
				goatesthelper.GetDraftresultNotFound(t, c, s, ctrl, "0-0")
				return nil
			},
		},
		{
			name:     "Bad request on wrong ID pattern",
			mockRepo: nil,
			goaFnWrapper: func(t *testing.T, c context.Context, s *goa.Service, ctrl app.DraftresultController) *app.ComJossemargtSaoDraftResultFull {
				// The controller should call to NotFound fn
				goatesthelper.GetDraftresultBadRequest(t, c, s, ctrl, "a-1")
				return nil
			},
		},
		{
			name:     "Should get Bad request on wrong ID",
			mockRepo: nil,
			goaFnWrapper: func(t *testing.T, c context.Context, s *goa.Service, ctrl app.DraftresultController) *app.ComJossemargtSaoDraftResultFull {
				// The controller should call to NotFound fn
				goatesthelper.GetDraftresultBadRequest(t, c, s, ctrl, "doesntfollowatall")
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
			ctrl := NewDraftResultController(service, tt.mockRepo)
			resource := tt.goaFnWrapper(t, ctx, service, ctrl)

			if !reflect.DeepEqual(tt.expectedResource, resource) {
				t.Errorf("Unexpected response body, expeted: %#v got: %#v", tt.expectedResource, resource)
			}
		})
	}
}

func TestDraftResultController_Show(t *testing.T) {
	scenarios := []struct {
		name             string
		mockRepo         storage.DraftResultRepository
		expectedResource app.ComJossemargtSaoDraftResultCollection
		goaFnWrapper     func(*testing.T, context.Context, *goa.Service, app.DraftresultController) app.ComJossemargtSaoDraftResultCollection
	}{
		{
			name: "Get a single Result that match search criteria",
			mockRepo: &mockDraftResultRepository{
				results: []model.DraftResult{
					model.DraftResult{
						EntryID:   5,
						DatasetID: 7,
						Execution: model.Execution{
							Memory:        12563,
							WallClockTime: 1,
							Time:          0,
							Output:        []byte("This shouldn't apper in the default view"),
						},
					},
				},
			},
			expectedResource: []*app.ComJossemargtSaoDraftResult{
				&app.ComJossemargtSaoDraftResult{
					ID: "5-7",
					Execution: &app.ExecutionResult{
						Memory:        12563,
						WallClockTime: 1,
						Time:          0,
						Output:        "",
					},
					Href: fmt.Sprintf("%s%s", app.DraftresultHref(), "5-7"),
				},
			},
			goaFnWrapper: func(t *testing.T, c context.Context, s *goa.Service, ctrl app.DraftresultController) app.ComJossemargtSaoDraftResultCollection {
				// The controller should call to Ok fn
				_, resource := goatesthelper.ShowDraftresultOK(t, c, s, ctrl, 1, "", 1,
					1, 10, "desc", 1, "", 1)
				return resource
			},
		},
		{
			name: "No Entries found",
			mockRepo: &mockDraftResultRepository{
				err: errors.Wrap(sql.ErrNoRows, "Test, no rows found"),
			},
			expectedResource: []*app.ComJossemargtSaoDraftResult{},
			goaFnWrapper: func(t *testing.T, c context.Context, s *goa.Service, ctrl app.DraftresultController) app.ComJossemargtSaoDraftResultCollection {
				// The controller should call to Ok fn
				_, resource := goatesthelper.ShowDraftresultOK(t, c, s, ctrl, 1, "", 1,
					1, 10, "desc", 1, "", 1)
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
			ctrl := NewDraftResultController(service, tt.mockRepo)
			resource := tt.goaFnWrapper(t, ctx, service, ctrl)

			if !reflect.DeepEqual(tt.expectedResource, resource) {
				t.Errorf("Unexpected response body, expected: %#v got: %#v", tt.expectedResource, resource)
			}
		})
	}
}

type mockDraftResultRepository struct {
	result  *model.DraftResult
	results []model.DraftResult
	err     error
}

func (m *mockDraftResultRepository) FindByID(_, _ int) (*model.DraftResult, error) {
	if m.err != nil {
		return nil, m.err
	}

	return m.result, nil
}

func (m *mockDraftResultRepository) FindBy(_ storage.DraftResultDTO) ([]model.DraftResult, error) {
	if m.err != nil {
		return nil, m.err
	}

	return m.results, nil
}
