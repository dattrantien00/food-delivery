package restaurantlikestorage

import (
	"context"
	"food-delivery/common"
	restaurantlikemodel "food-delivery/module/restaurantlike/model"
)

func (s *sqlStore) Delete(context context.Context, userId, restaurantId int) error {

	if err := s.db.Table(restaurantlikemodel.Like{}.TableName()).
		Where("id=? and restaurant_id=?", userId, restaurantId).
		Delete(nil).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
