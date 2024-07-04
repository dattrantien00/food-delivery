package restaurantlikebiz

import (
	"context"
	"food-delivery/component/asyncjob"
	restaurantlikemodel "food-delivery/module/restaurantlike/model"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

type IncLikedCountResStore interface {
	IncreaseLikedCount(context context.Context, id int) error
}
type userLikeRestaurantBiz struct {
	store     UserLikeRestaurantStore
	likeStore IncLikedCountResStore
}

func NewUserLikeRestaurant(store UserLikeRestaurantStore, likeStore IncLikedCountResStore) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{
		store:     store,
		likeStore: likeStore,
	}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(ctx context.Context, data *restaurantlikemodel.Like) error {
	err := biz.store.Create(ctx, data)
	if err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}
	j := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.likeStore.IncreaseLikedCount(ctx, data.RestaurantId)
	})

	asyncjob.NewGroup(true, j).Run(ctx)

	return nil
}
