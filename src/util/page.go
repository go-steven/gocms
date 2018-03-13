package util

type Pager struct {
	Page     int
	PageSize int
	Offset   int
}

func NewPager(page, pageSize int) Pager {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 15
	}
	return Pager{
		Page:     page,
		PageSize: pageSize,
		Offset:   (page - 1) * pageSize,
	}
}
