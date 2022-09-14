package restaurantbiz
import (
	"context"
	"quan/go/modules/restaurant/restaurantmodel"
)

type CreateRestaurantStore interface{
	Create(context context.Context, data *restaurantmodel.RestaurantCreate) error 
	Find(ctx context.Context, conditions map[string]interface{}, moreInfos ...string)(*restaurantmodel.Restaurant, error)
}
type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz( store CreateRestaurantStore) *createRestaurantBiz{
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	err := biz.store.Create(ctx, data)
	return err

}