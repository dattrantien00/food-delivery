package restaurantstorage

import (
	"context"
	"food-delivery/common"
	restaurantmodel "food-delivery/module/restaurant/model"

	"go.opencensus.io/trace"
)

func (s *sqlStore) ListDataWithCondition(context context.Context, filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string) (
	[]restaurantmodel.Restaurant, error) {

	var data []restaurantmodel.Restaurant

	db := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("status in (1)")

	if f := filter; f != nil {
		if f.OwnerId > 0 {
			db = db.Where("user_id=?", f.OwnerId)
		}
	}
	_, span := trace.StartSpan(context, "storagerestaurant.count_restaurant")
	if err := db.Count(&paging.Total).Error; err != nil {
		span.End()
		return nil, common.ErrDb(err)
	}
	span.End()
	if v := paging.FakeCursor; v != "" {
		uid, err := common.FromBase58(v)
		if err != nil {
			return nil, common.ErrDb(err)
		}
		db = db.Where("id<?", uid.GetLocalID())
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)
	}
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	_, span1 := trace.StartSpan(context, "storagerestaurant.list_restaurant")
	defer span1.End()
	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&data).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	if len(data) > 0 {
		last := data[len(data)-1]
		last.Mask(false)
		paging.NextCursor = last.FakeId.String()

	}
	return data, nil
}
