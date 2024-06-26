package ginrestaurant

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	restaurantbiz "food-delivery/module/restaurant/biz"
	restaurantmodel "food-delivery/module/restaurant/model"
	restaurantrepo "food-delivery/module/restaurant/repository"
	restaurantstorage "food-delivery/module/restaurant/storage"
	restaurantlikestorage "food-delivery/module/restaurantlike/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}
		pagingData.FulFill()

		var filter restaurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		store := restaurantstorage.NewSQLStore(db)
		likeStore := restaurantlikestorage.NewSQLStore(db)
		repo := restaurantrepo.NewListRestaurantRepo(store, likeStore)
		biz := restaurantbiz.NewListRestaurantBiz(repo)
		data, err := biz.ListRestaurant(c.Request.Context(), &filter, &pagingData)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		for i := range data {
			data[i].Mask(false)
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(data, pagingData, filter))
	}
}
