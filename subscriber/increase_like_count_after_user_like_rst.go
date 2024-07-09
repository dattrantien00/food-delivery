package subscriber

import (
	"context"
	"food-delivery/component/appctx"
	restaurantstorage "food-delivery/module/restaurant/storage"
	"food-delivery/pubsub"
)

type HasRestaurantId interface {
	GetRestaurantId() int
}

// func IncreaseLikeCountAfterUserLikeRestaurant(appctx appctx.AppContext, ctx context.Context) {
// 	c, _ := appctx.GetPubsub().Subscribe(ctx, common.TopicUserLikeRestaurant)

// 	store := restaurantstorage.NewSQLStore(appctx.GetMainDBConnection())

// 	go func() {
// 		defer common.AppRecover()
// 		for {
// 			msg := <-c
// 			// fmt.Println("==============", msg.Data())
// 			likeData := msg.Data().(HasRestaurantId)
// 			_ = store.IncreaseLikedCount(ctx, likeData.GetRestaurantId())
// 		}
// 	}()
// }

func IncreaseLikeCountAfterUserLikeRestaurant(appctx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Increase Like",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appctx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)
			return store.IncreaseLikedCount(ctx, likeData.GetRestaurantId())
		},
	}
}
