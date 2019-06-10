package storage

import "go.mongodb.org/mongo-driver/mongo"

const draftTrxCollectionName = "draft_trx"

func NewDraftSubmitTrxRepository(db *mongo.Database) EntrySubmitTrxRepository {
	dbColl := &DocumentDBCollection{
		coll: db.Collection(draftTrxCollectionName),
	}

	return &defaultEntrySubmitTrxRepository{
		read:  dbColl,
		write: dbColl,
	}
}
