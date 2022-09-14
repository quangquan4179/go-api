package authmodel

import "quan/go/modules/user/usermodel"

type LoginUser struct {
	Email    string `json:"email" form:"email" gorm:"column:email;"`
	Password string `json:"password" form:"password" gorm:"column:password;"`
}

func (LoginUser) TableName() string {
	return usermodel.User{}.TableName()
}
