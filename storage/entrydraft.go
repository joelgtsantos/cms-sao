package storage

import (
	"fmt"
)

// NewEntryDraftRepository returns a EntryRepository for Entry Draft (CMS user_test) entity
func NewEntryDraftRepository(dbx Queryer) EntryRepository {
	return &defaultEntryRepository{
		source:        dbx,
		findByIDQuery: buildEntryDraftFindByIDQuery,
		findByQuery:   buildEntryDraftFindQuery,
	}
}

func buildEntryDraftFindByIDQuery() (string, error) {
	return NewProjection(
		Select(
			"ut.id",
			"ut.task_id",
			"tsk.name AS task_slug",
			"utr.dataset_id AS result_prtl_id",
			"tsk.contest_id",
			"cts.name AS contest_slug",
		),
		From(fmt.Sprintf("%s AS ut", entryDraftTable)),
		Join("%s AS tsk ON tsk.id = ut.task_id", taskTable),
		Join("%s AS cts ON cts.id = tsk.contest_id", contestTable),
		Join("%s AS utr ON utr.user_test_id = ut.id", entryDraftResultTable),
		Where("ut.id = $1"),
	)
}

func buildEntryDraftFindQuery(dto EntryDTO) (string, error) {
	dto.prepare()

	sqlParts := []SQLBuilderOption{
		Select(
			"ut.id",
			"ut.task_id",
			"tsk.name AS task_slug",
			"utr.dataset_id AS result_prtl_id",
			"tsk.contest_id",
			"cts.name AS contest_slug",
		),
		From(fmt.Sprintf("%s AS ut", entryDraftTable)),
		Join("%s AS tsk ON tsk.id = ut.task_id", taskTable),
		Join("%s AS cts ON cts.id = tsk.contest_id", contestTable),
		Join("%s AS utr ON utr.user_test_id = ut.id", entryDraftResultTable),
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
