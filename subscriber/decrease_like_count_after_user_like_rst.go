package subscriber

import (
	"context"
	"food-delivery/component/appctx"
	restaurantstorage "food-delivery/module/restaurant/storage"
	"food-delivery/pubsub"
)

// type HasRestaurantId interface {
// 	GetRestaurantId() int
// }

func DecreaseLikeCountAfterUserLikeRestaurant(appctx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Decrease like",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appctx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)
			return store.DecreaseLikedCount(ctx, likeData.GetRestaurantId())
		},
	}
}
