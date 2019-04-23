// +build integration

package storage

import (
	"testing"

	"github.com/jmoiron/sqlx"
)

// WARNING: The following tests are tied to a database connection and real data
func TestDefaultResultRepository_FindByIDIntegration(t *testing.T) {
	var dbCred dbCredentials
	getCredentialsFromEnv(&dbCred)
	db, err := sqlx.Connect("postgres", dbCred.connString())

	if err != nil {
		t.Fatal(err)
	}

	repo := defaultResultRepository{
		source: db,
	}

	result, err := repo.FindByID(1, 1)
	if err != nil {
		t.Fatalf("Failed querying for entry. %v", err)
	}

	t.Logf("%#v", result)
}

// WARNING: The following tests are tied to a database connection and real data
func TestDefaultResultRepository_FindByIntegration(t *testing.T) {
	var dbCred dbCredentials
	getCredentialsFromEnv(&dbCred)
	db, err := sqlx.Connect("postgres", dbCred.connString())

	if err != nil {
		t.Fatal(err)
	}

	repo := defaultResultRepository{
		source: db,
	}

	results, err := repo.FindBy(ResultDTO{
		DTO: DTO{
			Page: 1,
		},
		ContestID: 0,
		UserID:    1,
		MaxScore:  true,
	})

	if err != nil {
		t.Fatalf("Failed querying for Result. %v", err)
	}

	t.Logf("Found '%d' entries", len(results))

	for _, result := range results {
		// t.Logf("%T -> %d-%d", result, result.EntryID, result.DatasetID)
		t.Logf("id: %d-%d, score: %f", result.EntryID, result.DatasetID, result.TaskScore)
	}
}
