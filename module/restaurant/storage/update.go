package restaurantstorage

import (
	"context"
	"food-delivery/common"
	restaurantmodel "food-delivery/module/restaurant/model"

	"gorm.io/gorm"
)

func (s *sqlStore) IncreaseLikedCount(context context.Context, id int) error {
	db := s.db
	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id =?", id).Update(
		"liked_count", gorm.Expr("liked_count+?", 1)).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}


func (s *sqlStore) DecreaseLikedCount(context context.Context, id int) error {
	db := s.db
	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id =?", id).Update(
		"liked_count", gorm.Expr("liked_count-?", 1)).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
