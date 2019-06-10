package storage

import (
	"time"

	"github.com/jossemargt/cms-sao/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type entrySubmitTrx struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
	Status    string             `bson:"status"`
	EntryID   int                `bson:"entryID"`
	ResultID  int                `bson:"resultID"`
}

func (e entrySubmitTrx) toEntrySubmitTrx() *model.EntrySubmitTrx {
	return &model.EntrySubmitTrx{
		ID:        e.ID.Hex(),
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
		Status:    e.Status,
		EntryID:   e.EntryID,
		ResultID:  e.ResultID,
	}
}

func newEntrySubmitTrx(src model.EntrySubmitTrx) *entrySubmitTrx {
	mEntry := entrySubmitTrx{
		CreatedAt: src.CreatedAt,
		UpdatedAt: src.UpdatedAt,
		Status:    src.Status,
		EntryID:   src.EntryID,
		ResultID:  src.ResultID,
	}

	id, err := primitive.ObjectIDFromHex(src.ID)
	if src.ID != "" && err == nil {
		mEntry.ID = id
	}

	return &mEntry
}
