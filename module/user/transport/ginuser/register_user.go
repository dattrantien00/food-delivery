package ginuser

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	hasher "food-delivery/component/hash"
	userbiz "food-delivery/module/user/biz"
	usermodel "food-delivery/module/user/model"
	userstorage "food-delivery/module/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserCreate

		if err := c.ShouldBindJSON(&data); err != nil {

			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		hasher := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBiz(store, hasher)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId))
	}
}

// func Createuser(c *gin.Context) {

// }
