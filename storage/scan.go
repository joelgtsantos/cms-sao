package storage

import "database/sql"

const (
	cmsLanguageDefault = "none"
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
