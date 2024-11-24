package common

import "gorm.io/gorm"

type Paginator struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type PagingInfo struct {
	Total int64 `json:"total"`
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func NewPaginator(page, limit int) *Paginator {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	return &Paginator{
		Page:  page,
		Limit: limit,
	}
}

func (p *Paginator) ApplyPagination(db *gorm.DB) *gorm.DB {
	offset := (p.Page - 1) * p.Limit
	return db.Offset(offset).Limit(p.Limit)
}
