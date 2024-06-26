package restaurantlikestorage

import (
	"context"
	"food-delivery/common"
	restaurantlikemodel "food-delivery/module/restaurantlike/model"
)

const timeLayout = "2006-01-02T15:04:05.999999"

func (s *sqlStore) GetRestaurantLikes(context context.Context, ids []int) (
	map[int]int, error) {

	result := map[int]int{}

	type sqlData struct {
		RestaurantId int `gorm:"column:restaurant_id"`
		LikeCount    int `gorm:"column:count"`
	}

	var listLike []sqlData
	if err := s.db.Table(restaurantlikemodel.Like{}.TableName()).
		Select("restaurant_id,count(restaurant_id) AS count").
		Where("restaurant_id in (?)", ids).Group("restaurant_id").
		Find(&listLike).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	for _, item := range listLike {
		result[item.RestaurantId] = item.LikeCount
	}
	return result, nil
}
