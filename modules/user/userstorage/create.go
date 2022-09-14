package userstorage

import (
	"context"
	"fmt"
	"log"
	"quan/go/common"
	"quan/go/modules/user/usermodel"
)

func (store *userMysql) Create(ctx context.Context, createUserData *usermodel.CreateUser) (int, error) {
	db := store.db.Begin()
	log.Println("create user data", createUserData)
	if err := db.Table(usermodel.CreateUser{}.TableName()).Create(&createUserData).Error; err != nil {
		db.Rollback()
		return 0, common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return 0, common.ErrDB(err)
	}

	fmt.Println("after create user data", createUserData)

	return createUserData.ID, nil

}
