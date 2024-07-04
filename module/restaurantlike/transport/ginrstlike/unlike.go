package ginrstlike

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	restaurantstorage "food-delivery/module/restaurant/storage"
	restaurantlikebiz "food-delivery/module/restaurantlike/biz"
	restaurantlikestorage "food-delivery/module/restaurantlike/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UnLikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(err)
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		// var data = restaurantlikemodel.Like{
		// 	RestaurantId: int(uid.GetLocalID()),
		// 	UserId:       requester.GetUserId(),
		// }

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		likeStore := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantlikebiz.NewUserUnLikeRestaurant(store,likeStore)
		if err := biz.UnLikeRestaurant(c.Request.Context(), requester.GetUserId(), int(uid.GetLocalID())); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		// data.User.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
