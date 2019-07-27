/**
 * @Author: huangw1
 * @Date: 2019/7/26 16:48
 */

package dao

import (
	"github.com/huangw1/bbs/database"
	"github.com/huangw1/bbs/model"
	"github.com/huangw1/bbs/pkg/extension"
	"github.com/jinzhu/gorm"
)

type TopicTagDao interface {
	Get(id int64) (*model.TopicTag, error)
	Take(where ...interface{}) (*model.TopicTag, error)
	QueryCondition(condition *model.Condition) ([]*model.TopicTag, error)
	QueryParams(params *model.Params) ([]*model.TopicTag, *model.Paging, error)
	Create(article *model.TopicTag) error
	Update(article *model.TopicTag) error
	Updates(id int64, columns map[string]interface{}) error
	UpdateColumn(id int64, name string, value interface{}) error
	Delete(id int64) (err error)

	AddTopicTags(articleId int64, tagIds []int64) error
	RemoveTopicTags(articleId int64) error
}

var TopicTagDaoImpl = NewTopicTagDao()

func NewTopicTagDao() TopicTagDao {
	return &topicTagDao{}
}

type topicTagDao struct {
}

func (dao *topicTagDao) Get(id int64) (*model.TopicTag, error) {
	item := new(model.TopicTag)
	err := database.GetDB().First(item, "id = ?", id).Error
	return item, err
}

func (dao *topicTagDao) Take(where ...interface{}) (*model.TopicTag, error) {
	item := new(model.TopicTag)
	err := database.GetDB().Take(item, where...).Error
	return item, err
}

func (dao *topicTagDao) QueryCondition(condition *model.Condition) ([]*model.TopicTag, error) {
	list := make([]*model.TopicTag, 0)
	err := DoQuery(database.GetDB(), condition).Find(&list).Error
	return list, err
}

func (dao *topicTagDao) QueryParams(params *model.Params) ([]*model.TopicTag, *model.Paging, error) {
	list := make([]*model.TopicTag, 0)
	err := StartQuery(database.GetDB(), params).Find(&list).Count(&params.Paging.Total).Error
	return list, params.Paging, err
}

func (dao *topicTagDao) Create(article *model.TopicTag) error {
	return database.GetDB().Create(article).Error
}

func (dao *topicTagDao) Update(article *model.TopicTag) error {
	return database.GetDB().Save(article).Error
}

func (dao *topicTagDao) Updates(id int64, columns map[string]interface{}) error {
	return database.GetDB().Model(&model.TopicTag{}).Where("id = ?", id).Updates(columns).Error
}

func (dao *topicTagDao) UpdateColumn(id int64, name string, value interface{}) error {
	return database.GetDB().Model(&model.TopicTag{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (dao *topicTagDao) Delete(id int64) (err error) {
	return database.GetDB().Model(&model.TopicTag{}).Delete("id", id).Error
}

func (dao *topicTagDao) AddTopicTags(topicId int64, tagIds []int64) error {
	if len(tagIds) == 0 {
		return nil
	}
	err := database.Tx(database.GetDB(), func(db *gorm.DB) error {
		for _, tagId := range tagIds {
			topicTag := &model.TopicTag{}
			topicTag.TopicId = topicId
			topicTag.TagId = tagId
			topicTag.CreateTime = extension.NowTimestamp()
			topicTag.UpdateTime = extension.NowTimestamp()
			if err := dao.Create(topicTag); err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func (dao *topicTagDao) RemoveTopicTags(topicId int64) error {
	return database.GetDB().Where("topicId = ?", topicId).Delete(&model.TopicTag{}).Error
}
