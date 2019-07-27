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

type ArticleTagDao interface {
	Get(id int64) (*model.ArticleTag, error)
	Take(where ...interface{}) (*model.ArticleTag, error)
	QueryCondition(condition *model.Condition) ([]*model.ArticleTag, error)
	QueryParams(params *model.Params) ([]*model.ArticleTag, *model.Paging, error)
	Create(article *model.ArticleTag) error
	Update(article *model.ArticleTag) error
	Updates(id int64, columns map[string]interface{}) error
	UpdateColumn(id int64, name string, value interface{}) error
	Delete(id int64) (err error)

	AddArticleTags(articleId int64, tagIds []int64) error
	RemoveArticleTags(articleId int64) error
	GetByArticleId(articleId int64) ([]*model.ArticleTag, error)
}

var ArticleTagDaoImpl = NewArticleTagDao()

func NewArticleTagDao() ArticleTagDao {
	return &articleTagDao{}
}

type articleTagDao struct {
}

func (dao *articleTagDao) Get(id int64) (*model.ArticleTag, error) {
	item := new(model.ArticleTag)
	err := database.GetDB().First(item, "id = ?", id).Error
	return item, err
}

func (dao *articleTagDao) Take(where ...interface{}) (*model.ArticleTag, error) {
	item := new(model.ArticleTag)
	err := database.GetDB().Take(item, where...).Error
	return item, err
}

func (dao *articleTagDao) QueryCondition(condition *model.Condition) ([]*model.ArticleTag, error) {
	list := make([]*model.ArticleTag, 0)
	err := DoQuery(database.GetDB(), condition).Find(&list).Error
	return list, err
}

func (dao *articleTagDao) QueryParams(params *model.Params) ([]*model.ArticleTag, *model.Paging, error) {
	list := make([]*model.ArticleTag, 0)
	err := StartQuery(database.GetDB(), params).Find(&list).Count(&params.Paging.Total).Error
	return list, params.Paging, err
}

func (dao *articleTagDao) Create(article *model.ArticleTag) error {
	return database.GetDB().Create(article).Error
}

func (dao *articleTagDao) Update(article *model.ArticleTag) error {
	return database.GetDB().Save(article).Error
}

func (dao *articleTagDao) Updates(id int64, columns map[string]interface{}) error {
	return database.GetDB().Model(&model.ArticleTag{}).Where("id = ?", id).Updates(columns).Error
}

func (dao *articleTagDao) UpdateColumn(id int64, name string, value interface{}) error {
	return database.GetDB().Model(&model.ArticleTag{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (dao *articleTagDao) Delete(id int64) (err error) {
	return database.GetDB().Model(&model.ArticleTag{}).Delete("id", id).Error
}

func (dao *articleTagDao) AddArticleTags(articleId int64, tagIds []int64) error {
	if len(tagIds) == 0 {
		return nil
	}
	err := database.Tx(database.GetDB(), func(db *gorm.DB) error {
		for _, tagId := range tagIds {
			articleTag := &model.ArticleTag{}
			articleTag.ArticleId = articleId
			articleTag.TagId = tagId
			articleTag.CreateTime = extension.NowTimestamp()
			articleTag.UpdateTime = extension.NowTimestamp()
			if err := dao.Create(articleTag); err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func (dao *articleTagDao) RemoveArticleTags(articleId int64) error {
	return database.GetDB().Where("articleId = ?", articleId).Delete(&model.ArticleTag{}).Error
}

func (dao *articleTagDao) GetByArticleId(articleId int64) ([]*model.ArticleTag, error) {
	return dao.QueryCondition(model.NewCondition("articleId = ?", articleId))
}
