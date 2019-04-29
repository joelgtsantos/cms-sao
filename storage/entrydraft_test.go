package storage

import (
	"database/sql"
	"reflect"
	"strings"
	"testing"
)

func TestDefaultEntryDraftRepository_FindByID(t *testing.T) {
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
					e := sqlEntry{
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
			repo := NewEntryDraftRepository(tt.queryer)

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

func TestDefaultEntryDraftRepository_FindBy(t *testing.T) {
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
					es := []sqlEntry{
						sqlEntry{
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
					es := []sqlEntry{
						sqlEntry{
							ID:          123,
							ContestSlug: "con_test",
							TaskSlug:    "a_task",
						},
						sqlEntry{
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
			repo := NewEntryDraftRepository(db)

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

func Test_buildEntryDraftFindByIDQuery(t *testing.T) {
	query, err := buildEntryDraftFindByIDQuery()
	if err != nil {
		t.Fatalf("Un-expeted error, %v", err)
	}

	keywords := []string{
		entryDraftTable,
		entryDraftResultTable,
		"utr.user_test_id = ut.id",
	}

	for _, word := range keywords {
		if !strings.Contains(query, word) {
			t.Errorf("Expected query containing '%s' keyword", word)
		}
	}
}

func Test_buildEntryDraftFindQuery(t *testing.T) {
	scenarios := []struct {
		name          string
		dto           EntryDTO
		expectedStmts []string
		excludedStmts []string
	}{
		{
			name: "Base query",
			dto:  EntryDTO{},
			expectedStmts: []string{
				entryDraftTable,
				entryDraftResultTable,
				"user_test_id = ut.id",
			},
			excludedStmts: []string{
				entryTable,
				resultTable,
				"OFFSET",
				"tsk.name LIKE",
				"cts.name LIKE",
			},
		},
		{
			name: "Base query with OFFSET",
			dto: EntryDTO{
				DTO: DTO{
					PageSize: 5,
					Page:     3,
				},
			},
			expectedStmts: []string{
				entryDraftTable,
				entryDraftResultTable,
				"user_test_id = ut.id",
				"LIMIT 5",
				"OFFSET 10",
			},
			excludedStmts: []string{
				"tsk.name LIKE",
				"cts.name LIKE",
			},
		},
		{
			name: "Base query with non-zero Entry DTO options",
			dto: EntryDTO{
				DTO: DTO{
					PageSize: 5,
					Page:     1,
				},
				TaskID:      7,
				TaskSlug:    "foo",
				ContestID:   7,
				ContestSlug: "bar",
			},
			expectedStmts: []string{
				entryDraftTable,
				entryDraftResultTable,
				"user_test_id = ut.id",
				"LIMIT 5",
				"tsk.id = 7",
				"tsk.name LIKE 'foo'",
				"cts.id = 7",
				"cts.name LIKE 'bar",
			},
			excludedStmts: []string{
				"OFFSET",
			},
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.name, func(t *testing.T) {
			query, err := buildEntryDraftFindQuery(tt.dto)
			if err != nil {
				t.Errorf("Un-expeted error, %v", err)
				return
			}

			for _, stmt := range tt.expectedStmts {
				if !strings.Contains(query, stmt) {
					t.Errorf("Expected query containing '%s' statement", stmt)
				}
			}

			for _, stmt := range tt.excludedStmts {
				if strings.Contains(query, stmt) {
					t.Errorf("Un-expected '%s' statement in query", stmt)
				}
			}
		})
	}
}
