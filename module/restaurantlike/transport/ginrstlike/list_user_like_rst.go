package ginrstlike

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	restaurantlikebiz "food-delivery/module/restaurantlike/biz"
	restaurantlikemodel "food-delivery/module/restaurantlike/model"
	restaurantlikestorage "food-delivery/module/restaurantlike/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListUsers(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}
		pagingData.FulFill()

		var filter restaurantlikemodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		store := restaurantlikestorage.NewSQLStore(db)
		// likeStore := restaurantlikestorage.NewSQLStore(db)
		// repo := restaurantrepo.NewListRestaurantRepo(store, likeStore)
		biz := restaurantlikebiz.NewListUserLikeRestaurant(store)
		data, err := biz.ListUsersLikeRestaurant(c.Request.Context(), &pagingData, &filter)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		for i := range data {
			data[i].Mask()
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(data, pagingData, filter))
	}
}
