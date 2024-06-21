package ginrestaurant

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	restaurantbiz "food-delivery/module/restaurant/biz"
	restaurantmodel "food-delivery/module/restaurant/model"
	restaurantstorage "food-delivery/module/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBindJSON(&data); err != nil {

			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewCreateRestaurantBiz(store)
		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId))
	}
}

// func CreateRestaurant(c *gin.Context) {

// }
