package restaurantstorage

import (
	"context"
	"quan/go/common"
	"quan/go/modules/restaurant/restaurantmodel"
)

func (store *SqlStore) Create(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	print("create")

	db := store.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
