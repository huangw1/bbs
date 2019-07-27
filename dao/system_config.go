/**
 * @Author: huangw1
 * @Date: 2019/7/26 16:48
 */

package dao

import (
	"github.com/huangw1/bbs/database"
	"github.com/huangw1/bbs/model"
)

type SystemConfigDao interface {
	Get(id int64) (*model.SystemConfig, error)
	Take(where ...interface{}) (*model.SystemConfig, error)
	QueryCondition(condition *model.Condition) ([]*model.SystemConfig, error)
	QueryParams(params *model.Params) ([]*model.SystemConfig, *model.Paging, error)
	Create(systemConfig *model.SystemConfig) error
	Update(systemConfig *model.SystemConfig) error
	Updates(id int64, columns map[string]interface{}) error
	UpdateColumn(id int64, name string, value interface{}) error
	Delete(id int64) (err error)

	GetByKey(key string) (*model.SystemConfig, error)
}

var SystemConfigDaoImpl = NewSystemConfigDao()

func NewSystemConfigDao() SystemConfigDao {
	return &systemConfigDao{}
}

type systemConfigDao struct {
}

func (dao *systemConfigDao) Get(id int64) (*model.SystemConfig, error) {
	item := new(model.SystemConfig)
	err := database.GetDB().First(item, "id = ?", id).Error
	return item, err
}

func (dao *systemConfigDao) Take(where ...interface{}) (*model.SystemConfig, error) {
	item := new(model.SystemConfig)
	err := database.GetDB().Take(item, where...).Error
	return item, err
}

func (dao *systemConfigDao) QueryCondition(condition *model.Condition) ([]*model.SystemConfig, error) {
	list := make([]*model.SystemConfig, 0)
	err := DoQuery(database.GetDB(), condition).Find(&list).Error
	return list, err
}

func (dao *systemConfigDao) QueryParams(params *model.Params) ([]*model.SystemConfig, *model.Paging, error) {
	list := make([]*model.SystemConfig, 0)
	err := StartQuery(database.GetDB(), params).Find(&list).Count(&params.Paging.Total).Error
	return list, params.Paging, err
}

func (dao *systemConfigDao) Create(systemConfig *model.SystemConfig) error {
	return database.GetDB().Create(systemConfig).Error
}

func (dao *systemConfigDao) Update(systemConfig *model.SystemConfig) error {
	return database.GetDB().Save(systemConfig).Error
}

func (dao *systemConfigDao) Updates(id int64, columns map[string]interface{}) error {
	return database.GetDB().Model(&model.SystemConfig{}).Where("id = ?", id).Updates(columns).Error
}

func (dao *systemConfigDao) UpdateColumn(id int64, name string, value interface{}) error {
	return database.GetDB().Model(&model.SystemConfig{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (dao *systemConfigDao) Delete(id int64) (err error) {
	return database.GetDB().Model(&model.SystemConfig{}).Delete("id", id).Error
}

func (dao *systemConfigDao) GetByKey(key string) (*model.SystemConfig, error) {
	return dao.Take("key = ?", key)
}
