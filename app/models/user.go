package models

import "sync"

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey;column:id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

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

func (*UserDao) QueryUser(username string) *User {
	return &User{}
}

func (*UserDao) CreateUser() (rowsAffected int64, err error) {
	return 1, nil
}
