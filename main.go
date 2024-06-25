package main

import (
	"food-delivery/component/appctx"
	uploadprovider "food-delivery/component/provider"
	"food-delivery/middleware"
	"food-delivery/module/restaurant/transport/ginrestaurant"
	"food-delivery/module/upload/transport/ginupload"
	"food-delivery/module/user/transport/ginuser"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:Quajvat12@@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3ApiKey := os.Getenv("S3ApiKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")

	// jwtSecretKey := os.Getenv()
	secretKey := os.Getenv("SYSTEM_SECRET")

	// fmt.Println(s3BucketName, s3Region, s3ApiKey)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = db.Debug()

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3ApiKey, s3SecretKey, s3Domain)
	appCtx := appctx.NewAppContext(db, s3Provider, secretKey)
	g := gin.Default()
	g.Use(middleware.Recover(appCtx), )

	v1 := g.Group("/v1")
	v1.GET("restaurant", ginrestaurant.ListRestaurant(appCtx))
	v1.POST("restaurant", ginrestaurant.CreateRestaurant(appCtx))
	v1.DELETE("restaurant/:id", ginrestaurant.DeleteRestaurant(appCtx))

	v1.POST("/upload", ginupload.UploadImage(appCtx))
	v1.POST("/register", ginuser.RegisterUser(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))
	v1.GET("/profile", middleware.RequireAuth(appCtx),ginuser.Profile(appCtx))
	g.Run()
	// fmt.Println(os.Getenv("BUCKET_NAME"))
}
