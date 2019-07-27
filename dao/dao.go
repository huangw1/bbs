/**
 * @Author: huangw1
 * @Date: 2019/7/29 15:56
 */

package dao

import (
	"fmt"
	"github.com/huangw1/bbs/model"
	"github.com/jinzhu/gorm"
)

func DoQuery(db *gorm.DB, q *model.Condition) *gorm.DB {
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

func StartQuery(db *gorm.DB, q *model.Params) *gorm.DB {
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

func StartCount(db *gorm.DB, q *model.Params) *gorm.DB {
	result := db
	if len(q.QueryPairs) > 0 {
		for _, pair := range q.QueryPairs {
			result = result.Where(pair.Query, pair.Args...)
		}
	}
	return result
}
