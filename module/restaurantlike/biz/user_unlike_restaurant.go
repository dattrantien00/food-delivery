package restaurantlikebiz

import (
	"context"
	"food-delivery/component/asyncjob"
	restaurantlikemodel "food-delivery/module/restaurantlike/model"
)

type UserUnLikeRestaurantStore interface {
	Delete(context context.Context, userId, restaurantId int) error
}

type DecLikeRestaurantStore interface {
	DecreaseLikedCount(context context.Context, id int) error
}
type userUnLikeRestaurantBiz struct {
	store     UserUnLikeRestaurantStore
	likeStore DecLikeRestaurantStore
}

func NewUserUnLikeRestaurant(store UserUnLikeRestaurantStore, likeStore DecLikeRestaurantStore) *userUnLikeRestaurantBiz {
	return &userUnLikeRestaurantBiz{
		store:     store,
		likeStore: likeStore,
	}
}

func (biz *userUnLikeRestaurantBiz) UnLikeRestaurant(ctx context.Context, userId, restaurantId int) error {
	err := biz.store.Delete(ctx, userId, restaurantId)
	if err != nil {
		return restaurantlikemodel.ErrCannotUnLikeRestaurant(err)
	}

	j := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.likeStore.DecreaseLikedCount(ctx, restaurantId)
	})

	asyncjob.NewGroup(true, j).Run(ctx)
	return nil
}
