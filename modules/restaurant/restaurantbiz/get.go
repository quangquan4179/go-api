package restaurantbiz

import (
	"context"
	"quan/go/common"
	"quan/go/modules/restaurant/restaurantmodel"
)

type GetRestaurantStore interface {
	FindDataByCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type getRestaurantBiz struct {
	store GetRestaurantStore
}

func NewGetRestaurantBiz(store GetRestaurantStore) *getRestaurantBiz {
	return &getRestaurantBiz{store: store}
}

func (biz *getRestaurantBiz) GetRestaurant(ctx context.Context, id int) (*restaurantmodel.Restaurant, error) {
	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err != common.ErrRecordNotFound {
			return nil, common.ErrCannotCreateEntity(restaurantmodel.EntityName,err)
		}
		return nil, common.ErrCannotCreateEntity(restaurantmodel.EntityName,err)
	}

	if data.Status == 0 {
		return nil, common.ErrEntityDeleted(restaurantmodel.EntityName,nil)
	}
	return data, nil

}
