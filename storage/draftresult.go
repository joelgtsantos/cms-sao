package storage

import (
	"fmt"

	"github.com/jossemargt/cms-sao/model"
	"github.com/pkg/errors"
)

type DraftResultRepository interface {
	FindByID(entryID, datasetID int) (*model.DraftResult, error)
	FindBy(dto DraftResultDTO) ([]model.DraftResult, error)
}

func NewDraftResultRepository(db Queryer) DraftResultRepository {
	return &defaultDraftResultRepository{source: db}
}

type defaultDraftResultRepository struct {
	source Queryer
}

func (repo *defaultDraftResultRepository) FindByID(entryID, datasetID int) (*model.DraftResult, error) {
	result := model.DraftResult{}
	query, err := buildFindDraftResultByIDQuery()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed building Result SQL projection")
	}

	err = repo.source.Get(&result, query, entryID, datasetID)

	if err != nil {
		return nil, errors.Wrapf(err, "Failed query with Result ID %d-%d", entryID, datasetID)
	}

	return &result, nil
}

func (repo *defaultDraftResultRepository) FindBy(dto DraftResultDTO) ([]model.DraftResult, error) {
	query, err := buildFindDraftResultByQuery(dto)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed building Entry SQL projection")
	}

	results := make([]model.DraftResult, 0, dto.limit)
	err = repo.source.Select(&results, query)

	if err != nil {
		return nil, errors.Wrapf(err, "Failed Result query")
	}

	return results, nil
}

func buildFindDraftResultByIDQuery() (string, error) {
	return NewProjection(
		Select(
			"utr.dataset_id AS dataset_id",
			"utr.user_test_id AS entry_id",
			"utr.compilation_outcome AS compilation_status",
			"utr.compilation_tries",
			"utr.compilation_stdout",
			"utr.compilation_stderr",
			"utr.compilation_time",
			"utr.compilation_wall_clock_time",
			"utr.compilation_memory",
			"utr.evaluation_outcome IS NOT NULL AS evaluation_done",
			"utr.evaluation_tries",
			"utr.execution_time",
			"utr.execution_wall_clock_time",
			"utr.execution_memory",
			"lo.data AS output_data",
		),
		From(fmt.Sprintf("%s AS utr", entryDraftResultTable)),
		Join("%s AS fso ON fso.digest = utr.output", pgFsObjects),
		Join("%s AS lo ON lo.loid = fso.loid", pgLargeObject),
		Where("utr.user_test_id = $1"),
		Where("utr.dataset_id = $2"),
	)
}

func buildFindDraftResultByQuery(dto DraftResultDTO) (string, error) {
	dto.prepare()

	sqlParts := []SQLBuilderOption{
		Select(
			"utr.dataset_id AS dataset_id",
			"utr.user_test_id AS entry_id",
			"utr.compilation_outcome AS compilation_status",
			"utr.compilation_tries",
			"utr.compilation_stdout",
			"utr.compilation_stderr",
			"utr.compilation_time",
			"utr.compilation_wall_clock_time",
			"utr.compilation_memory",
			"utr.evaluation_outcome IS NOT NULL AS evaluation_done",
			"utr.evaluation_tries",
			"utr.execution_time",
			"utr.execution_wall_clock_time",
			"utr.execution_memory",
			"lo.data AS output_data",
		),
		From(fmt.Sprintf("%s AS utr", entryDraftResultTable)),
		Join("%s AS fso ON fso.digest = utr.output", pgFsObjects),
		Join("%s AS lo ON lo.loid = fso.loid", pgLargeObject),
		Join("%s AS sb ON sb.id = utr.user_test_id", entryDraftTable),
		Join("%s AS tsk ON tsk.id = sb.task_id", taskTable),
		Join("%s AS cts ON cts.id = tsk.contest_id", contestTable),
		OrderBy(fmt.Sprintf("utr.user_test_id %s, utr.dataset_id %s", dto.Order, dto.Order)),
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

	if dto.DraftID > 0 {
		sqlParts = append(sqlParts, Where(fmt.Sprintf("utr.user_test_id = %d", dto.DraftID)))
	}

	if dto.UserID > 0 {
		sqlParts = append(sqlParts,
			Join("%s AS prts ON prts.id = sb.participation_id", contestUserAsignationTable),
			Where(fmt.Sprintf("prts.user_id = %d", dto.UserID)),
		)
	}

	if dto.Page > 1 {
		sqlParts = append(sqlParts, Offset(dto.offset))
	}

	return NewProjection(sqlParts...)
}
