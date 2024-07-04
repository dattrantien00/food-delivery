package restaurantlikebiz

import (
	"context"
	"food-delivery/common"
	restaurantlikemodel "food-delivery/module/restaurantlike/model"
)

type listUserLikeRestaurantStore interface {
	GetUsersLikedRestaurant(context context.Context, conditions map[string]interface{},
		paging *common.Paging,
		filter *restaurantlikemodel.Filter,
		moreKeys ...string) (
		[]common.SimpleUser, error)
}

type listUserLikeRestaurantBiz struct {
	store listUserLikeRestaurantStore
}

func NewListUserLikeRestaurant(store listUserLikeRestaurantStore) *listUserLikeRestaurantBiz {
	return &listUserLikeRestaurantBiz{
		store: store,
	}
}

func (biz *listUserLikeRestaurantBiz) ListUsersLikeRestaurant(context context.Context,
	paging *common.Paging,
	filter *restaurantlikemodel.Filter) (
	[]common.SimpleUser, error) {
	return biz.store.GetUsersLikedRestaurant(context, nil, paging, filter)
}
