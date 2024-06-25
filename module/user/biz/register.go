package userbusiness

import (
	"context"
	"food-delivery/common"
	usermodel "food-delivery/module/user/model"
)

type registerUser interface {
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
	FindUserByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		relations ...string,
	) (*usermodel.User, error)
}

type hasher interface {
	Hash(data string) string
}

type registerBiz struct {
	store  registerUser
	hasher hasher
}

func NewRegisterBiz(store registerUser, hash hasher) *registerBiz {
	return &registerBiz{
		store:  store,
		hasher: hash,
	}
}

func (biz *registerBiz) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, _ := biz.store.FindUserByCondition(ctx, map[string]interface{}{
		"email": data.Email,
	})

	if user != nil {
		return usermodel.ErrEmailExist
	}

	salt := common.GenSalt(50)
	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user"

	if err := biz.store.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}
	return nil
}
