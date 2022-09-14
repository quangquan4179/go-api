package restaurantstorage

import (
	"context"
	"quan/go/common"
	"quan/go/modules/restaurant/restaurantmodel"
)

func (store *SqlStore) Find(ctx context.Context, conditions map[string]interface{}, moreInfos ...string)(*restaurantmodel.Restaurant, error){
	db:= store.db

	for i:= range moreInfos {
		db=db.Preload(moreInfos[i])
	}
	var restaurant restaurantmodel.Restaurant

	if  err:= db.First(&restaurant).Error; err!=nil {
		return nil, common.ErrDB(err)
	}

	return &restaurant, nil
}