// +build integration

package storage

import (
	"fmt"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func init() {
	fmt.Println("WARNING: The following tests are tied to a database connection and real data")
}

func TestDefaultEntryRepository_FindByIDIntegration(t *testing.T) {
	var dbCred dbCredentials
	getCredentialsFromEnv(&dbCred)
	db, err := sqlx.Connect("postgres", dbCred.connString())
	if err != nil {
		t.Fatal(err)
	}

	repo := NewEntryRepository(db)

	entry, err := repo.FindByID(5)

	if err != nil {
		t.Fatalf("Failed querying for entry. %v", err)
	}

	t.Logf("%#v", entry)
}

func TestDefaultEntryRepository_FindByIntegration(t *testing.T) {
	var dbCred dbCredentials
	getCredentialsFromEnv(&dbCred)
	db, err := sqlx.Connect("postgres", dbCred.connString())

	if err != nil {
		t.Fatal(err)
	}

	repo := NewEntryRepository(db)

	entries, err := repo.FindBy(EntryDTO{
		TaskSlug: "batch_file",
		DTO: DTO{
			Page: 1,
		},
	})

	if err != nil {
		t.Fatalf("Failed querying for entry. %v", err)
	}

	t.Logf("Found '%d' entries", len(entries))

	for _, entry := range entries {
		t.Logf("%#v", entry)
	}
}

type dbCredentials struct {
	host     string
	dbname   string
	user     string
	password string
	sslmode  string
}

func (c dbCredentials) connString() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=%s",
		c.user,
		c.password,
		c.dbname,
		c.host,
		c.sslmode,
	)
}

func getCredentialsFromEnv(credentials *dbCredentials) {
	credentials.host = getEnvFallback("TEST_DBHOST", "192.168.7.10") // Vagrant box IP addr
	credentials.dbname = getEnvFallback("TEST_DBNAME", "cmsdb")
	credentials.user = getEnvFallback("TEST_DBUSER", "cmsuser")
	credentials.password = getEnvFallback("TEST_DBPSWD", "notsecure")
	credentials.sslmode = getEnvFallback("TEST_DBSSLMODE", "disable")
}

func getEnvFallback(key, fallback string) string {
	t := os.Getenv(key)
	if t != "" {
		return t
	}

	return fallback
}
