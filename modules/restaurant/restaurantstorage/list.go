package restaurantstorage

import (
	"context"
	"quan/go/common"
	"quan/go/modules/restaurant/restaurantmodel"
)


func (store *SqlStore) ListDataByCondition(ctx context.Context, 
	conditions map[string]interface{}, 
	filter *restaurantmodel.Filter, 
	paging *common.Paging,
	moreInfos ...string)([]restaurantmodel.Restaurant, error){
		var result []restaurantmodel.Restaurant
		db:= store.db
		for i:= range moreInfos {
			db = db.Preload(moreInfos[i])
		}
		db  = db.Table(restaurantmodel.Restaurant{}.TableName()).Where(conditions)

		if v:= filter; v !=nil {
			if v.CityId>0 {
				db = db.Where("city_id = ?", v.CityId)
			}
		}
		if err:=db.Count(&paging.Total).Error; err !=nil{
			return nil, common.ErrDB(err)
		}

		if err := db.Offset((paging.Page-1)*paging.Limit).Limit(paging.Limit).Find(&result).Error; err!=nil{
			return nil, common.ErrDB(err)

		}
	return result,nil
}