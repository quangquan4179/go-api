package restaurantmodel

type RestaurantUpate struct {
	
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"address" gorm:"column:addr;"`
}
func (RestaurantUpate) TableName() string {
	return Restaurant{}.TableName()
}