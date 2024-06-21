package restaurantstorage

import (
	"context"
	"food-delivery/common"
	restaurantmodel "food-delivery/module/restaurant/model"
)

func (s *sqlStore) ListDataWithCondition(context context.Context, filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string) (
	[]restaurantmodel.Restaurant, error) {

	var data []restaurantmodel.Restaurant

	db := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("status in (1)")

	if f := filter; f != nil {
		if f.OwnerId > 0 {
			db = db.Where("owner_id=?", f.OwnerId)
		}
	}
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	offset := (paging.Page - 1) * paging.Limit
	if err := db.
		Offset(offset).
		Limit(paging.Limit).
		Order("id desc").
		Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
