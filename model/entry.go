package model

const (
	KindUserTest = iota
	KindEntry
)

type Entry struct {
	ID          int64
	Kind        int
	ContestID   int64  `db:"contest_id"`
	ContestSlug string `db:"contest_slug"`
	TaskID      int64  `db:"task_id"`
	TaskSlug    string `db:"task_slug"`
	ResultID    int64  `db:"result_prtl_id"`
	ScoreID     int64
}


func (e *Entry) isDryRun () bool {
	if e.Kind == KindUserTest {
		return true
	}

	return false
}
