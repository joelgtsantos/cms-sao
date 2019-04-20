package model

const (
	KindUserTest = iota
	KindEntry
)

type Entry struct {
	ID          int
	Kind        int
	ContestID   int    `db:"contest_id"`
	ContestSlug string `db:"contest_slug"`
	TaskID      int    `db:"task_id"`
	TaskSlug    string `db:"task_slug"`
	ResultID    int    `db:"result_prtl_id"`
	ScoreID     int
}

func (e *Entry) isDryRun() bool {
	if e.Kind == KindUserTest {
		return true
	}

	return false
}
