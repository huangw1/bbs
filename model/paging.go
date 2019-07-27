/**
 * @Author: huangw1
 * @Date: 2019/7/29 15:59
 */

package model

type Paging struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

func (p *Paging) Offset() int {
	offset := 0
	if p.Page > 0 {
		offset = (p.Page - 1) * p.Limit
	}
	return offset
}

func (p *Paging) TotalPage() int {
	if p.Total == 0 || p.Limit == 0 {
		return 0
	}
	totalPage := p.Total / p.Limit
	if p.Total%p.Limit > 0 {
		totalPage += 1
	}
	return totalPage
}

type PageResult struct {
	Page    *Paging     `json:"page"`
	Results interface{} `json:"results"`
}
