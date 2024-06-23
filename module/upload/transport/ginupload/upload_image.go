package ginupload

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	uploadbusiness "food-delivery/module/upload/biz"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadImage(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// db := appCtx.GetMainDBConnection()
		file, fileHeader, err := c.Request.FormFile("file")
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "img")

		defer file.Close()
		dataBytes := make([]byte, fileHeader.Size)
		_, err = file.Read(dataBytes)
		if err != nil {
			panic(common.ErrInvalidRequest(err))

		}

		biz := uploadbusiness.NewUploadBiz(appCtx.UploadProvider(), nil)
		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
