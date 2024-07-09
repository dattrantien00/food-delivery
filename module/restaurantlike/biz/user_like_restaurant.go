package restaurantlikebiz

import (
	"context"
	"food-delivery/common"
	restaurantlikemodel "food-delivery/module/restaurantlike/model"
	"food-delivery/pubsub"
	"log"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

// type IncLikedCountResStore interface {
// 	IncreaseLikedCount(context context.Context, id int) error
// }
type userLikeRestaurantBiz struct {
	store UserLikeRestaurantStore
	// likeStore IncLikedCountResStore
	pubsub pubsub.Pubsub
}

func NewUserLikeRestaurant(store UserLikeRestaurantStore,
	// likeStore IncLikedCountResStore,
	pubsub pubsub.Pubsub) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{
		store: store,
		// likeStore: likeStore,
		pubsub: pubsub,
	}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(ctx context.Context, data *restaurantlikemodel.Like) error {
	err := biz.store.Create(ctx, data)
	if err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	err = biz.pubsub.Publish(ctx, common.TopicUserLikeRestaurant, pubsub.NewMessage(data))
	if err != nil {
		log.Println(err)
	}
	// j := asyncjob.NewJob(func(ctx context.Context) error {
	// 	return biz.likeStore.IncreaseLikedCount(ctx, data.RestaurantId)
	// })

	// asyncjob.NewGroup(true, j).Run(ctx)

	return nil
}
