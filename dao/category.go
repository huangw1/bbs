/**
 * @Author: huangw1
 * @Date: 2019/7/26 16:48
 */

package dao

import (
	"github.com/huangw1/bbs/database"
	"github.com/huangw1/bbs/model"
)

type CategoryDao interface {
	Get(id int64) (*model.Category, error)
	Take(where ...interface{}) (*model.Category, error)
	QueryCondition(condition *model.Condition) ([]*model.Category, error)
	QueryParams(params *model.Params) ([]*model.Category, *model.Paging, error)
	Create(category *model.Category) error
	Update(category *model.Category) error
	Updates(id int64, columns map[string]interface{}) error
	UpdateColumn(id int64, name string, value interface{}) error
	Delete(id int64) (err error)

	GetCategories() ([]*model.Category, error)
}

var CategoryDaoImpl = NewCategoryDao()

func NewCategoryDao() CategoryDao {
	return &categoryDao{}
}

type categoryDao struct {
}

func (dao *categoryDao) Get(id int64) (*model.Category, error) {
	item := new(model.Category)
	err := database.GetDB().First(item, "id = ?", id).Error
	return item, err
}

func (dao *categoryDao) Take(where ...interface{}) (*model.Category, error) {
	item := new(model.Category)
	err := database.GetDB().Take(item, where...).Error
	return item, err
}

func (dao *categoryDao) QueryCondition(condition *model.Condition) ([]*model.Category, error) {
	list := make([]*model.Category, 0)
	err := DoQuery(database.GetDB(), condition).Find(&list).Error
	return list, err
}

func (dao *categoryDao) QueryParams(params *model.Params) ([]*model.Category, *model.Paging, error) {
	list := make([]*model.Category, 0)
	err := StartQuery(database.GetDB(), params).Find(&list).Count(&params.Paging.Total).Error
	return list, params.Paging, err
}

func (dao *categoryDao) Create(category *model.Category) error {
	return database.GetDB().Create(category).Error
}

func (dao *categoryDao) Update(category *model.Category) error {
	return database.GetDB().Save(category).Error
}

func (dao *categoryDao) Updates(id int64, columns map[string]interface{}) error {
	return database.GetDB().Model(&model.Category{}).Where("id = ?", id).Updates(columns).Error
}

func (dao *categoryDao) UpdateColumn(id int64, name string, value interface{}) error {
	return database.GetDB().Model(&model.Category{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (dao *categoryDao) Delete(id int64) (err error) {
	return database.GetDB().Model(&model.Category{}).Delete("id", id).Error
}

func (dao *categoryDao) GetCategories() ([]*model.Category, error) {
	return dao.QueryCondition(model.NewCondition("status = ?", model.CategoryStatusOk))
}
