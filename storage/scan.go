package storage

import (
	"database/sql"

	"github.com/jossemargt/cms-sao/model"
)

const (
	cmsLanguageDefault          = "none"
	cmsCompilationStatusDefault = "unprocessed"
)

type sqlEntry struct {
	ID          int
	ContestID   int            `db:"contest_id"`
	ContestSlug string         `db:"contest_slug"`
	TaskID      int            `db:"task_id"`
	TaskSlug    string         `db:"task_slug"`
	DatasetID   sql.NullInt64  `db:"result_prtl_id"`
	Language    sqlCMSLanguage `db:"language"`
	Token       bool           `db:"token"`
}

func (e sqlEntry) toEntry() model.Entry {
	return model.Entry{
		ID:          e.ID,
		ContestID:   e.ContestID,
		ContestSlug: e.ContestSlug,
		TaskID:      e.TaskID,
		TaskSlug:    e.TaskSlug,
		DatasetID:   int(e.DatasetID.Int64),
		Token:       e.Token,
		Language:    e.Language.String,
	}
}

type sqlResult struct {
	DatasetID int `db:"dataset_id"`
	EntryID   int `db:"entry_id"`
	compilation
	evaluation
	scoring
}

func (r sqlResult) toResult() model.Result {
	return model.Result{
		EntryID:   r.EntryID,
		DatasetID: r.DatasetID,
		Compilation: model.Compilation{
			Status:        r.compilation.Status.String,
			Tries:         r.compilation.Tries,
			Stdout:        r.compilation.Stdout.String,
			Stderr:        r.compilation.Stderr.String,
			Time:          r.compilation.Time.Float64,
			WallClockTime: r.compilation.WallClockTime.Float64,
			Memory:        int(r.compilation.Memory.Int64),
		},
		Scoring: model.Scoring{
			TaskScore:    r.scoring.TaskScore.Float64,
			ContestScore: r.scoring.ContestScore.Float64,
		},
		Evaluation: model.Evaluation{
			Tries: r.evaluation.Tries,
			Done:  r.evaluation.Done,
		},
	}
}

type sqlDraftResult struct {
	DatasetID int `db:"dataset_id"`
	EntryID   int `db:"entry_id"`
	compilation
	evaluation
	execution
}

func (r sqlDraftResult) toDraftResult() model.DraftResult {
	return model.DraftResult{
		EntryID:   r.EntryID,
		DatasetID: r.DatasetID,
		Compilation: model.Compilation{
			Status:        r.compilation.Status.String,
			Tries:         r.compilation.Tries,
			Stdout:        r.compilation.Stdout.String,
			Stderr:        r.compilation.Stderr.String,
			Time:          r.compilation.Time.Float64,
			WallClockTime: r.compilation.WallClockTime.Float64,
			Memory:        int(r.compilation.Memory.Int64),
		},
		Evaluation: model.Evaluation{
			Tries: r.evaluation.Tries,
			Done:  r.evaluation.Done,
		},
		Execution: model.Execution{
			Output:        r.execution.Output,
			Memory:        int(r.execution.Memory.Int64),
			WallClockTime: r.execution.WallClockTime.Float64,
			Time:          r.execution.Time.Float64,
		},
	}
}

type compilation struct {
	Status        sqlCMSCompilation `db:"compilation_status"`
	Tries         int               `db:"compilation_tries"`
	Stdout        sql.NullString    `db:"compilation_stdout"`
	Stderr        sql.NullString    `db:"compilation_stderr"`
	Time          sql.NullFloat64   `db:"compilation_time"`
	WallClockTime sql.NullFloat64   `db:"compilation_wall_clock_time"`
	Memory        sql.NullInt64     `db:"compilation_memory"`
}

type evaluation struct {
	Done  bool `db:"evaluation_done"`
	Tries int  `db:"evaluation_tries"`
}

type scoring struct {
	TaskScore    sql.NullFloat64 `db:"score"`
	ContestScore sql.NullFloat64 `db:"public_score"`
}

type execution struct {
	Time          sql.NullFloat64 `db:"execution_time"`
	WallClockTime sql.NullFloat64 `db:"execution_wall_clock_time"`
	Memory        sql.NullInt64   `db:"execution_memory"`
	Output        []byte          `db:"output_data"`
}

type sqlCMSLanguage struct {
	sql.NullString
}

// Scan implements the Scanner interface.
func (l *sqlCMSLanguage) Scan(value interface{}) error {
	err := l.NullString.Scan(value)

	if err != nil {
		return err
	}

	if !l.Valid {
		l.String = cmsLanguageDefault
	}

	return nil
}

type sqlCMSCompilation struct {
	sql.NullString
}

// Scan implements the Scanner interface.
func (l *sqlCMSCompilation) Scan(value interface{}) error {
	err := l.NullString.Scan(value)

	if err != nil {
		return err
	}

	if !l.Valid {
		l.String = cmsCompilationStatusDefault
	}

	return nil
}
