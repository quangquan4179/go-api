package ginupload

import (
	"fmt"
	"quan/go/common"
	"quan/go/component"

	"github.com/gin-gonic/gin"
)
func Upload(appCtx component.AppContext) func(*gin.Context){
	return func(ctx *gin.Context) {
		fileHeader, err :=ctx.FormFile("file")
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		ctx.SaveUploadedFile(fileHeader, fmt.Sprintf("./static/%s",fileHeader.Filename))
		ctx.JSON(200,common.SimpleSuccessResponse(true))
	}
}