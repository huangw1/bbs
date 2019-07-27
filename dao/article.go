/**
 * @Author: huangw1
 * @Date: 2019/7/26 16:48
 */

package dao

import (
	"github.com/huangw1/bbs/database"
	"github.com/huangw1/bbs/model"
)

type ArticleDao interface {
	Get(id int64) (*model.Article, error)
	Take(where ...interface{}) (*model.Article, error)
	QueryCondition(condition *model.Condition) ([]*model.Article, error)
	QueryParams(params *model.Params) ([]*model.Article, *model.Paging, error)
	Create(article *model.Article) error
	Update(article *model.Article) error
	Updates(id int64, columns map[string]interface{}) error
	UpdateColumn(id int64, name string, value interface{}) error
	Delete(id int64) (err error)
}

var ArticleDaoImpl = NewArticleDao()

func NewArticleDao() ArticleDao {
	return &articleDao{}
}

type articleDao struct {
}

func (dao *articleDao) Get(id int64) (*model.Article, error) {
	item := new(model.Article)
	err := database.GetDB().First(item, "id = ?", id).Error
	return item, err
}

func (dao *articleDao) Take(where ...interface{}) (*model.Article, error) {
	item := new(model.Article)
	err := database.GetDB().Take(item, where...).Error
	return item, err
}

func (dao *articleDao) QueryCondition(condition *model.Condition) ([]*model.Article, error) {
	list := make([]*model.Article, 0)
	err := DoQuery(database.GetDB(), condition).Find(&list).Error
	return list, err
}

func (dao *articleDao) QueryParams(params *model.Params) ([]*model.Article, *model.Paging, error) {
	list := make([]*model.Article, 0)
	err := StartQuery(database.GetDB(), params).Find(&list).Count(&params.Paging.Total).Error
	return list, params.Paging, err
}

func (dao *articleDao) Create(article *model.Article) error {
	return database.GetDB().Create(article).Error
}

func (dao *articleDao) Update(article *model.Article) error {
	return database.GetDB().Save(article).Error
}

func (dao *articleDao) Updates(id int64, columns map[string]interface{}) error {
	return database.GetDB().Model(&model.Article{}).Where("id = ?", id).Updates(columns).Error
}

func (dao *articleDao) UpdateColumn(id int64, name string, value interface{}) error {
	return database.GetDB().Model(&model.Article{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (dao *articleDao) Delete(id int64) (err error) {
	return database.GetDB().Model(&model.Article{}).Delete("id", id).Error
}
