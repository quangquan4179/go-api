package userstorage

import (
	"context"
	"quan/go/common"
	"quan/go/modules/user/usermodel"

	"gorm.io/gorm"
)

func (store *userMysql) FindUserByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	relations ...string,
) (*usermodel.User, error) {

	var user usermodel.User
	db := store.db.Table(usermodel.User{}.TableName())

	if err := db.Where(conditions).First(&user).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, common.ErrDB(err)
		}

		return nil, err
	}

	return &user, nil
}
