// +build integration

package storage

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jossemargt/cms-sao/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestDefaultEntrySubmitTrxRepository_Save(t *testing.T) {
	var mongoCred mongoDBCredentials
	getMongoCredentialsFromEnv(&mongoCred)

	client, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI(mongoCred.connString()),
	)

	dbColl := &DocumentDBCollection{
		coll: client.Database("cmsdb").Collection("entry_trx_test"),
	}

	repo := &defaultEntrySubmitTrxRepository{
		read:  dbColl,
		write: dbColl,
	}

	id, err := repo.Save(context.Background(), &model.EntrySubmitTrx{
		CreatedAt: time.Date(2000, time.June, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Now(),
		Status:    "created",
		EntryID:   -1,
		ResultID:  "",
	})

	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)
}

func TestDefaultEntrySubmitTrxRepository_FindByID(t *testing.T) {
	var mongoCred mongoDBCredentials
	getMongoCredentialsFromEnv(&mongoCred)

	client, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI(mongoCred.connString()),
	)

	dbColl := &DocumentDBCollection{
		coll: client.Database("cmsdb").Collection("entry_trx_test"),
	}

	repo := &defaultEntrySubmitTrxRepository{
		read:  dbColl,
		write: dbColl,
	}

	id, err := repo.Save(context.Background(), &model.EntrySubmitTrx{
		CreatedAt: time.Date(2000, time.June, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Now(),
		Status:    "created",
		EntryID:   -1,
		ResultID:  "",
	})

	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)

	entry, err := repo.FindByID(context.Background(), id)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(entry)
}

type mongoDBCredentials struct {
	host     string
	dbname   string
	user     string
	password string
}

func (c mongoDBCredentials) connString() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:27017/%s",
		c.user,
		c.password,
		c.host,
		c.dbname,
	)
}

func getMongoCredentialsFromEnv(credentials *mongoDBCredentials) {
	credentials.host = getEnvFallback("TEST_MONGO_HOST", "192.168.99.100") // Docker machine IP
	credentials.dbname = getEnvFallback("TEST_MONGO_DBNAME", "cmsdb")
	credentials.user = getEnvFallback("TEST_MONGO_USER", "cmsuser")
	credentials.password = getEnvFallback("TEST_MONGO_PSWD", "notsecure")
}
