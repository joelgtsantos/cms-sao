package storage

import (
	"fmt"

	"github.com/jossemargt/cms-sao/model"
	"github.com/pkg/errors"
)

const entryTable = "submissions"
const taskTable = "tasks"
const contestTable = "contests"
const resultTable = "submission_results"

type EntryRepository interface {
	FindByID(int) (*model.Entry, error)
	FindBy(EntryDTO) ([]model.Entry, error)
}

func NewEntryRepository(dbx Queryer) EntryRepository {
	return &defaultEntryRepository{dbx}
}

type defaultEntryRepository struct {
	source Queryer
}

func (entryRepo *defaultEntryRepository) FindByID(entryID int) (*model.Entry, error) {
	entry := model.Entry{}

	query, err := NewProjection(
		Select(
			"sb.id",
			"sb.task_id",
			"tsk.name AS task_slug",
			"sbr.dataset_id AS result_prtl_id",
			"tsk.contest_id",
			"cts.name AS contest_slug",
		),
		From(fmt.Sprintf("%s AS sb", entryTable)),
		Join("%s AS tsk ON tsk.id = sb.task_id", taskTable),
		Join("%s AS cts ON cts.id = tsk.contest_id", contestTable),
		Join("%s AS sbr ON sbr.submission_id = sb.id", resultTable),
		Where("sb.id = $1"),
	)

	if err != nil {
		return nil, errors.Wrapf(err, "Failed building Entry SQL projection")
	}

	err = entryRepo.source.Get(&entry, query, entryID)

	if err != nil {
		return nil, errors.Wrapf(err, "Failed query with Entry ID %d", entryID)
	}

	return &entry, nil
}

func (entryRepo *defaultEntryRepository) FindBy(dto EntryDTO) ([]model.Entry, error) {
	query, err := buildEntryQuery(dto)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed building Entry SQL projection")
	}

	entries := make([]model.Entry, dto.limit)
	err = entryRepo.source.Select(&entries, query)

	if err != nil {
		return nil, errors.Wrapf(err, "Failed Entries query")
	}

	return entries, nil
}

func buildEntryQuery(dto EntryDTO) (string, error) {
	dto.prepare()

	sqlParts := []SQLBuilderOption{
		Select(
			"sb.id",
			"sb.task_id",
			"tsk.name AS task_slug",
			"sbr.dataset_id AS result_prtl_id",
			"tsk.contest_id",
			"cts.name AS contest_slug",
		),
		From(fmt.Sprintf("%s AS sb", entryTable)),
		Join("%s AS tsk ON tsk.id = sb.task_id", taskTable),
		Join("%s AS cts ON cts.id = tsk.contest_id", contestTable),
		Join("%s AS sbr ON sbr.submission_id = sb.id", resultTable),
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

type EntryDTO struct {
	DTO
	ContestID   int
	ContestSlug string
	TaskID      int
	TaskSlug    string
}
