package main

import (
	"github.com/gin-gonic/gin"
	"github.com/juliofeng/GinGateway/controllers"
	"log"
)

func main() {
	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.Static("/resources", "./resources")
	server.StaticFile("/mycoolvideo", "./resources/Disney.mp4")

	videoController := controllers.NewVideoController()
	videoGroup := server.Group("/video")
	videoGroup.GET("/", videoController.GetAll)
	videoGroup.PUT("/:id", videoController.Update)
	videoGroup.POST("/", videoController.Create)
	videoGroup.DELETE("/:id", videoController.Delete)

	log.Fatalln(server.Run("localhost:8080"))
}
