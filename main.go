package main

import (
	"food-delivery/component/appctx"
	"food-delivery/middleware"
	"food-delivery/module/restaurant/transport/ginrestaurant"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:Quajvat12@@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = db.Debug()
	appCtx := appctx.NewAppContext(db)
	g := gin.Default()
	g.Use(middleware.Recover(appCtx))

	g.GET("restaurant", ginrestaurant.ListRestaurant(appCtx))
	g.POST("restaurant", ginrestaurant.CreateRestaurant(appCtx))
	g.DELETE("restaurant/:id", ginrestaurant.DeleteRestaurant(appCtx))
	g.Run()
}
