package ginrestaurant

import (
	"net/http"
	"quan/go/common"
	"quan/go/modules/restaurant/restaurantbiz"
	"quan/go/modules/restaurant/restaurantmodel"
	"quan/go/modules/restaurant/restaurantstorage"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data restaurantmodel.RestaurantUpdate
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}
		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(400, map[string]interface{}{
				"error": err.Error(),
			})
			return

		}

		store := restaurantstorage.NewSqlStore(db)
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurant(ctx.Request.Context(), id, &data); err != nil {

			ctx.JSON(400, err)
			return

		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}
