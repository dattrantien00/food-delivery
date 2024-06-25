package ginuser

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		
		user := c.MustGet(common.CurrentUser)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(user))
	}
}

// func Createuser(c *gin.Context) {

// }
