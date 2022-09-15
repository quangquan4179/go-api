package restaurantstorage

import (
	"context"
	"quan/go/common"
	"quan/go/modules/restaurant/restaurantmodel"

	"gorm.io/gorm"
)

func (store *SqlStore) FindDataByCondition(ctx context.Context, conditions map[string]interface{}, moreInfos ...string)(*restaurantmodel.Restaurant, error){
	db:= store.db

	for i:= range moreInfos {
		db=db.Preload(moreInfos[i])
	}
	db  = db.Table(restaurantmodel.Restaurant{}.TableName()).Where(conditions)
	var restaurant restaurantmodel.Restaurant

	if  err:= db.First(&restaurant).Error; err!=nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &restaurant, nil
}