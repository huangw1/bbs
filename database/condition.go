/**
 * @Author: huangw1
 * @Date: 2019/7/25 11:32
 */

package database

import "github.com/jinzhu/gorm"

type Condition struct {
	Query  string
	Args   []interface{}
	Orders []string
	Limit  int
	Offset int
}

func NewCondition(query string, args ...interface{}) *Condition {
	return &Condition{
		Query: query,
		Args:  args,
	}
}

func (q *Condition) Order(order string) *Condition {
	q.Orders = append(q.Orders, order)
	return q
}

func (q *Condition) Size(size int) *Condition {
	q.Limit = size
	return q
}

func (q *Condition) Page(page int, size int) *Condition {
	p := Paging{Page: page, Limit: size}
	q.Limit = p.Limit
	q.Offset = p.Offset()
	return q
}

func (q *Condition) DoQuery(db *gorm.DB) *gorm.DB {
	result := db.Where(q.Query, q.Args...)
	if q.Limit > 0 {
		result = result.Limit(q.Limit)
	}
	if q.Offset > 0 {
		result = result.Offset(q.Offset)
	}
	if len(q.Orders) > 0 {
		for _, order := range q.Orders {
			result = result.Order(order)
		}
	}
	return result
}
