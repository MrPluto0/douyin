package models

import (
	"sync"
)

// User Table = Gorm Model
type User struct {
	CommonModel
	Name          string
	Password      string
	FollowCount   int
	FollowerCount int
	IsFollow      bool
}

// User Dao
type UserDao struct{}

var (
	userDao *UserDao
	once    sync.Once
)

func NewUserDaoInstance() *UserDao {
	once.Do(func() {
		userDao = &UserDao{}
	})
	return userDao
}

func (uD *UserDao) QueryUser(name string) (u User, err error) {
	err = DB.Where("name = ?", name).First(&u).Error
	return u, err
}

func (uD *UserDao) CreateUser() (rowsAffected int64, err error) {
	return 1, nil
}
