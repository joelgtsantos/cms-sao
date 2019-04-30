package storage

import (
	"fmt"

	"github.com/jossemargt/cms-sao/model"
	"github.com/pkg/errors"
)

type EntryRepository interface {
	FindByID(int) (*model.Entry, error)
	FindBy(EntryDTO) ([]model.Entry, error)
}

// NewEntryDraftRepository returns a EntryRepository for Entry (CMS submittion) entity
func NewEntryRepository(dbx Queryer) EntryRepository {
	return &defaultEntryRepository{
		source:        dbx,
		findByIDQuery: buildEntryFindByIDQuery,
		findByQuery:   buildEntryFindByQuery,
	}
}

type defaultEntryRepository struct {
	source        Queryer
	findByIDQuery func() (string, error)
	findByQuery   func(dto EntryDTO) (string, error)
}

func (entryRepo *defaultEntryRepository) FindByID(entryID int) (*model.Entry, error) {
	entry := sqlEntry{}

	query, err := entryRepo.findByIDQuery()

	if err != nil {
		return nil, errors.Wrapf(err, "Failed building Entry SQL projection")
	}

	err = entryRepo.source.Get(&entry, query, entryID)

	if err != nil {
		return nil, errors.Wrapf(err, "Failed query with Entry ID %d", entryID)
	}

	rEntry := entry.toEntry()
	return &rEntry, nil
}

func (entryRepo *defaultEntryRepository) FindBy(dto EntryDTO) ([]model.Entry, error) {
	query, err := entryRepo.findByQuery(dto)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed building Entry SQL projection")
	}

	nullableEntries := make([]sqlEntry, 0, dto.limit)
	err = entryRepo.source.Select(&nullableEntries, query)

	if err != nil {
		return nil, errors.Wrapf(err, "Failed Entries query")
	}

	entries := make([]model.Entry, len(nullableEntries))
	for i, entry := range nullableEntries {
		entries[i] = entry.toEntry()
	}

	return entries, nil
}

func buildEntryFindByIDQuery() (string, error) {
	return NewProjection(
		Select(
			"sb.id",
			"sb.task_id",
			"tsk.name AS task_slug",
			"sbr.dataset_id AS result_prtl_id",
			"tsk.contest_id",
			"cts.name AS contest_slug",
			"sb.language",
			"tkn.id IS NOT NULL AS token",
		),
		From(fmt.Sprintf("%s AS sb", entryTable)),
		Join("%s AS tsk ON tsk.id = sb.task_id", taskTable),
		Join("%s AS cts ON cts.id = tsk.contest_id", contestTable),
		LeftJoin("%s AS sbr ON sbr.submission_id = sb.id", resultTable),
		LeftJoin("%s AS tkn ON sb.id = tkn.submission_id", tokenTable),
		Where("sb.id = $1"),
	)
}

func buildEntryFindByQuery(dto EntryDTO) (string, error) {
	dto.prepare()

	sqlParts := []SQLBuilderOption{
		Select(
			"sb.id",
			"sb.task_id",
			"tsk.name AS task_slug",
			"sbr.dataset_id AS result_prtl_id",
			"tsk.contest_id",
			"cts.name AS contest_slug",
			"sb.language",
			"tkn.id IS NOT NULL AS token",
		),
		From(fmt.Sprintf("%s AS sb", entryTable)),
		Join("%s AS tsk ON tsk.id = sb.task_id", taskTable),
		Join("%s AS cts ON cts.id = tsk.contest_id", contestTable),
		LeftJoin("%s AS sbr ON sbr.submission_id = sb.id", resultTable),
		LeftJoin("%s AS tkn ON sb.id = tkn.submission_id", tokenTable),
		OrderBy(fmt.Sprintf("id %s", dto.Order)),
		Limit(dto.limit),
	}

	if dto.TaskSlug != "" {
		sqlParts = append(sqlParts, Where(fmt.Sprintf("tsk.name LIKE '%s'", dto.TaskSlug)))
	}

	if dto.ContestSlug != "" {
		sqlParts = append(sqlParts, Where(fmt.Sprintf("cts.name LIKE '%s'", dto.ContestSlug)))
	}

	if dto.TaskID > 0 {
		sqlParts = append(sqlParts, Where(fmt.Sprintf("tsk.id = %d", dto.TaskID)))
	}

	if dto.ContestID > 0 {
		sqlParts = append(sqlParts, Where(fmt.Sprintf("cts.id = %d", dto.ContestID)))
	}

	if dto.Page > 1 {
		sqlParts = append(sqlParts, Offset(dto.offset))
	}

	return NewProjection(sqlParts...)
}
