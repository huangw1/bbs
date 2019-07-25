/**
 * @Author: huangw1
 * @Date: 2019/7/26 16:48
 */

package dao

import (
	"github.com/huangw1/bbs/database"
	"github.com/huangw1/bbs/model"
	"github.com/jinzhu/gorm"
)

type ArticleDao interface {
	Get(id int64) (*model.Article, error)
	Take(where ...interface{}) (*model.Article, error)
	QueryCondition(condition *database.Condition) ([]*model.Article, error)
	QueryParams(params *database.Params) ([]*model.Article, *database.Paging, error)
	Create(article *model.Article) error
	Update(article *model.Article) error
	Updates(id int64, columns map[string]interface{}) error
	UpdateColumn(id int64, name string, value interface{}) error
	Delete(id int64) (err error)
}

func NewArticleDao(db *gorm.DB) ArticleDao {
	return &articleDao{db}
}

type articleDao struct {
	db *gorm.DB
}

func (dao *articleDao) Get(id int64) (*model.Article, error) {
	item := new(model.Article)
	err := dao.db.First(item, "id = ?", id).Error
	return item, err
}

func (dao *articleDao) Take(where ...interface{}) (*model.Article, error) {
	item := new(model.Article)
	err := dao.db.Take(item, where...).Error
	return item, err
}

func (dao *articleDao) QueryCondition(condition *database.Condition) ([]*model.Article, error) {
	list := make([]*model.Article, 0)
	err := condition.DoQuery(dao.db).Find(&list).Error
	return list, err
}

func (dao *articleDao) QueryParams(params *database.Params) ([]*model.Article, *database.Paging, error) {
	list := make([]*model.Article, 0)
	err := params.StartQuery(dao.db).Find(&list).Count(&params.Paging.Total).Error
	return list, params.Paging, err
}

func (dao *articleDao) Create(article *model.Article) error {
	return dao.db.Create(article).Error
}

func (dao *articleDao) Update(article *model.Article) error {
	return  dao.db.Save(article).Error
}

func (dao *articleDao) Updates(id int64, columns map[string]interface{}) error {
	return dao.db.Model(&model.Article{}).Where("id = ?", id).Updates(columns).Error
}

func (dao *articleDao) UpdateColumn(id int64, name string, value interface{}) error {
	return dao.db.Model(&model.Article{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (dao *articleDao) Delete(id int64) (err error) {
	return dao.db.Model(&model.Article{}).Delete("id", id).Error
}