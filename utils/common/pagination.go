package common

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Paginator struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type PagingInfo struct {
	Total int64 `json:"total"`
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
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

func GetLimitAndPage(ctx *gin.Context) (int, int, error) {
	// Ambil nilai page dan limit dari query parameter, dengan default value
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "10")

	// Konversi page dari string ke integer
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		return 0, 0, fmt.Errorf("invalid page value: %s", pageStr)
	}

	// Konversi limit dari string ke integer
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		return 0, 0, fmt.Errorf("invalid limit value: %s", limitStr)
	}

	return page, limit, nil
}
