package restaurantmodel

import "quan/go/common"

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"address" gorm:"column:addr;"`
	Logo  *common.Image `json:"logo" gorm:"colum:logo;"`
	Cover  *common.Images `json:"cover" gorm:"colum:cover;"`


}
func (RestaurantCreate) TableName() string {
	return  Restaurant{}.TableName()
}