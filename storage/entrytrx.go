package storage

import (
	"context"

	"github.com/jossemargt/cms-sao/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const entryTrxCollectionName = "entry_trx"

type EntrySubmitTrxRepository interface {
	FindByID(ctx context.Context, id string) (*model.EntrySubmitTrx, error)
	Save(ctx context.Context, trx *model.EntrySubmitTrx) (string, error)
}

func NewEntrySubmitTrxRepository(db *mongo.Database) EntrySubmitTrxRepository {
	dbColl := &DocumentDBCollection{
		coll: db.Collection(entryTrxCollectionName),
	}

	return &defaultEntrySubmitTrxRepository{
		read:  dbColl,
		write: dbColl,
	}
}

type defaultEntrySubmitTrxRepository struct {
	read  DocumentQuerier
	write DocumentPersister
}

func (r *defaultEntrySubmitTrxRepository) FindByID(ctx context.Context, id string) (*model.EntrySubmitTrx, error) {
	var m entrySubmitTrx
	q, err := NewDocumentDBQuery(ID(id))
	if err != nil {
		return nil, err
	}

	err = r.read.FindOne(ctx, &m, q)
	return m.toEntrySubmitTrx(), err
}

func (r *defaultEntrySubmitTrxRepository) Save(ctx context.Context, trx *model.EntrySubmitTrx) (string, error) {
	if trx.ID == "" {
		return r.write.InsertOne(ctx, newEntrySubmitTrx(trx))
	}

	q, err := NewDocumentDBQuery(ID(trx.ID))
	if err != nil {
		return "", err
	}

	err = r.write.UpdateOne(ctx, newEntrySubmitTrx(trx), q)
	return "", err
}
