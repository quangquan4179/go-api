package ginrestaurant

import (
	"net/http"
	"quan/go/common"
	"quan/go/component"
	"quan/go/modules/restaurant/restaurantbiz"
	"quan/go/modules/restaurant/restaurantstorage"

	"github.com/gin-gonic/gin"
)


func GetRestaurant(atx component.AppContext) gin.HandlerFunc{
	return func(ctx *gin.Context) {
		// id, err := strconv.Atoi(ctx.Param("id"))
		uid, err := common.FromBase58(ctx.Param("id"))



		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}


		store := restaurantstorage.NewSqlStore(atx.GetMainDBConnection())
		biz :=restaurantbiz.NewGetRestaurantBiz(store)

		data, errNew := biz.GetRestaurant(ctx.Request.Context(),int(uid.GetLocalID()))

		if errNew!=nil{
			panic(errNew)
		}

		data.Mask(false)

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
		
	}
}