package memcache

import (
	"context"
	"fmt"
	usermodel "food-delivery/module/user/model"
)

type RealStore interface {
	FindUserByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		relations ...string,
	) (*usermodel.User, error)
}

type userCaching struct {
	store     Caching
	realStore RealStore
}

func NewUserCaching(store Caching, realStore RealStore) *userCaching {
	return &userCaching{
		store:     store,
		realStore: realStore,
	}
}

func (c *userCaching) FindUserByCondition(ctx context.Context, conditions map[string]interface{}, relations ...string) (*usermodel.User, error) {
	userId := conditions["id"].(int)
	key := fmt.Sprintf("user-%d", userId)
	userInCache := c.store.Read(key)

	if userInCache != nil {
		return userInCache.(*usermodel.User), nil
	}

	user, err := c.realStore.FindUserByCondition(ctx, conditions, relations...)
	if err != nil {
		return nil, err
	}

	c.store.Write(key, user)
	return user, nil
}
