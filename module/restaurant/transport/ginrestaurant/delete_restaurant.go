package ginrestaurant

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	restaurantbiz "food-delivery/module/restaurant/biz"
	restaurantstorage "food-delivery/module/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// id, err := strconv.Atoi(c.Param("id"))

		// if err != nil {

		// 	c.JSON(http.StatusBadRequest, err)
		// 	return
		// }
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {

			c.JSON(http.StatusBadRequest, common.ErrInternal(err))
			return
		}
		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)
		if err := biz.DeleteRestaurant(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

// func CreateRestaurant(c *gin.Context) {

// }
