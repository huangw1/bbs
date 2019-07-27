/**
 * @Author: huangw1
 * @Date: 2019/7/30 11:45
 */

package service

import (
	"fmt"
	"github.com/huangw1/bbs/dao"
	"github.com/huangw1/bbs/database"
	"github.com/huangw1/bbs/model"
)

var UserService = NewUserService()

func NewUserService() *userService {
	return &userService{}
}

type userService struct{}

func (service *userService) Get(id int64) (*model.User, error) {
	return dao.UserDaoImpl.Get(id)
}

func (service *userService) Take(where ...interface{}) (*model.User, error) {
	return dao.UserDaoImpl.Take(where...)
}

func (service *userService) QueryCondition(condition *model.Condition) ([]*model.User, error) {
	return dao.UserDaoImpl.QueryCondition(condition)
}

func (service *userService) QueryParams(params *model.Params) ([]*model.User, *model.Paging, error) {
	return dao.UserDaoImpl.QueryParams(params)
}

func (service *userService) Create(user *model.User) error {
	return dao.UserDaoImpl.Create(user)
}

func (service *userService) Update(user *model.User) error {
	return dao.UserDaoImpl.Update(user)
}

func (service *userService) Updates(id int64, columns map[string]interface{}) error {
	return dao.UserDaoImpl.Updates(id, columns)
}

func (service *userService) UpdateColumn(id int64, name string, value interface{}) error {
	return dao.UserDaoImpl.UpdateColumn(id, name, value)
}

func (service *userService) Delete(id int64) (err error) {
	return dao.UserDaoImpl.Delete(id)
}

func (service *userService) GetByEmail(email string) (*model.User, error) {
	return dao.UserDaoImpl.GetByEmail(email)
}

func (service *userService) GetByUsername(username string) (*model.User, error) {
	return dao.UserDaoImpl.GetByUsername(username)
}

func (service *userService) GetActiveUserIds() ([]int64, error) {
	// todo 根据积分、文章、话题、登录？暂取文章。
	var userIds []int64
	rows, err := database.GetDB().Raw("select userId, count(*) c from t_article group by userId order by c desc limit 20").Rows()
	if err != nil {
		return userIds, err
	}
	for rows.Next() {
		var c int
		var userId int64
		err := rows.Scan(&userId, &c)
		if err != nil {
			continue
		}
		userIds = append(userIds, userId)
	}
	fmt.Printf("userIds: %+v", userIds)
	return userIds, nil
}
