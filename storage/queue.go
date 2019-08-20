package storage

import (
	"context"
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const queueName = "sao_queue"

type QueueWriter interface {
	Write(ctx context.Context, payload interface{}) error
}

func NewQueueWriter(db *mongo.Database) QueueWriter {
	dbColl := &DocumentDBCollection{
		coll: db.Collection(queueName),
	}
	return &defaultQueueWriter{
		write: dbColl,
	}
}

type defaultQueueWriter struct {
	write DocumentPersister
}

type queueMessage struct {
	Timestamp time.Time
	Payload   interface{}
}

func (m *queueMessage) MarshalBSON() ([]byte, error) {
	payload, err := json.Marshal(m.Payload)
	if err != nil {
		return nil, err
	}

	return bson.Marshal(&struct {
		T string      `bson:"visible"`
		P interface{} `bson:"payload"`
	}{
		T: m.Timestamp.UTC().Format(time.RFC3339Nano),
		P: string(payload),
	})
}

func (r *defaultQueueWriter) Write(ctx context.Context, payload interface{}) error {
	message := &queueMessage{
		Timestamp: time.Now(),
		Payload:   payload,
	}

	_, err := r.write.InsertOne(ctx, message)
	return err
}
