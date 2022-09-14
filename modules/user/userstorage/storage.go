package userstorage

import "gorm.io/gorm"

type userMysql struct {
	db *gorm.DB
}

func NewUserMysql(db*gorm.DB) *userMysql{
	return &userMysql{db: db}
}