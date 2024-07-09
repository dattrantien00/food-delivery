package restaurantlikebiz

import (
	"context"
	"food-delivery/common"
	restaurantlikemodel "food-delivery/module/restaurantlike/model"
	"food-delivery/pubsub"
	"log"
)

type UserUnLikeRestaurantStore interface {
	Delete(context context.Context, userId, restaurantId int) error
}

type userUnLikeRestaurantBiz struct {
	store  UserUnLikeRestaurantStore
	pubsub pubsub.Pubsub
}

func NewUserUnLikeRestaurant(store UserUnLikeRestaurantStore, pubsub pubsub.Pubsub) *userUnLikeRestaurantBiz {
	return &userUnLikeRestaurantBiz{
		store:  store,
		pubsub: pubsub,
	}
}

func (biz *userUnLikeRestaurantBiz) UnLikeRestaurant(ctx context.Context, userId, restaurantId int) error {
	err := biz.store.Delete(ctx, userId, restaurantId)
	if err != nil {
		return restaurantlikemodel.ErrCannotUnLikeRestaurant(err)
	}

	err = biz.pubsub.Publish(ctx, common.TopicUserUnLikeRestaurant, pubsub.NewMessage(&restaurantlikemodel.Like{
		RestaurantId: restaurantId,
		UserId:       userId,
		// CreatedAt:    (time.Now().Local()),
	}))
	if err != nil {
		log.Println(err)
	}
	return nil
}
