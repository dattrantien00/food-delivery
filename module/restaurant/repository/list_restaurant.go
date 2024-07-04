package restaurantrepo

import (
	"context"
	"food-delivery/common"
	restaurantmodel "food-delivery/module/restaurant/model"
)

type listRestaurantStore interface {
	ListDataWithCondition(context context.Context, filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string) (
		[]restaurantmodel.Restaurant, error)
}

//	type likeRestaurantStore interface {
//		GetRestaurantLikes(context context.Context, ids []int) (
//			map[int]int, error)
//	}
type listRestaurantRepo struct {
	store listRestaurantStore
	// likerestaurantstore likeRestaurantStore
}

func NewListRestaurantRepo(store listRestaurantStore) *listRestaurantRepo {
	return &listRestaurantRepo{
		store: store,
		// likerestaurantstore: likerestaurantstore,
	}
}

func (biz *listRestaurantRepo) ListRestaurant(context context.Context, filter *restaurantmodel.Filter,
	paging *common.Paging) ([]restaurantmodel.Restaurant, error) {
	restaurants, err := biz.store.ListDataWithCondition(context, filter, paging, "User")
	if err != nil {
		return nil, err
	}

	// ids := make([]int, len(restaurants))

	// for i := range ids {
	// 	ids[i] = restaurants[i].Id
	// }
	// likeMap, err := biz.likerestaurantstore.GetRestaurantLikes(context, ids)
	// if err != nil {
	// 	log.Println(err)
	// 	return restaurants, nil
	// }

	// for i, res := range restaurants {
	// 	restaurants[i].LikedCount = likeMap[res.Id]
	// }
	return restaurants, nil
}
