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

type TagDao interface {
	Get(id int64) (*model.Tag, error)
	Take(where ...interface{}) (*model.Tag, error)
	QueryCondition(condition *model.Condition) ([]*model.Tag, error)
	QueryParams(params *model.Params) ([]*model.Tag, *model.Paging, error)
	Create(tag *model.Tag) error
	Update(tag *model.Tag) error
	Updates(id int64, columns map[string]interface{}) error
	UpdateColumn(id int64, name string, value interface{}) error
	Delete(id int64) (err error)

	GetTagInIds(tagIds []int64) ([]*model.Tag, error)
	FindByName(name string) (*model.Tag, error)
	GetOrCreate(name string) (*model.Tag, error)
	GetOrCreates(tags []string) ([]int64, error)
}

var TagDaoImpl = NewTagDao()

func NewTagDao() TagDao {
	return &tagDao{}
}

type tagDao struct {
}

func (dao *tagDao) Get(id int64) (*model.Tag, error) {
	item := new(model.Tag)
	err := database.GetDB().First(item, "id = ?", id).Error
	return item, err
}

func (dao *tagDao) Take(where ...interface{}) (*model.Tag, error) {
	item := new(model.Tag)
	err := database.GetDB().Take(item, where...).Error
	return item, err
}

func (dao *tagDao) QueryCondition(condition *model.Condition) ([]*model.Tag, error) {
	list := make([]*model.Tag, 0)
	err := DoQuery(database.GetDB(), condition).Find(&list).Error
	return list, err
}

func (dao *tagDao) QueryParams(params *model.Params) ([]*model.Tag, *model.Paging, error) {
	list := make([]*model.Tag, 0)
	err := StartQuery(database.GetDB(), params).Find(&list).Count(&params.Paging.Total).Error
	return list, params.Paging, err
}

func (dao *tagDao) Create(tag *model.Tag) error {
	return database.GetDB().Create(tag).Error
}

func (dao *tagDao) Update(tag *model.Tag) error {
	return database.GetDB().Save(tag).Error
}

func (dao *tagDao) Updates(id int64, columns map[string]interface{}) error {
	return database.GetDB().Model(&model.Tag{}).Where("id = ?", id).Updates(columns).Error
}

func (dao *tagDao) UpdateColumn(id int64, name string, value interface{}) error {
	return database.GetDB().Model(&model.Tag{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (dao *tagDao) Delete(id int64) (err error) {
	return database.GetDB().Model(&model.Tag{}).Delete("id", id).Error
}

func (dao *tagDao) GetTagInIds(tagIds []int64) ([]*model.Tag, error) {
	if len(tagIds) == 0 {
		return make([]*model.Tag, 0), nil
	}
	return dao.QueryCondition(model.NewCondition("id in (?)", tagIds))
}

func (dao *tagDao) FindByName(name string) (*model.Tag, error) {
	return dao.Take("name = ?", name)
}

func (dao *tagDao) GetOrCreate(name string) (*model.Tag, error) {
	tag, err := dao.FindByName(name)
	if err == nil {
		return tag, nil
	}
	tag = &model.Tag{}
	tag.Name = name
	tag.Status = model.TagStatusOk
	tag.CreateTime = extension.NowTimestamp()
	tag.UpdateTime = extension.NowTimestamp()
	err = dao.Create(tag)
	return tag, err
}

func (dao *tagDao) GetOrCreates(names []string) ([]int64, error) {
	tagIds := make([]int64, 0)
	err := database.Tx(database.GetDB(), func(db *gorm.DB) error {
		for _, name := range names {
			tag, err := dao.GetOrCreate(name)
			if err != nil {
				return err
			}
			tagIds = append(tagIds, tag.Id)
		}
		return nil
	})
	return tagIds, err
}
