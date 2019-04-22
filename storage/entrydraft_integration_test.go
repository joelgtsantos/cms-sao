// +build integration

package storage

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// WARNING: The following tests are tied to a database connection and real data
func TestDefaultEntryDraftRepository_FindByIDIntegration(t *testing.T) {
	var dbCred dbCredentials
	getCredentialsFromEnv(&dbCred)
	db, err := sqlx.Connect("postgres", dbCred.connString())
	if err != nil {
		t.Fatal(err)
	}

	repo := NewEntryDraftRepository(db)

	entry, err := repo.FindByID(1)

	if err != nil {
		t.Fatalf("Failed querying for entry. %v", err)
	}

	t.Logf("%#v", entry)
}

// WARNING: The following tests are tied to a database connection and real data
func TestDefaultEntryDraftRepository_FindByIntegration(t *testing.T) {
	var dbCred dbCredentials
	getCredentialsFromEnv(&dbCred)
	db, err := sqlx.Connect("postgres", dbCred.connString())

	if err != nil {
		t.Fatal(err)
	}

	repo := NewEntryDraftRepository(db)

	entries, err := repo.FindBy(EntryDTO{
		TaskSlug: "batch",
		DTO: DTO{
			Page: 1,
		},
		ContestID: 0,
	})

	if err != nil {
		t.Fatalf("Failed querying for entry. %v", err)
	}

	t.Logf("Found '%d' entries", len(entries))

	for _, entry := range entries {
		t.Logf("%#v", entry)
	}
}
