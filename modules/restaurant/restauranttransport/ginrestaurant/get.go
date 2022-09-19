package ginrestaurant

import (
	"net/http"
	"quan/go/common"
	"quan/go/component"
	"quan/go/modules/restaurant/restaurantbiz"
	"quan/go/modules/restaurant/restaurantstorage"
	"strconv"

	"github.com/gin-gonic/gin"
)


func GetRestaurant(atx component.AppContext) gin.HandlerFunc{
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}


		store := restaurantstorage.NewSqlStore(atx.GetMainDBConnection())
		biz :=restaurantbiz.NewGetRestaurantBiz(store)

		data, errNew := biz.GetRestaurant(ctx.Request.Context(),id)

		if errNew!=nil{
			panic(errNew)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
		
	}
}