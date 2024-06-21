package restaurantbiz

import (
	"context"
	"food-delivery/common"
	restaurantmodel "food-delivery/module/restaurant/model"
)

type listRestaurantStore interface {
	ListDataWithCondition(context context.Context, filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string) (
		[]restaurantmodel.Restaurant, error)
}
type listRestaurantBiz struct {
	store listRestaurantStore
}

func NewListRestaurantBiz(store listRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{
		store: store,
	}
}

func (biz *listRestaurantBiz) ListRestaurant(context context.Context, filter *restaurantmodel.Filter,
	paging *common.Paging) ([]restaurantmodel.Restaurant, error) {
	return biz.store.ListDataWithCondition(context, filter, paging)
}
