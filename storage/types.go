package storage

const (
	entryTable                  = "submissions"
	taskTable                   = "tasks"
	contestTable                = "contests"
	resultTable                 = "submission_results"
	entryDraftTable             = "user_tests"
	entryDraftResultTable       = "user_test_results"
	tokenTable                  = "tokens"
	contestUserAssignationTable = "participations"

	pgFsObjects   = "fsobjects"
	pgLargeObject = "pg_largeobject"
)

type DTO struct {
	limit    int
	offset   int
	Page     int
	PageSize int
	Order    string
}

func (d *DTO) prepare() {
	if d.PageSize < 1 {
		d.PageSize = 10
	}

	if d.Page < 1 {
		d.Page = 1
	}

	d.limit = d.PageSize
	d.offset = d.Page*d.PageSize - d.PageSize

	if len(d.Order) == 0 || !validOrder(d.Order) {
		d.Order = "DESC"
	}

}

func validOrder(order string) bool {
	return order == "ASC" || order == "DESC"
}

type EntryDTO struct {
	DTO
	ContestID   int
	ContestSlug string
	TaskID      int
	TaskSlug    string
}

type ResultDTO struct {
	DTO
	ContestID   int
	ContestSlug string
	TaskID      int
	TaskSlug    string
	EntryID     int
	UserID      int
	MaxScore    bool
}

type DraftResultDTO struct {
	DTO
	ContestID   int
	ContestSlug string
	TaskID      int
	TaskSlug    string
	DraftID     int
	UserID      int
}
