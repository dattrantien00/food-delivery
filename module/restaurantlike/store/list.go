package restaurantlikestorage

import (
	"context"
	"food-delivery/common"
	restaurantlikemodel "food-delivery/module/restaurantlike/model"
	"time"

	"github.com/btcsuite/btcutil/base58"
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

func (s *sqlStore) GetUsersLikedRestaurant(context context.Context, conditions map[string]interface{},
	paging *common.Paging,
	filter *restaurantlikemodel.Filter,
	moreKeys ...string) (
	[]common.SimpleUser, error) {

	var result []restaurantlikemodel.Like
	db := s.db.Table(restaurantlikemodel.Like{}.TableName()).Where(conditions)
	if f := filter; f != nil {
		if f.RestaurantId > 0 {
			db = db.Where("restaurant_id=?", f.RestaurantId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	db = db.Preload("User")
	if v := paging.FakeCursor; v != "" {
		timeCreated, err := time.Parse(timeLayout, string(base58.Decode(v)))
		if err != nil {
			return nil, common.ErrDb(err)
		}
		db = db.Where("created_at <?", timeCreated.Format("2006-01-02 15:04:05"))
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.Limit(paging.Limit).Order("created_at desc").Find(&result).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	users := make([]common.SimpleUser, len(result))
	for i, item := range result {
		item.User.UpdatedAt = nil
		item.User.CreatedAt = item.CreatedAt
		users[i] = *item.User

		if i == len(result)-1 {
			item.User.Mask()
			paging.NextCursor = base58.Encode([]byte(item.CreatedAt.Format(timeLayout)))
		}
	}
	return users, nil
}
