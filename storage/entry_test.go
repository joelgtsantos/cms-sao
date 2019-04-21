package storage

import (
	"database/sql"
	"reflect"
	"strings"
	"testing"

	"github.com/jossemargt/cms-sao/model"
)

func TestDefaultEntryRepository_FindByID(t *testing.T) {
	scenarios := []struct {
		name       string
		queryer    *mockDB
		shouldFail bool
	}{
		{
			name: "Gets single result from database",
			queryer: func() *mockDB {
				m := new(mockDB)
				m.fnGet = func(d interface{}, q string, a ...interface{}) error {
					e := model.Entry{
						ID:          123,
						ContestSlug: "con_test",
						TaskSlug:    "a_task",
					}

					// Kids don't do this at home, unless you need to
					reflect.ValueOf(d).Elem().Set(reflect.ValueOf(e))
					return nil
				}
				return m
			}(),
			shouldFail: false,
		},
		{
			name: "Fails finding result from database",
			queryer: func() *mockDB {
				m := new(mockDB)
				m.fnGet = func(d interface{}, q string, a ...interface{}) error {
					return sql.ErrNoRows
				}
				return m
			}(),
			shouldFail: true,
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewEntryRepository(tt.queryer)

			entry, err := repo.FindByID(123)

			if err != nil {
				if tt.shouldFail {
					return
				}

				t.Errorf("Unexpected error %v", err)
			} else if tt.shouldFail {
				t.Error("An Error was expected")
			}

			if entry == nil || entry.ID == 0 {
				t.Errorf("Expected entry to be found, got: %#v", entry)
			}
		})
	}
}

func TestDefaultEntryRepository_FindBy(t *testing.T) {
	scenarios := []struct {
		name           string
		queryerFactory func(t *testing.T) *mockDB
		dto            EntryDTO
		shouldFail     bool
	}{
		{
			name: "Gets single result from database using DTO zero values",
			queryerFactory: func(t *testing.T) *mockDB {
				m := new(mockDB)
				m.fnQuery = func(d interface{}, q string, a ...interface{}) error {
					es := []model.Entry{
						model.Entry{
							ID:          123,
							ContestSlug: "con_test",
							TaskSlug:    "a_task",
						},
					}
					reflect.ValueOf(d).Elem().Set(reflect.ValueOf(es))

					// Checking query
					if !strings.Contains(q, "LIMIT 10") {
						t.Error("Expected LIMIT 10 in query statement")
					}

					if strings.Contains(q, "OFFSET") {
						t.Error("Expected no OFFSET in query statement for zero value DTO")
					}

					return nil
				}
				return m
			},
			dto:        EntryDTO{},
			shouldFail: false,
		},
		{
			name: "Gets multiple results from database",
			queryerFactory: func(t *testing.T) *mockDB {
				m := new(mockDB)
				m.fnQuery = func(d interface{}, q string, a ...interface{}) error {
					es := []model.Entry{
						model.Entry{
							ID:          123,
							ContestSlug: "con_test",
							TaskSlug:    "a_task",
						},
						model.Entry{
							ID:          124,
							ContestSlug: "con_test",
							TaskSlug:    "b_task",
						},
					}
					reflect.ValueOf(d).Elem().Set(reflect.ValueOf(es))

					// Checking query
					if !strings.Contains(q, "LIMIT 15") {
						t.Error("Expected LIMIT 15 in query statement")
					}

					if !strings.Contains(q, "OFFSET 15") {
						t.Error("Expected OFFSET 15 in query statement")
					}

					return nil
				}
				return m
			},
			dto: EntryDTO{
				DTO: DTO{
					Page:     2,
					PageSize: 15,
				},
			},
			shouldFail: false,
		},
		{
			name: "Fails finding result from database",
			queryerFactory: func(t *testing.T) *mockDB {
				m := new(mockDB)
				m.fnQuery = func(d interface{}, q string, a ...interface{}) error {
					return sql.ErrNoRows
				}
				return m
			},
			shouldFail: true,
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.name, func(t *testing.T) {
			db := tt.queryerFactory(t)
			repo := NewEntryRepository(db)

			entries, err := repo.FindBy(tt.dto)

			if err != nil {
				if tt.shouldFail {
					return
				}

				t.Errorf("Unexpected error %v", err)
			} else if tt.shouldFail {
				t.Error("An Error was expected")
			}

			if entries == nil {
				t.Error("Expected entries to be found, got: nil")
			}
		})
	}
}

type mockDB struct {
	fnGet   func(interface{}, string, ...interface{}) error
	fnQuery func(interface{}, string, ...interface{}) error
}

func (m *mockDB) Select(destination interface{}, query string, args ...interface{}) error {
	return m.fnQuery(destination, query, args...)
}

func (m *mockDB) Get(destination interface{}, query string, args ...interface{}) error {
	return m.fnGet(destination, query, args...)
}
