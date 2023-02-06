package models

import (
	"sync"
)

// User Table = Gorm Model
type User struct {
	CommonModel
	Name          string `gorm:"type:varchar(32);not null;unique;index:name_idx"`
	Password      string `gorm:"type:varchar(32);not null"`
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
	err = DB.Where("name = ?", name).Limit(1).Find(&u).Error
	return u, err
}

func (uD *UserDao) CreateUser(name string, pwd string) (rowsAffected int64, err error) {
	user := User{Name: name, Password: pwd}
	result := DB.Create(&user)
	return result.RowsAffected, result.Error
}
