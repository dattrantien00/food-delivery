package restaurantbiz

import (
	"context"
	"errors"
	"food-delivery/common"
	restaurantmodel "food-delivery/module/restaurant/model"
	"testing"
)

type mockCreateStore struct{}

func (mockCreateStore) Create(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if data.Name == "Dat" {
		return common.ErrDb(errors.New("something went wrong in DB"))
	}
	return nil
}
func TestCreateRestaurantBiz(t *testing.T) {
	biz := NewCreateRestaurantBiz(mockCreateStore{})
	dataTest := &restaurantmodel.RestaurantCreate{Name: ""}
	err := biz.CreateRestaurant(context.Background(), dataTest)
	if err == nil || err.Error() != "invalid request" {
		t.Error("fail")
	}

	dataTest = &restaurantmodel.RestaurantCreate{Name: "Dat"}
	err = biz.CreateRestaurant(context.Background(), dataTest)
	if err == nil  {
		t.Error("fail")
	}

	dataTest = &restaurantmodel.RestaurantCreate{Name: "Dat1"}
	err = biz.CreateRestaurant(context.Background(), dataTest)
	if err != nil {
		t.Error("fail")
	}
}
