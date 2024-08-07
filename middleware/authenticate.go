package middleware

import (
	"context"
	"errors"
	"food-delivery/common"
	"food-delivery/component/appctx"
	"food-delivery/component/tokenprovider/jwt"
	usermodel "food-delivery/module/user/model"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthenStore interface {
	FindUserByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		relations ...string,
	) (*usermodel.User, error)
}

func extractTokenFromHeader(s string) (string, error) {
	part := strings.Split(s, " ")
	if part[0] != "Bearer" || len(part) != 2 || strings.TrimSpace(part[1]) == "" {
		return "", ErrWrongAuthHeader
	}
	return part[1], nil
}

func RequireAuth(appCtx appctx.AppContext, authenStore AuthenStore) func(c *gin.Context) {
	tokenProvider := jwt.NewJwtProvider(appCtx.SecretKey())

	return func(c *gin.Context) {
		token, err := extractTokenFromHeader(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}
		// db := appCtx.GetMainDBConnection()
		// store := userstorage.NewSQLStore(db)

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}

		user, err := authenStore.FindUserByCondition(c.Request.Context(), map[string]interface{}{
			"id": payload.UserId,
		})
		if err != nil {
			panic(err)
		}

		if user.Status == 0 {
			panic(common.ErrNoPermission(err))
		}

		user.Mask()

		c.Set(common.CurrentUser, user)
		c.Next()
	}
}

var (
	ErrWrongAuthHeader = common.NewCustomError(
		errors.New("Wrong auth header"),
		"Wrong auth header",
		"ErrWrongAuthHeader",
	)
)
