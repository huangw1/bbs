/**
 * @Author: huangw1
 * @Date: 2019/7/26 16:48
 */

package dao

import (
	"github.com/huangw1/bbs/database"
	"github.com/huangw1/bbs/model"
)

type ThirdUserDao interface {
	Get(id int64) (*model.ThirdUser, error)
	Take(where ...interface{}) (*model.ThirdUser, error)
	QueryCondition(condition *model.Condition) ([]*model.ThirdUser, error)
	QueryParams(params *model.Params) ([]*model.ThirdUser, *model.Paging, error)
	Create(thirdUser *model.ThirdUser) error
	Update(thirdUser *model.ThirdUser) error
	Updates(id int64, columns map[string]interface{}) error
	UpdateColumn(id int64, name string, value interface{}) error
	Delete(id int64) (err error)

	GetByThirdId(thirdId int64) (*model.ThirdUser, error)
}

var ThirdUserDaoImpl = NewThirdUserDao()

func NewThirdUserDao() ThirdUserDao {
	return &thirdUserDao{}
}

type thirdUserDao struct {
}

func (dao *thirdUserDao) Get(id int64) (*model.ThirdUser, error) {
	item := new(model.ThirdUser)
	err := database.GetDB().First(item, "id = ?", id).Error
	return item, err
}

func (dao *thirdUserDao) Take(where ...interface{}) (*model.ThirdUser, error) {
	item := new(model.ThirdUser)
	err := database.GetDB().Take(item, where...).Error
	return item, err
}

func (dao *thirdUserDao) QueryCondition(condition *model.Condition) ([]*model.ThirdUser, error) {
	list := make([]*model.ThirdUser, 0)
	err := DoQuery(database.GetDB(), condition).Find(&list).Error
	return list, err
}

func (dao *thirdUserDao) QueryParams(params *model.Params) ([]*model.ThirdUser, *model.Paging, error) {
	list := make([]*model.ThirdUser, 0)
	err := StartQuery(database.GetDB(), params).Find(&list).Count(&params.Paging.Total).Error
	return list, params.Paging, err
}

func (dao *thirdUserDao) Create(thirdUser *model.ThirdUser) error {
	return database.GetDB().Create(thirdUser).Error
}

func (dao *thirdUserDao) Update(thirdUser *model.ThirdUser) error {
	return database.GetDB().Save(thirdUser).Error
}

func (dao *thirdUserDao) Updates(id int64, columns map[string]interface{}) error {
	return database.GetDB().Model(&model.ThirdUser{}).Where("id = ?", id).Updates(columns).Error
}

func (dao *thirdUserDao) UpdateColumn(id int64, name string, value interface{}) error {
	return database.GetDB().Model(&model.ThirdUser{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (dao *thirdUserDao) Delete(id int64) (err error) {
	return database.GetDB().Model(&model.ThirdUser{}).Delete("id", id).Error
}

func (dao *thirdUserDao) GetByThirdId(thirdId int64) (*model.ThirdUser, error) {
	return dao.Take("thirdId = ?", thirdId)
}
