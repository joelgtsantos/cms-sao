package storage

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DocumentDBCollection struct {
	coll *mongo.Collection
}

type DocumentDBQuery struct {
	body bson.M
}

type DocumentQuerier interface {
	FindOne(ctx context.Context, destination interface{}, query *DocumentDBQuery) error
}

type DocumentPersister interface {
	InsertOne(ctx context.Context, subject interface{}) (string, error)
	UpdateOne(ctx context.Context, subject interface{}, query *DocumentDBQuery) error
}

func (m *DocumentDBCollection) FindOne(ctx context.Context, destination interface{}, query *DocumentDBQuery) error {
	r := m.coll.FindOne(ctx, query.body)
	if r.Err() != nil {
		errors.Wrap(r.Err(), "Failed fetching value")
	}

	if err := r.Decode(destination); err != nil {
		return errors.Wrap(err, "Failed decoding value")
	}

	return nil
}

func (m *DocumentDBCollection) InsertOne(ctx context.Context, subject interface{}) (string, error) {
	r, err := m.coll.InsertOne(ctx, subject)
	if err != nil {
		return "", err
	}

	objID, ok := r.InsertedID.(primitive.ObjectID)
	if ok {
		return objID.Hex(), nil
	}

	return fmt.Sprint(r.InsertedID), nil
}

func (m *DocumentDBCollection) UpdateOne(ctx context.Context, subject interface{}, query *DocumentDBQuery) error {
	_, err := m.coll.UpdateOne(ctx, subject, query.body)

	return err
}

func NewDocumentDBQuery(options ...DocumentDBQueryBuilderOption) (*DocumentDBQuery, error) {
	query := &DocumentDBQuery{
		body: make(primitive.M, 0),
	}

	for _, opt := range options {
		if err := opt(query); err != nil {
			return nil, err
		}
	}

	return query, nil
}

type DocumentDBQueryBuilderOption func(b *DocumentDBQuery) error

func ID(id string) DocumentDBQueryBuilderOption {
	return func(b *DocumentDBQuery) error {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return errors.Wrapf(err, "Unable to create ObjectID from %s", id)
		}

		b.body["_id"] = objID
		return nil
	}
}

func Field(name string, value interface{}) DocumentDBQueryBuilderOption {
	return func(b *DocumentDBQuery) error {
		b.body[name] = value
		return nil
	}
}
