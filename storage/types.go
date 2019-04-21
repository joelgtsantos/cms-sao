package storage

type DTO struct {
	limit int
	offset int
	Page int
	PageSize int
	Order string
}

func (d *DTO) prepare() {
	if d.PageSize < 1 {
		d.PageSize = 10
	}

	if d.Page < 1 {
		d.Page = 1
	}

	d.limit = d.PageSize
	d.offset = d.Page * d.PageSize - d.PageSize

	if len(d.Order) == 0 || !validOrder(d.Order) {
		d.Order = "DESC"
	}

}

func validOrder(order string) bool {
	return order == "ASC" || order == "DESC"
}