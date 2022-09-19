package ginrestaurant

import (
	"net/http"
	"quan/go/common"
	"quan/go/component"
	"quan/go/modules/restaurant/restaurantbiz"
	"quan/go/modules/restaurant/restaurantstorage"

	"github.com/gin-gonic/gin"
)

func DeleteRestaurant(atx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))
		
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSqlStore(atx.GetMainDBConnection())
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(ctx.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}
