package restaurantmodel

type RestaurantUpdate struct {
	
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"address" gorm:"column:addr;"`
}
func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}