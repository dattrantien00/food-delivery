package restaurantbiz

import (
	"context"
	"errors"
	restaurantmodel "food-delivery/module/restaurant/model"
)

type deleteRestaurantStore interface {
	FindDataWithCondition(context context.Context, condition map[string]interface{}, moreKeys ...string) (
		*restaurantmodel.Restaurant, error)
	Delete(context context.Context, id int) error
}
type deleteRestaurantBiz struct {
	store deleteRestaurantStore
}

func NewDeleteRestaurantBiz(store deleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{
		store: store,
	}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(context context.Context, id int) error {
	oldData,err := biz.store.FindDataWithCondition(context,map[string]interface{}{
		"id":id,
	})
	if err != nil{
		return err
	}

	if oldData.Status == 0 {
		return errors.New("data has been deleted")
	}
	return biz.store.Delete(context, id)
}