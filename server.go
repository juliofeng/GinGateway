package main

import (
	"github.com/gin-gonic/gin"
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

	videoController := NewVideoController()

	log.Fatalln(server.Run("localhost:8080"))
}
