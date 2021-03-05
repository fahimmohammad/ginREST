package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/haquenafeem/boilerplate-gin/auth"
	"github.com/haquenafeem/boilerplate-gin/todo"
)

func main() {
	gin.SetMode(gin.DebugMode)

	router := gin.New()  // Gin instance
	router.Use(Logger()) // loger
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))

	v1 := router.Group("/api/v1")

	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		fmt.Println("Cannot connect to database......")
		return
	}

	fmt.Println("Server Started.......")
	initializeAllServices(v1, session)

	router.Run(":3002")
}

func initializeAllServices(router *gin.RouterGroup, dbSession *mgo.Session) {
	todo.Init(router, dbSession)
	auth.Init(router, dbSession)
}

// Logger - middleware
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("======================================>")
		fmt.Println("Url Hit : " + c.Request.URL.String() + " Method : " + c.Request.Method)
		c.Next()
		since := time.Since(t)
		fmt.Println("Time Took : " + since.String())
	}
}
