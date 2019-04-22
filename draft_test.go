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

func TestDraftController_Get(t *testing.T) {
	scenarios := []struct {
		name             string
		mockRepo         storage.EntryRepository
		expectedResource *app.ComJossemargtSaoDraftFull
		goaFnWrapper     func(*testing.T, context.Context, *goa.Service, app.DraftController) *app.ComJossemargtSaoDraftFull
	}{
		{
			name: "Get existing unprocessed Entry (without linked Result nor Score)",
			mockRepo: &mockEntryRepository{
				entry: &model.Entry{
					ID:          5,
					ContestID:   7,
					TaskID:      7,
					TaskSlug:    "batch_test",
					ContestSlug: "con_test",
				},
			},
			expectedResource: &app.ComJossemargtSaoDraftFull{
				ID:          5,
				ContestID:   7,
				TaskID:      7,
				ContestSlug: "con_test",
				TaskSlug:    "batch_test",
				Href:        fmt.Sprintf("%s%d", app.DraftHref(), 5),
			},
			goaFnWrapper: func(t *testing.T, c context.Context, s *goa.Service, ctrl app.DraftController) *app.ComJossemargtSaoDraftFull {
				// The controller should call to Ok fn
				_, resource := goatesthelper.GetDraftOKFull(t, c, s, ctrl, 5)
				return resource
			},
		},
		{
			name: "Get existing Entry",
			mockRepo: &mockEntryRepository{
				entry: &model.Entry{
					ID:          5,
					ContestID:   7,
					TaskID:      7,
					TaskSlug:    "batch_test",
					ContestSlug: "con_test",
					DatasetID:   7,
				},
			},
			expectedResource: &app.ComJossemargtSaoDraftFull{
				ID:          5,
				ContestID:   7,
				TaskID:      7,
				ContestSlug: "con_test",
				TaskSlug:    "batch_test",
				Href:        fmt.Sprintf("%s%d", app.DraftHref(), 5),
				Links: &app.ComJossemargtSaoDraftLinks{
					Result: &app.ComJossemargtSaoDraftResultLink{
						ID:   "5-7",
						Href: fmt.Sprintf("%s%s", app.DraftresultHref(), "5-7"),
					},
				},
			},
			goaFnWrapper: func(t *testing.T, c context.Context, s *goa.Service, ctrl app.DraftController) *app.ComJossemargtSaoDraftFull {
				// The controller should call to Ok fn
				_, resource := goatesthelper.GetDraftOKFull(t, c, s, ctrl, 5)
				return resource
			},
		},
		{
			name: "No Entry found",
			mockRepo: &mockEntryRepository{
				err: errors.Wrap(sql.ErrNoRows, "Test, no rows found"),
			},
			goaFnWrapper: func(t *testing.T, c context.Context, s *goa.Service, ctrl app.DraftController) *app.ComJossemargtSaoDraftFull {
				// The controller should call to NotFound fn
				goatesthelper.GetDraftNotFound(t, c, s, ctrl, -1)
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
			ctrl := NewDraftController(service, tt.mockRepo)
			resource := tt.goaFnWrapper(t, ctx, service, ctrl)

			if !reflect.DeepEqual(tt.expectedResource, resource) {
				t.Errorf("Unexpected response body, expeted: %#v got: %#v", tt.expectedResource, resource)
			}
		})
	}
}

func TestDraftController_Show(t *testing.T) {
	scenarios := []struct {
		name             string
		mockRepo         storage.EntryRepository
		expectedResource app.ComJossemargtSaoDraftCollection
		goaFnWrapper     func(*testing.T, context.Context, *goa.Service, app.DraftController) app.ComJossemargtSaoDraftCollection
	}{
		{
			name: "Get a single Entry that match search criteria",
			mockRepo: &mockEntryRepository{
				entries: []model.Entry{
					model.Entry{
						ID:          5,
						ContestID:   7,
						TaskID:      7,
						TaskSlug:    "batch_test",
						ContestSlug: "con_test",
					},
				},
			},
			expectedResource: []*app.ComJossemargtSaoDraft{
				&app.ComJossemargtSaoDraft{
					ID:          5,
					ContestSlug: "con_test",
					TaskSlug:    "batch_test",
					Href:        fmt.Sprintf("%s%d", app.DraftHref(), 5),
				},
			},
			goaFnWrapper: func(t *testing.T, c context.Context, s *goa.Service, ctrl app.DraftController) app.ComJossemargtSaoDraftCollection {
				// The controller should call to Ok fn
				_, resource := goatesthelper.ShowDraftOK(t, c, s, ctrl, 1, "", 1, 5, "desc", 1, "", 1)
				return resource
			},
		},
		{
			name: "Get set of existing Entries that match search criteria",
			mockRepo: &mockEntryRepository{
				entries: func() []model.Entry {
					list := make([]model.Entry, 0, 3)
					for i := 1; i < 4; i++ {
						list = append(list, model.Entry{
							ID:          i,
							ContestID:   7,
							TaskID:      7,
							TaskSlug:    "batch_test",
							ContestSlug: "con_test",
						})
					}
					return list
				}(),
			},
			expectedResource: func() []*app.ComJossemargtSaoDraft {
				list := make([]*app.ComJossemargtSaoDraft, 0, 3)
				for i := 1; i < 4; i++ {
					list = append(list, &app.ComJossemargtSaoDraft{
						ID:          i,
						ContestSlug: "con_test",
						TaskSlug:    "batch_test",
						Href:        fmt.Sprintf("%s%d", app.DraftHref(), i),
					})
				}
				return list
			}(),
			goaFnWrapper: func(t *testing.T, c context.Context, s *goa.Service, ctrl app.DraftController) app.ComJossemargtSaoDraftCollection {
				// The controller should call to Ok fn
				_, resource := goatesthelper.ShowDraftOK(t, c, s, ctrl, 1, "", 1, 5, "desc", 1, "", 1)
				return resource
			},
		},
		{
			name: "No Entries found",
			mockRepo: &mockEntryRepository{
				err: errors.Wrap(sql.ErrNoRows, "Test, no rows found"),
			},
			expectedResource: []*app.ComJossemargtSaoDraft{},
			goaFnWrapper: func(t *testing.T, c context.Context, s *goa.Service, ctrl app.DraftController) app.ComJossemargtSaoDraftCollection {
				// The controller should call to Ok fn
				_, resource := goatesthelper.ShowDraftOK(t, c, s, ctrl, 1, "", 1, 5, "desc", 1, "", 1)
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
			ctrl := NewDraftController(service, tt.mockRepo)
			resource := tt.goaFnWrapper(t, ctx, service, ctrl)

			if !reflect.DeepEqual(tt.expectedResource, resource) {
				t.Errorf("Unexpected response body, expeted: %#v got: %#v", tt.expectedResource, resource)
			}
		})
	}
}
