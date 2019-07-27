/**
 * @Author: huangw1
 * @Date: 2019/7/26 16:48
 */

package dao

import (
	"github.com/huangw1/bbs/database"
	"github.com/huangw1/bbs/model"
)

type CommentDao interface {
	Get(id int64) (*model.Comment, error)
	Take(where ...interface{}) (*model.Comment, error)
	QueryCondition(condition *model.Condition) ([]*model.Comment, error)
	QueryParams(params *model.Params) ([]*model.Comment, *model.Paging, error)
	Create(comment *model.Comment) error
	Update(comment *model.Comment) error
	Updates(id int64, columns map[string]interface{}) error
	UpdateColumn(id int64, name string, value interface{}) error
	Delete(id int64) (err error)
}

var CommentDaoImpl = NewCommentDao()

func NewCommentDao() CommentDao {
	return &commentDao{}
}

type commentDao struct {
}

func (dao *commentDao) Get(id int64) (*model.Comment, error) {
	item := new(model.Comment)
	err := database.GetDB().First(item, "id = ?", id).Error
	return item, err
}

func (dao *commentDao) Take(where ...interface{}) (*model.Comment, error) {
	item := new(model.Comment)
	err := database.GetDB().Take(item, where...).Error
	return item, err
}

func (dao *commentDao) QueryCondition(condition *model.Condition) ([]*model.Comment, error) {
	list := make([]*model.Comment, 0)
	err := DoQuery(database.GetDB(), condition).Find(&list).Error
	return list, err
}

func (dao *commentDao) QueryParams(params *model.Params) ([]*model.Comment, *model.Paging, error) {
	list := make([]*model.Comment, 0)
	err := StartQuery(database.GetDB(), params).Find(&list).Count(&params.Paging.Total).Error
	return list, params.Paging, err
}

func (dao *commentDao) Create(comment *model.Comment) error {
	return database.GetDB().Create(comment).Error
}

func (dao *commentDao) Update(comment *model.Comment) error {
	return database.GetDB().Save(comment).Error
}

func (dao *commentDao) Updates(id int64, columns map[string]interface{}) error {
	return database.GetDB().Model(&model.Comment{}).Where("id = ?", id).Updates(columns).Error
}

func (dao *commentDao) UpdateColumn(id int64, name string, value interface{}) error {
	return database.GetDB().Model(&model.Comment{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (dao *commentDao) Delete(id int64) (err error) {
	return database.GetDB().Model(&model.Comment{}).Delete("id", id).Error
}
