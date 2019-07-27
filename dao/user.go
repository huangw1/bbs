/**
 * @Author: huangw1
 * @Date: 2019/7/26 16:48
 */

package dao

import (
	"github.com/huangw1/bbs/database"
	"github.com/huangw1/bbs/model"
)

type UserDao interface {
	Get(id int64) (*model.User, error)
	Take(where ...interface{}) (*model.User, error)
	QueryCondition(condition *model.Condition) ([]*model.User, error)
	QueryParams(params *model.Params) ([]*model.User, *model.Paging, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Updates(id int64, columns map[string]interface{}) error
	UpdateColumn(id int64, name string, value interface{}) error
	Delete(id int64) (err error)

	GetByEmail(email string) (*model.User, error)
	GetByUsername(username string) (*model.User, error)
}

var UserDaoImpl = NewUserDao()

func NewUserDao() UserDao {
	return &userDao{}
}

type userDao struct {
}

func (dao *userDao) Get(id int64) (*model.User, error) {
	item := new(model.User)
	err := database.GetDB().First(item, "id = ?", id).Error
	return item, err
}

func (dao *userDao) Take(where ...interface{}) (*model.User, error) {
	item := new(model.User)
	err := database.GetDB().Take(item, where...).Error
	return item, err
}

func (dao *userDao) QueryCondition(condition *model.Condition) ([]*model.User, error) {
	list := make([]*model.User, 0)
	err := DoQuery(database.GetDB(), condition).Find(&list).Error
	return list, err
}

func (dao *userDao) QueryParams(params *model.Params) ([]*model.User, *model.Paging, error) {
	list := make([]*model.User, 0)
	err := StartQuery(database.GetDB(), params).Find(&list).Count(&params.Paging.Total).Error
	return list, params.Paging, err
}

func (dao *userDao) Create(user *model.User) error {
	return database.GetDB().Create(user).Error
}

func (dao *userDao) Update(user *model.User) error {
	return database.GetDB().Save(user).Error
}

func (dao *userDao) Updates(id int64, columns map[string]interface{}) error {
	return database.GetDB().Model(&model.User{}).Where("id = ?", id).Updates(columns).Error
}

func (dao *userDao) UpdateColumn(id int64, name string, value interface{}) error {
	return database.GetDB().Model(&model.User{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (dao *userDao) Delete(id int64) (err error) {
	return database.GetDB().Model(&model.User{}).Delete("id", id).Error
}

func (dao *userDao) GetByEmail(email string) (*model.User, error) {
	return dao.Take("email = ?", email)
}

func (dao *userDao) GetByUsername(username string) (*model.User, error) {
	return dao.Take("username = ?", username)
}
