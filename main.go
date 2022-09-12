package main

import (
	"fmt"
	"log"
	"os"
	appctx "quan/go/appCtx"
	"quan/go/modules/auth/authhdl"

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
	dns := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s", dbUser, dbPW, dbPORT, dbName)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	appCtx := appctx.NewAppContext(db)
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	v1 := r.Group("/v1")

	auth := v1.Group("/auth")

	auth.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "auth"})
	})
	auth.POST("/login", authhdl.Login(appCtx, "aa"))
	r.Run(":8080") 

}
