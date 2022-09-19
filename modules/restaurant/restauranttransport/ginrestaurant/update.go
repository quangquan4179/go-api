package ginrestaurant

import (
	"net/http"
	"quan/go/common"
	"quan/go/component"
	"quan/go/modules/restaurant/restaurantbiz"
	"quan/go/modules/restaurant/restaurantmodel"
	"quan/go/modules/restaurant/restaurantstorage"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateRestaurant(atz component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data restaurantmodel.RestaurantUpdate
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		if err := ctx.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))

		}

		store := restaurantstorage.NewSqlStore(atz.GetMainDBConnection())
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurant(ctx.Request.Context(), id, &data); err != nil {
			panic(err)

		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}
