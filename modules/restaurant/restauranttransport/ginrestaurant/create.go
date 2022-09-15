package ginrestaurant

import (
	"net/http"
	"quan/go/common"
	"quan/go/modules/restaurant/restaurantbiz"
	"quan/go/modules/restaurant/restaurantmodel"
	"quan/go/modules/restaurant/restaurantstorage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func CreateRestaurant(db  *gorm.DB) gin.HandlerFunc{
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate
		if err:= c.ShouldBind(&data); err!=nil{
			c.JSON(http.StatusBadRequest,common.ErrInvalidRequest(err))
			return

		}

		store := restaurantstorage.NewSqlStore(db);
		biz := restaurantbiz.NewCreateRestaurantBiz(store)
		if err:=biz.CreateRestaurant(c.Request.Context(),&data); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"error":err.Error(),
			})
			return

		}
		c.JSON(http.StatusOK,data)
	}
}