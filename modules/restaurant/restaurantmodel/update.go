package restaurantmodel

import "quan/go/common"

type RestaurantUpdate struct {
	
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"address" gorm:"column:addr;"`
	Logo  *common.Image `json:"logo" gorm:"colum:logo;"`
	Cover  *common.Images `json:"cover" gorm:"colum:cover;"`


}
func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}