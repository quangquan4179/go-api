package authhdl

import (
	"net/http"
	"quan/go/common"
	"quan/go/component"

	"quan/go/modules/auth/authmodel"
	"quan/go/modules/auth/authrepo"
	"quan/go/modules/user/userstorage"

	"github.com/gin-gonic/gin"
)


func Register(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		db := appCtx.GetMainDBConnection()
		var user authmodel.CreateUser

		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		if err := user.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := userstorage.NewUserMysql(db)
		repo := authrepo.NewAuthRepo(store)

		userId, err := repo.Register(c.Request.Context(), &user)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(userId))
	}
}