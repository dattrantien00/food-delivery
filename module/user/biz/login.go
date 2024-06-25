package userbusiness

import (
	"context"
	"food-delivery/common"
	"food-delivery/component/tokenprovider"
	usermodel "food-delivery/module/user/model"
)

type loginUser interface {
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
	FindUserByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		relations ...string,
	) (*usermodel.User, error)
}

type loginBiz struct {
	store         loginUser
	hasher        hasher
	expiry        int
	tokenProvider tokenprovider.Provider
}

func NewLoginBiz(storeUser loginUser, tokenProvider tokenprovider.Provider, hasher hasher, expiry int) *loginBiz {
	return &loginBiz{
		store:         storeUser,
		hasher:        hasher,
		expiry:        expiry,
		tokenProvider: tokenProvider,
	}
}

func (biz *loginBiz) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	user, err := biz.store.FindUserByCondition(ctx, map[string]interface{}{
		"email": data.Email,
	})

	if err != nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	passHash := biz.hasher.Hash(data.Password + user.Salt)
	if passHash != user.Password {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	token, err := biz.tokenProvider.Generate(tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}, biz.expiry)

	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return token, nil
}
