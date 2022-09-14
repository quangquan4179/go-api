package restaurantstorage

import "gorm.io/gorm"



type SqlStore struct {
	db *gorm.DB
}

func NewSqlStore(db*gorm.DB)*SqlStore {
	return &SqlStore{db: db}
}