package model

import "math"

type Pagination struct {
	Limit        int   `json:"limit"`
	Page         int   `json:"page"`
	TotalPage    int   `json:"total_page"`
	TotalRecords int64 `json:"total_records"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit <= 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) TotalToPage(total int64) *Pagination {
	p.TotalRecords = int64(total)
	p.TotalPage = int(math.Ceil(float64(p.TotalRecords) / float64(p.GetLimit())))
	return p
}

type Cursor struct {
	TotalRecords int64  `json:"total_records"`
	Limit        int    `json:"limit"`
	Next         string `json:"next"`
	Prev         string `json:"prev"`
}
