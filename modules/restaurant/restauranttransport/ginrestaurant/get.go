package ginrestaurant

import (
	"net/http"
	"quan/go/common"
	"quan/go/modules/restaurant/restaurantbiz"
	"quan/go/modules/restaurant/restaurantstorage"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func GetRestaurant(db*gorm.DB) gin.HandlerFunc{
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest,common.ErrInvalidRequest(err))
			return
		}
		store := restaurantstorage.NewSqlStore(db)
		biz :=restaurantbiz.NewGetRestaurantBiz(store)

		data, errNew := biz.GetRestaurant(ctx.Request.Context(),id)

		if errNew!=nil{
			ctx.JSON(400,errNew)
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
		
	}
}