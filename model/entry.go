package model

import "fmt"

type Entry struct {
	ID          int
	ContestID   int     `db:"contest_id"`
	ContestSlug string  `db:"contest_slug"`
	TaskID      int     `db:"task_id"`
	TaskSlug    string  `db:"task_slug"`
	DatasetID   int     `db:"result_prtl_id"`
	Language    *string `db:"language"`
	Token       bool    `db:"token"`
}

func (e Entry) ResultID() string {
	if e.DatasetID == 0 {
		return ""
	}

	return fmt.Sprintf("%d-%d", e.ID, e.DatasetID)
}
