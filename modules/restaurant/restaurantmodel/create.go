package restaurantmodel

type RestaurantCreate struct {
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"address" gorm:"column:addr;"`
}
func (RestaurantCreate) TableName() string {
	return  Restaurant{}.TableName()
}