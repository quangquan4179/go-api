package ginrestaurant

import (
	"net/http"
	"quan/go/common"
	"quan/go/component"
	"quan/go/modules/restaurant/restaurantbiz"
	"quan/go/modules/restaurant/restaurantmodel"
	"quan/go/modules/restaurant/restaurantstorage"

	"github.com/gin-gonic/gin"
)


func ListRestaurant(atx component.AppContext)gin.HandlerFunc{
	return func(c *gin.Context) {

		var filter restaurantmodel.Filter
		if err:= c.ShouldBind(&filter); err!=nil{
			panic(common.ErrInvalidRequest(err))

		}
		var paging common.Paging
		if err:= c.ShouldBind(&paging); err!=nil{
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		store := restaurantstorage.NewSqlStore(atx.GetMainDBConnection());
		biz := restaurantbiz.NewListRestaurantBiz(store)
		result, err:=biz.ListRestaurant(c.Request.Context(),nil,&filter,&paging,"")

		for i:= range result {
			result[i].Mask(false)
		}
		if  err!=nil{
			panic(err)

		}
		c.JSON(http.StatusOK,common.NewSuccessResponse(result,paging, filter))
	}
}