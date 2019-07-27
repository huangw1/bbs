/**
 * @Author: huangw1
 * @Date: 2019/7/26 16:48
 */

package dao

import (
	"github.com/huangw1/bbs/database"
	"github.com/huangw1/bbs/model"
)

type TopicDao interface {
	Get(id int64) (*model.Topic, error)
	Take(where ...interface{}) (*model.Topic, error)
	QueryCondition(condition *model.Condition) ([]*model.Topic, error)
	QueryParams(params *model.Params) ([]*model.Topic, *model.Paging, error)
	Create(topic *model.Topic) error
	Update(topic *model.Topic) error
	Updates(id int64, columns map[string]interface{}) error
	UpdateColumn(id int64, name string, value interface{}) error
	Delete(id int64) (err error)
}

var TopicDaoImpl = NewTopicDao()

func NewTopicDao() TopicDao {
	return &topicDao{}
}

type topicDao struct {
}

func (dao *topicDao) Get(id int64) (*model.Topic, error) {
	item := new(model.Topic)
	err := database.GetDB().First(item, "id = ?", id).Error
	return item, err
}

func (dao *topicDao) Take(where ...interface{}) (*model.Topic, error) {
	item := new(model.Topic)
	err := database.GetDB().Take(item, where...).Error
	return item, err
}

func (dao *topicDao) QueryCondition(condition *model.Condition) ([]*model.Topic, error) {
	list := make([]*model.Topic, 0)
	err := DoQuery(database.GetDB(), condition).Find(&list).Error
	return list, err
}

func (dao *topicDao) QueryParams(params *model.Params) ([]*model.Topic, *model.Paging, error) {
	list := make([]*model.Topic, 0)
	err := StartQuery(database.GetDB(), params).Find(&list).Count(&params.Paging.Total).Error
	return list, params.Paging, err
}

func (dao *topicDao) Create(topic *model.Topic) error {
	return database.GetDB().Create(topic).Error
}

func (dao *topicDao) Update(topic *model.Topic) error {
	return database.GetDB().Save(topic).Error
}

func (dao *topicDao) Updates(id int64, columns map[string]interface{}) error {
	return database.GetDB().Model(&model.Topic{}).Where("id = ?", id).Updates(columns).Error
}

func (dao *topicDao) UpdateColumn(id int64, name string, value interface{}) error {
	return database.GetDB().Model(&model.Topic{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (dao *topicDao) Delete(id int64) (err error) {
	return database.GetDB().Model(&model.Topic{}).Delete("id", id).Error
}
