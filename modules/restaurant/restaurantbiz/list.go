package restaurantbiz

import (
	"context"
	"quan/go/common"
	"quan/go/modules/restaurant/restaurantmodel"
)

type ListRestaurantStore interface {
	ListDataByCondition(ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreInfos ...string) ([]restaurantmodel.Restaurant, error)
}
type listRestaurantBiz struct {
	store ListRestaurantStore
}

func NewListRestaurantBiz(store ListRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}

func (biz *listRestaurantBiz) ListRestaurant(ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreInfos ...string,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListDataByCondition(ctx, conditions, filter, paging)
	return result, err

}
