package main

import (
	"github.com/gin-gonic/gin"
	"github.com/juliofeng/GinGateway/controllers"
	"github.com/juliofeng/GinGateway/middlewares"
	"log"
)

func main() {
	server := gin.Default()
	//server.Use(middlewares.MyAuth())

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.Static("/resources", "./resources")
	server.StaticFile("/mycoolvideo", "./resources/Disney.mp4")

	videoController := controllers.NewVideoController()
	videoGroup := server.Group("/videos")
	videoGroup.Use(middlewares.MyLogger())

	videoGroup.GET("/", videoController.GetAll)
	videoGroup.POST("/", videoController.Create)
	videoGroup.PUT("/:id", videoController.Update)
	videoGroup.DELETE("/:id", videoController.Delete)

	log.Fatalln(server.Run("localhost:8080"))
}
