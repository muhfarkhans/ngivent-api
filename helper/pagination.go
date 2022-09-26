package helper

import (
	"math"
	"regexp"

	"gorm.io/gorm"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^\p{L}\p{N} ]+`)

func clearString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

type Pagination struct {
	Limit      int    `json:"limit" query:"limit" form:"limit"`
	Page       int    `json:"page" query:"page" form:"page"`
	Sort       string `json:"sort" query:"sort" form:"sort"`
	TotalRows  int64  `json:"total_rows"`
	TotalPages int    `json:"total_pages"`
	UserId     int
	Q          string `json:"q" form:"q"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetQ() string {
	return clearString(p.Q)
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = clearString("Id desc'")
	}
	return p.Sort
}

func Paginate(value interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	d := db.Model(value)
	if pagination.UserId != 0 {
		d = d.Where("user_id=?", pagination.UserId)
	}
	d.Count(&totalRows)

	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}
