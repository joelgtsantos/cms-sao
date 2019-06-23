package model

import "time"

type EntrySubmitTrx struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    string
	EntryID   int
	ResultID  string
}
