package pagination

import "github.com/doug-martin/goqu/v9"

type Pagination struct {
	page  uint
	limit uint
}

func NewPagination(page, limit uint) *Pagination {
	return &Pagination{
		page:  page,
		limit: limit,
	}
}
func (p *Pagination) Calc() (offset uint, limit uint) {
	limit = p.limit
	if p.page > 1 {
		offset = (p.limit * p.page) - p.limit
	}
	return
}

func (p *Pagination) DataSet(ds *goqu.SelectDataset) *goqu.SelectDataset {
	offset, limit := p.Calc()
	ds = ds.Limit(limit)
	if offset > 0 {
		ds = ds.Offset(offset)
	}
	return ds
}
