package main

import (
	"fmt"
	"log"
	"os"
	"quan/go/component"
	"quan/go/middleware"
	"quan/go/modules/auth/authhdl"
	"quan/go/modules/restaurant/restauranttransport/ginrestaurant"
	"quan/go/modules/upload/uploadtransport/ginupload"


	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Can't load env file")
	}
	dbName := os.Getenv("MYSQL_DATABASE")
	dbUser := os.Getenv("MYSQL_USER")
	dbPW := os.Getenv("MYSQL_PASSWORD")
	dbPORT := os.Getenv("MYSQL_PORT")
	// secretKey := os.Getenv("SECRET_KEY")
	dns := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPW, dbPORT, dbName)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	runServer(db)

}

func runServer(db *gorm.DB) {
	AppCtx := component.NewAppContext(db)
	// gin.SetMode(gin.ReleaseMode

	r := gin.Default()
	r.Use(middleware.Recover(AppCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.POST("/upload", ginupload.Upload(AppCtx))
	v1 := r.Group("/v1")

	restaurant := v1.Group("/restaurants")
	{
		restaurant.GET("/:id", ginrestaurant.GetRestaurant(AppCtx))
		restaurant.DELETE("/:id", ginrestaurant.DeleteRestaurant(AppCtx))
		restaurant.POST("", ginrestaurant.CreateRestaurant(AppCtx))
		restaurant.GET("", ginrestaurant.ListRestaurant(AppCtx))
		restaurant.PATCH("/:id", ginrestaurant.UpdateRestaurant(AppCtx))

	}

	auth := v1.Group("/auth")
	{
		auth.POST("/register", authhdl.Register(AppCtx))

	}
	r.Run(":8080")

}
