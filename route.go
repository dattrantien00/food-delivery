package main

import (
	"food-delivery/component/appctx"
	"food-delivery/middleware"
	"food-delivery/module/restaurant/transport/ginrestaurant"
	"food-delivery/module/upload/transport/ginupload"
	"food-delivery/module/user/transport/ginuser"

	"github.com/gin-gonic/gin"
)

func setRoute(appCtx appctx.AppContext, g *gin.Engine) {
	v1 := g.Group("/v1")
	restaurant := v1.Group("restaurant", middleware.RequireAuth(appCtx))
	{
		restaurant.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurant.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurant.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
	}

	v1.POST("/upload", ginupload.UploadImage(appCtx))
	v1.POST("/register", ginuser.RegisterUser(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))
	v1.GET("/profile", ginuser.Profile(appCtx))

	admin := v1.Group("/admin", middleware.RequireAuth(appCtx), middleware.RoleRequire(appCtx, "admin"))
	{
		admin.GET("/", ginuser.Profile(appCtx))
	}
}
