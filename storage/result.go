package storage

import (
	"fmt"

	"github.com/jossemargt/cms-sao/model"
	"github.com/pkg/errors"
)

type ResultRepository interface {
	FindByID(entryID, datasetID int) (*model.Result, error)
	FindBy(dto ResultDTO) ([]model.Result, error)
}

func NewResultRepository(db Queryer) ResultRepository {
	return &defaultResultRepository{source: db}
}

type defaultResultRepository struct {
	source Queryer
}

func (repo *defaultResultRepository) FindByID(entryID, datasetID int) (*model.Result, error) {
	result := model.Result{}
	query, err := buildFindResultByIDQuery()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed building Result SQL projection")
	}

	err = repo.source.Get(&result, query, entryID, datasetID)

	if err != nil {
		return nil, errors.Wrapf(err, "Failed query with Result ID %d-%d", entryID, datasetID)
	}

	return &result, nil
}

func (repo *defaultResultRepository) FindBy(dto ResultDTO) ([]model.Result, error) {
	query, err := buildFindResultByQuery(dto)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed building Entry SQL projection")
	}

	results := make([]model.Result, 0, dto.limit)
	err = repo.source.Select(&results, query)

	if err != nil {
		return nil, errors.Wrapf(err, "Failed Result query")
	}

	return results, nil
}

func buildFindResultByIDQuery() (string, error) {
	return NewProjection(
		Select(
			"sr.dataset_id AS dataset_id",
			"sr.submission_id AS entry_id",
			"sr.compilation_outcome AS compilation_status",
			"sr.compilation_tries",
			"sr.compilation_stdout",
			"sr.compilation_stderr",
			"sr.compilation_time",
			"sr.compilation_wall_clock_time",
			"sr.compilation_memory",
			"sr.evaluation_outcome IS NOT NULL AS evaluation_done",
			"sr.evaluation_tries",
			"sr.score",
			"sr.public_score",
		),
		From(fmt.Sprintf("%s AS sr", resultTable)),
		Where("sr.submission_id = $1"),
		Where("sr.dataset_id = $2"),
	)
}

func buildFindResultByQuery(dto ResultDTO) (string, error) {
	dto.prepare()

	sqlParts := []SQLBuilderOption{
		Select(
			"sr.dataset_id AS dataset_id",
			"sr.submission_id AS entry_id",
			"sr.compilation_outcome AS compilation_status",
			"sr.compilation_tries",
			"sr.compilation_stdout",
			"sr.compilation_stderr",
			"sr.compilation_time",
			"sr.compilation_wall_clock_time",
			"sr.compilation_memory",
			"sr.evaluation_outcome IS NOT NULL AS evaluation_done",
			"sr.evaluation_tries",
			"sr.score",
			"sr.public_score",
		),
		From(fmt.Sprintf("%s AS sr", resultTable)),
		Join("%s AS sb ON sb.id = sr.submission_id", entryTable),
		Join("%s AS tsk ON tsk.id = sb.task_id", taskTable),
		Join("%s AS cts ON cts.id = tsk.contest_id", contestTable),
		OrderBy(fmt.Sprintf("sr.submission_id %s, sr.dataset_id %s", dto.Order, dto.Order)),
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

	if dto.EntryID > 0 {
		sqlParts = append(sqlParts, Where(fmt.Sprintf("sr.submission_id = %d", dto.EntryID)))
	}

	if dto.UserID > 0 {
		sqlParts = append(sqlParts,
			Join("%s AS prts ON prts.id = sb.participation_id", contestUserAssignationTable),
			Where(fmt.Sprintf("prts.user_id = %d", dto.UserID)),
		)
	}

	if dto.MaxScore {
		sqlParts = append(sqlParts,
			LeftJoin(
				"%s AS sr2 "+
					"ON sr.dataset_id = sr2.dataset_id "+
					"AND sr.score < sr2.score", resultTable),
			Where("sr2.score IS NULL"),
		)
	}

	if dto.Page > 1 {
		sqlParts = append(sqlParts, Offset(dto.offset))
	}

	return NewProjection(sqlParts...)
}
