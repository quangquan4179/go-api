package restaurantstorage

import (
	"context"
	"quan/go/common"
	"quan/go/modules/restaurant/restaurantmodel"
)

func (store *SqlStore) UpdateData(ctx context.Context,
	id int,
	data *restaurantmodel.RestaurantUpdate,
) error {
	db := store.db

	if err:= db.Where("id = ?", id).Updates(data).Error; err !=nil {
		return common.ErrDB(err)
	}
	return nil
}
