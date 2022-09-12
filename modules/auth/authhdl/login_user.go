package authhdl

import (
	"quan/go/common"

	"github.com/gin-gonic/gin"
)

func Login(appctx common.AppContext, secretKey string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}