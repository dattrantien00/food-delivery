package restaurantbiz

import (
	"context"
	"food-delivery/common"
	restaurantmodel "food-delivery/module/restaurant/model"

	"go.opencensus.io/trace"
)

type listRestaurantRepo interface {
	ListRestaurant(context context.Context, filter *restaurantmodel.Filter,
		paging *common.Paging) ([]restaurantmodel.Restaurant, error)
}

type likeRestaurantStore interface {
	GetRestaurantLikes(context context.Context, ids []int) (
		map[int]int, error)
}
type listRestaurantBiz struct {
	repo listRestaurantRepo
}

func NewListRestaurantBiz(repo listRestaurantRepo) *listRestaurantBiz {
	return &listRestaurantBiz{
		repo: repo,
	}
}

func (biz *listRestaurantBiz) ListRestaurant(context context.Context, filter *restaurantmodel.Filter,
	paging *common.Paging) ([]restaurantmodel.Restaurant, error) {
	ctx,span := trace.StartSpan(context, "biz.list_restaurant")
	defer span.End()
	result,err := biz.repo.ListRestaurant(ctx, filter, paging)
	if err != nil{
		return nil,err
	}
	
	return result,nil
}
