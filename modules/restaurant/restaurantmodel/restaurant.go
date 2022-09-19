package restaurantmodel

import "quan/go/common"

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
	Addr            string `json:"address" gorm:"column:addr;"`
	Logo  *common.Image `json:"logo" gorm:"colum:logo;"`
	Cover  *common.Images `json:"cover" gorm:"colum:cover;"`

}

func (Restaurant) TableName() string {
	return "restaurants"
}


func (data*Restaurant)Mask(isAdminOrOwner bool){
	data.GenUID(common.DbTypeRestaurant)
}