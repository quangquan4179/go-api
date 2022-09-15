package restaurantstorage

import (
	"context"
	"quan/go/common"
	"quan/go/modules/restaurant/restaurantmodel"
)

func (store *SqlStore) DeleteByConfition(ctx context.Context,
	id int,
) error {
	db := store.db
	db  = db.Table(restaurantmodel.Restaurant{}.TableName())


	if err:= db.Where("id = ?", id).Updates(map[string]interface{}{
		"status":0,}).Error; err !=nil {
		return common.ErrDB(err)
	}
	return nil
}