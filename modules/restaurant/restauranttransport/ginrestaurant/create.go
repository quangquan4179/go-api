package ginrestaurant

import (
	"net/http"
	"quan/go/common"
	"quan/go/modules/restaurant/restaurantbiz"
	"quan/go/modules/restaurant/restaurantmodel"
	"quan/go/modules/restaurant/restaurantstorage"
	"quan/go/component"


	"github.com/gin-gonic/gin"
)


func CreateRestaurant(atx component.AppContext) gin.HandlerFunc{
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate
		if err:= c.ShouldBind(&data); err!=nil{
			panic(common.ErrInvalidRequest(err))

		}

		store := restaurantstorage.NewSqlStore(atx.GetMainDBConnection());
		biz := restaurantbiz.NewCreateRestaurantBiz(store)
		if err:=biz.CreateRestaurant(c.Request.Context(),&data); err != nil{
			panic(err)

		}
		data.GenUID(common.DbTypeRestaurant)
		c.JSON(http.StatusOK,common.SimpleSuccessResponse(data.FakeId))
	}
}