// +build integration

package storage

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestDefaultQueueWriter_Write(t *testing.T) {
	var mongoCred mongoDBCredentials
	getMongoCredentialsFromEnv(&mongoCred)

	client, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI(mongoCred.connString()),
	)

	dbColl := &DocumentDBCollection{
		coll: client.Database("cmsdb").Collection("test_queue"),
	}

	writer := &defaultQueueWriter{
		write: dbColl,
	}

	err = writer.Write(context.Background(), &struct {
		Foo   string
		Count int
	}{
		Foo:   "bar",
		Count: 0,
	})

	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
}
