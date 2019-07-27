/**
 * @Author: huangw1
 * @Date: 2019/7/26 16:48
 */

package dao

import (
	"github.com/huangw1/bbs/database"
	"github.com/huangw1/bbs/model"
)

type FavoriteDao interface {
	Get(id int64) (*model.Favorite, error)
	Take(where ...interface{}) (*model.Favorite, error)
	QueryCondition(condition *model.Condition) ([]*model.Favorite, error)
	QueryParams(params *model.Params) ([]*model.Favorite, *model.Paging, error)
	Create(favorite *model.Favorite) error
	Update(favorite *model.Favorite) error
	Updates(id int64, columns map[string]interface{}) error
	UpdateColumn(id int64, name string, value interface{}) error
	Delete(id int64) (err error)
}

var FavoriteDaoImpl = NewFavoriteDao()

func NewFavoriteDao() FavoriteDao {
	return &favoriteDao{}
}

type favoriteDao struct {
}

func (dao *favoriteDao) Get(id int64) (*model.Favorite, error) {
	item := new(model.Favorite)
	err := database.GetDB().First(item, "id = ?", id).Error
	return item, err
}

func (dao *favoriteDao) Take(where ...interface{}) (*model.Favorite, error) {
	item := new(model.Favorite)
	err := database.GetDB().Take(item, where...).Error
	return item, err
}

func (dao *favoriteDao) QueryCondition(condition *model.Condition) ([]*model.Favorite, error) {
	list := make([]*model.Favorite, 0)
	err := DoQuery(database.GetDB(), condition).Find(&list).Error
	return list, err
}

func (dao *favoriteDao) QueryParams(params *model.Params) ([]*model.Favorite, *model.Paging, error) {
	list := make([]*model.Favorite, 0)
	err := StartQuery(database.GetDB(), params).Find(&list).Count(&params.Paging.Total).Error
	return list, params.Paging, err
}

func (dao *favoriteDao) Create(favorite *model.Favorite) error {
	return database.GetDB().Create(favorite).Error
}

func (dao *favoriteDao) Update(favorite *model.Favorite) error {
	return database.GetDB().Save(favorite).Error
}

func (dao *favoriteDao) Updates(id int64, columns map[string]interface{}) error {
	return database.GetDB().Model(&model.Favorite{}).Where("id = ?", id).Updates(columns).Error
}

func (dao *favoriteDao) UpdateColumn(id int64, name string, value interface{}) error {
	return database.GetDB().Model(&model.Favorite{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (dao *favoriteDao) Delete(id int64) (err error) {
	return database.GetDB().Model(&model.Favorite{}).Delete("id", id).Error
}
