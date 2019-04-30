package model

import "fmt"

type Entry struct {
	ID          int
	ContestID   int
	ContestSlug string
	TaskID      int
	TaskSlug    string
	DatasetID   int
	Language    string
	Token       bool
}

func (e Entry) ResultID() string {
	if e.DatasetID == 0 {
		return ""
	}

	return fmt.Sprintf("%d-%d", e.ID, e.DatasetID)
}
