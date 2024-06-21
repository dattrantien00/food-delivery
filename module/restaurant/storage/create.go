package restaurantstorage

import (
	"context"
	restaurantmodel "food-delivery/module/restaurant/model"
)

func (s *sqlStore) Create(context context.Context, data *restaurantmodel.RestaurantCreate) error{
	return s.db.Create(&data).Error
}