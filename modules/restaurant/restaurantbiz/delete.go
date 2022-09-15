package restaurantbiz

import (
	"context"
	"quan/go/common"
	"quan/go/modules/restaurant/restaurantmodel"
)

type DeleteRestaurantStore interface {
	FindDataByCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
	DeleteByConfition(ctx context.Context,
		id int,
	) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id int) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id":id})

	if err !=nil {
		if err!=common.ErrRecordNotFound {
			return common.ErrCannotDeleteEntity(restaurantmodel.EntityName,err)
		}
		return common.ErrEntityNotExist(restaurantmodel.EntityName,err)
	}
	if oldData.Status==0 {
		return common.ErrEntityDeleted(restaurantmodel.EntityName,nil)
	}
	if err := biz.store.DeleteByConfition(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.EntityName,err)
	}
	return nil
}
