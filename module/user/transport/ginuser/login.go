package ginuser

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	hasher "food-delivery/component/hash"
	"food-delivery/component/tokenprovider/jwt"
	userbiz "food-delivery/module/user/biz"
	usermodel "food-delivery/module/user/model"
	userstorage "food-delivery/module/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserLogin

		if err := c.ShouldBindJSON(&data); err != nil {

			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		hasher := hasher.NewMd5Hash()
		jwtProvider := jwt.NewJwtProvider(appCtx.SecretKey())
		biz := userbiz.NewLoginBiz(store, jwtProvider, hasher, 60*60*24*30)

		token, err := biz.Login(c.Request.Context(), &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(token))
	}
}

// func Createuser(c *gin.Context) {

// }
