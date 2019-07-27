/**
 * @Author: huangw1
 * @Date: 2019/7/26 16:48
 */

package dao

import (
	"github.com/huangw1/bbs/database"
	"github.com/huangw1/bbs/model"
)

type MessageDao interface {
	Get(id int64) (*model.Message, error)
	Take(where ...interface{}) (*model.Message, error)
	QueryCondition(condition *model.Condition) ([]*model.Message, error)
	QueryParams(params *model.Params) ([]*model.Message, *model.Paging, error)
	Create(message *model.Message) error
	Update(message *model.Message) error
	Updates(id int64, columns map[string]interface{}) error
	UpdateColumn(id int64, name string, value interface{}) error
	Delete(id int64) (err error)
}

var MessageDaoImpl = NewMessageDao()

func NewMessageDao() MessageDao {
	return &messageDao{}
}

type messageDao struct {
}

func (dao *messageDao) Get(id int64) (*model.Message, error) {
	item := new(model.Message)
	err := database.GetDB().First(item, "id = ?", id).Error
	return item, err
}

func (dao *messageDao) Take(where ...interface{}) (*model.Message, error) {
	item := new(model.Message)
	err := database.GetDB().Take(item, where...).Error
	return item, err
}

func (dao *messageDao) QueryCondition(condition *model.Condition) ([]*model.Message, error) {
	list := make([]*model.Message, 0)
	err := DoQuery(database.GetDB(), condition).Find(&list).Error
	return list, err
}

func (dao *messageDao) QueryParams(params *model.Params) ([]*model.Message, *model.Paging, error) {
	list := make([]*model.Message, 0)
	err := StartQuery(database.GetDB(), params).Find(&list).Count(&params.Paging.Total).Error
	return list, params.Paging, err
}

func (dao *messageDao) Create(message *model.Message) error {
	return database.GetDB().Create(message).Error
}

func (dao *messageDao) Update(message *model.Message) error {
	return database.GetDB().Save(message).Error
}

func (dao *messageDao) Updates(id int64, columns map[string]interface{}) error {
	return database.GetDB().Model(&model.Message{}).Where("id = ?", id).Updates(columns).Error
}

func (dao *messageDao) UpdateColumn(id int64, name string, value interface{}) error {
	return database.GetDB().Model(&model.Message{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (dao *messageDao) Delete(id int64) (err error) {
	return database.GetDB().Model(&model.Message{}).Delete("id", id).Error
}
