/**
 * @Author: huangw1
 * @Date: 2019/7/25 11:32
 */

package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type queryPair struct {
	Query string
	Args  []interface{}
}

type orderPair struct {
	Column string
	Sort   string
}

type Params struct {
	QueryPairs []queryPair
	OrderPairs []orderPair
	Paging     *Paging
}

func NewParams() *Params {
	return &Params{}
}

func (q *Params) Where(query string, args ...interface{}) *Params {
	q.QueryPairs = append(q.QueryPairs, queryPair{Query: query, Args: args})
	return q
}

func (q *Params) Eq(column string, args ...interface{}) *Params {
	q.Where(fmt.Sprintf("%s = ?", column), args)
	return q
}

func (q *Params) NotEq(column string, args ...interface{}) *Params {
	q.Where(fmt.Sprintf("%s <> ?", column), args)
	return q
}

func (q *Params) Gt(column string, args ...interface{}) *Params {
	q.Where(fmt.Sprintf("%s > ?", column), args)
	return q
}

func (q *Params) Gte(column string, args ...interface{}) *Params {
	q.Where(fmt.Sprintf("%s >= ?", column), args)
	return q
}

func (q *Params) Lt(column string, args ...interface{}) *Params {
	q.Where(fmt.Sprintf("%s < ?", column), args)
	return q
}

func (q *Params) Lte(column string, args ...interface{}) *Params {
	q.Where(fmt.Sprintf("%s <= ?", column), args)
	return q
}

func (q *Params) Like(column string, arg string) *Params {
	q.Where(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", arg))
	return q
}

func (q *Params) OrderBy(column string, sort string) *Params {
	q.OrderPairs = append(q.OrderPairs, orderPair{Column: column, Sort: sort})
	return q
}

func (q *Params) Asc(column string) *Params {
	return q.OrderBy(column, "asc")
}

func (q *Params) Desc(column string) *Params {
	return q.OrderBy(column, "desc")
}

func (q *Params) Page(page, limit int) *Params {
	if q.Paging == nil {
		q.Paging = &Paging{Page: page, Limit: limit}
	} else {
		q.Paging.Page = page
		q.Paging.Limit = limit
	}
	return q
}

func (q *Params) limit(limit int) *Params {
	q.Page(1, limit)
	return q
}

func (q *Params) StartQuery(db *gorm.DB) *gorm.DB {
	result := db
	if len(q.QueryPairs) > 0 {
		for _, pair := range q.QueryPairs {
			result = result.Where(pair.Query, pair.Args...)
		}
	}
	if len(q.OrderPairs) > 0 {
		for _, pair := range q.OrderPairs {
			result = result.Order(fmt.Sprintf("%s %s", pair.Column, pair.Sort))
		}
	}
	if q.Paging != nil && q.Paging.Limit > 0 {
		result = result.Limit(q.Paging.Limit)
	}

	if q.Paging != nil && q.Paging.Offset() > 0 {
		result = result.Offset(q.Paging.Offset())
	}
	return result
}

func (q *Params) StartCount(db *gorm.DB) *gorm.DB {
	result := db
	if len(q.QueryPairs) > 0 {
		for _, pair := range q.QueryPairs {
			result = result.Where(pair.Query, pair.Args...)
		}
	}
	return result
}
