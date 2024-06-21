package restaurantbiz

import (
	"context"
	restaurantmodel "food-delivery/module/restaurant/model"
)

type createRestaurantStore interface {
	Create(context context.Context, data *restaurantmodel.RestaurantCreate) error
}
type createRestaurantBiz struct {
	store createRestaurantStore
}

func NewCreateRestaurantBiz(store createRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{
		store: store,
	}
}

func (biz *createRestaurantBiz) CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	return biz.store.Create(context, data)
}
