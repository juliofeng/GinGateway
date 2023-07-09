package controllers

import (
	"github.com/HanFa/learn-go/gin-example/models"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	GetAll(context *gin.Context)
	Update(context *gin.Context)
	Create(context *gin.Context)
	Delete(context *gin.Context)
}

type controller struct {
	videos []models.Video
}

func NewVideoController() VideoController {
	return &controller{videos: make([]models.Video, 0)}
}
