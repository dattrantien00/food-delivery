package middleware

import (
	"errors"
	"food-delivery/common"
	"food-delivery/component/appctx"

	"github.com/gin-gonic/gin"
)

func RoleRequire(appCtx appctx.AppContext, allowRoles ...string) func(c *gin.Context) {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)

		for _, item := range allowRoles {
			if user.GetRole() == item {
				c.Next()
				return
			}
		}
		panic(common.ErrNoPermission(errors.New("invalid role user")))
	}
}
