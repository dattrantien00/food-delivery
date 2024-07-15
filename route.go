package main

import (
	"food-delivery/component/appctx"
	"food-delivery/memcache"
	"food-delivery/middleware"
	"food-delivery/module/restaurant/transport/ginrestaurant"
	"food-delivery/module/restaurantlike/transport/ginrstlike"
	"food-delivery/module/upload/transport/ginupload"
	userstorage "food-delivery/module/user/storage"
	"food-delivery/module/user/transport/ginuser"

	"github.com/gin-gonic/gin"
)

func setRoute(appCtx appctx.AppContext, g *gin.Engine) {
	userStorage := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
	userCaching := memcache.NewUserCaching(memcache.NewCache(), userStorage)

	g.StaticFile("/demo", "./demo.html")
	v1 := g.Group("/v1")
	restaurant := v1.Group("restaurant", middleware.RequireAuth(appCtx,userCaching))
	{
		restaurant.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurant.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurant.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))

		restaurant.POST("/:id/liked-users", ginrstlike.LikeRestaurant(appCtx))
		restaurant.DELETE("/:id/liked-users", ginrstlike.UnLikeRestaurant(appCtx))
		restaurant.GET("/:id/liked-users", ginrstlike.ListUsers(appCtx))
	}

	v1.POST("/upload", ginupload.UploadImage(appCtx))
	v1.POST("/register", ginuser.RegisterUser(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))
	v1.GET("/profile", ginuser.Profile(appCtx))

	admin := v1.Group("/admin", middleware.RequireAuth(appCtx,userCaching), middleware.RoleRequire(appCtx, "admin"))
	{
		admin.GET("/", ginuser.Profile(appCtx))
	}
}
