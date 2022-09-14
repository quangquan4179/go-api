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


func ListRestaurant(db* gorm.DB)gin.HandlerFunc{
	return func(c *gin.Context) {

		var filter restaurantmodel.Filter
		if err:= c.ShouldBind(&filter); err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"error":err.Error(),
			})
			return

		}
		var paging common.Paging
		if err:= c.ShouldBind(&paging); err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"error":err.Error(),
			})
			return

		}
		paging.Fulfill()

		store := restaurantstorage.NewSqlStore(db);
		biz := restaurantbiz.NewListRestaurantBiz(store)
		result, err:=biz.ListRestaurant(c.Request.Context(),nil,&filter,&paging,"")
		if  err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"error":err.Error(),
			})
			return

		}
		c.JSON(http.StatusOK,common.NewSuccessResponse(result,paging, filter))
	}
}