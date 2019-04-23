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

func TestEntryController_Get(t *testing.T) {
	scenarios := []struct {
		name             string
		mockRepo         storage.EntryRepository
		expectedResource *app.ComJossemargtSaoEntryFull
		goaFnWrapper     func(*testing.T, context.Context, *goa.Service, app.EntryController) *app.ComJossemargtSaoEntryFull
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
			expectedResource: &app.ComJossemargtSaoEntryFull{
				ID:          5,
				ContestID:   7,
				TaskID:      7,
				ContestSlug: "con_test",
				TaskSlug:    "batch_test",
				Token:       true,
				Href:        fmt.Sprintf("%s%d", app.EntryHref(), 5),
			},
			goaFnWrapper: func(t *testing.T, c context.Context, s *goa.Service, ctrl app.EntryController) *app.ComJossemargtSaoEntryFull {
				// The controller should call to Ok fn
				_, resource := goatesthelper.GetEntryOKFull(t, c, s, ctrl, 5)
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
			expectedResource: &app.ComJossemargtSaoEntryFull{
				ID:          5,
				ContestID:   7,
				TaskID:      7,
				ContestSlug: "con_test",
				TaskSlug:    "batch_test",
				Token:       true,
				Href:        fmt.Sprintf("%s%d", app.EntryHref(), 5),
				Links: &app.ComJossemargtSaoEntryLinks{
					Result: &app.ComJossemargtSaoResultLink{
						ID:   "5-7",
						Href: fmt.Sprintf("%s%s", app.ResultHref(), "5-7"),
					},
				},
			},
			goaFnWrapper: func(t *testing.T, c context.Context, s *goa.Service, ctrl app.EntryController) *app.ComJossemargtSaoEntryFull {
				// The controller should call to Ok fn
				_, resource := goatesthelper.GetEntryOKFull(t, c, s, ctrl, 5)
				return resource
			},
		},
		{
			name: "No Entry found",
			mockRepo: &mockEntryRepository{
				err: errors.Wrap(sql.ErrNoRows, "Test, no rows found"),
			},
			goaFnWrapper: func(t *testing.T, c context.Context, s *goa.Service, ctrl app.EntryController) *app.ComJossemargtSaoEntryFull {
				// The controller should call to NotFound fn
				goatesthelper.GetEntryNotFound(t, c, s, ctrl, -1)
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
			ctrl := NewEntryController(service, tt.mockRepo)
			resource := tt.goaFnWrapper(t, ctx, service, ctrl)

			if !reflect.DeepEqual(tt.expectedResource, resource) {
				t.Errorf("Unexpected response body, expeted: %#v got: %#v", tt.expectedResource, resource)
			}
		})
	}
}

func TestEntryController_Show(t *testing.T) {
	scenarios := []struct {
		name             string
		mockRepo         storage.EntryRepository
		expectedResource app.ComJossemargtSaoEntryCollection
		goaFnWrapper     func(*testing.T, context.Context, *goa.Service, app.EntryController) app.ComJossemargtSaoEntryCollection
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
			expectedResource: []*app.ComJossemargtSaoEntry{
				&app.ComJossemargtSaoEntry{
					ID:          5,
					ContestSlug: "con_test",
					TaskSlug:    "batch_test",
					Token:       true,
					Href:        fmt.Sprintf("%s%d", app.EntryHref(), 5),
				},
			},
			goaFnWrapper: func(t *testing.T, c context.Context, s *goa.Service, ctrl app.EntryController) app.ComJossemargtSaoEntryCollection {
				// The controller should call to Ok fn
				_, resource := goatesthelper.ShowEntryOK(t, c, s, ctrl, 1, "", 1, 5, "desc", 1, "", 1)
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
			expectedResource: func() []*app.ComJossemargtSaoEntry {
				list := make([]*app.ComJossemargtSaoEntry, 0, 3)
				for i := 1; i < 4; i++ {
					list = append(list, &app.ComJossemargtSaoEntry{
						ID:          i,
						ContestSlug: "con_test",
						TaskSlug:    "batch_test",
						Token:       true,
						Href:        fmt.Sprintf("%s%d", app.EntryHref(), i),
					})
				}
				return list
			}(),
			goaFnWrapper: func(t *testing.T, c context.Context, s *goa.Service, ctrl app.EntryController) app.ComJossemargtSaoEntryCollection {
				// The controller should call to Ok fn
				_, resource := goatesthelper.ShowEntryOK(t, c, s, ctrl, 1, "", 1, 5, "desc", 1, "", 1)
				return resource
			},
		},
		{
			name: "No Entries found",
			mockRepo: &mockEntryRepository{
				err: errors.Wrap(sql.ErrNoRows, "Test, no rows found"),
			},
			expectedResource: []*app.ComJossemargtSaoEntry{},
			goaFnWrapper: func(t *testing.T, c context.Context, s *goa.Service, ctrl app.EntryController) app.ComJossemargtSaoEntryCollection {
				// The controller should call to Ok fn
				_, resource := goatesthelper.ShowEntryOK(t, c, s, ctrl, 1, "", 1, 5, "desc", 1, "", 1)
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
			ctrl := NewEntryController(service, tt.mockRepo)
			resource := tt.goaFnWrapper(t, ctx, service, ctrl)

			if !reflect.DeepEqual(tt.expectedResource, resource) {
				t.Errorf("Unexpected response body, expeted: %#v got: %#v", tt.expectedResource, resource)
			}
		})
	}
}

type mockEntryRepository struct {
	entry   *model.Entry
	entries []model.Entry
	err     error
}

func (m *mockEntryRepository) FindByID(_ int) (*model.Entry, error) {
	if m.err != nil {
		return nil, m.err
	}

	return m.entry, nil
}

func (m *mockEntryRepository) FindBy(_ storage.EntryDTO) ([]model.Entry, error) {
	if m.err != nil {
		return nil, m.err
	}

	return m.entries, nil
}
