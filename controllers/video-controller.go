package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/juliofeng/GinGateway/models"
	"sync"
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

type generator struct {
	counter int
	mtx     sync.Mutex
}

func (g *generator) getNextId() int {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	g.counter++
	return g.counter
}

var g *generator = &generator{}

func (c *controller) GetAll(context *gin.Context) {
	context.JSON(200, c.videos)
}

func (c *controller) Update(context *gin.Context) {
	var videoToUpdate models.Video
	if err := context.ShouldBindUri(&videoToUpdate); err != nil {
		context.String(400, "bad request %v", err)
		return
	}
	if err := context.ShouldBind(&videoToUpdate); err != nil {
		context.String(400, "bad request %v", err)
		return
	}
	for idx, video := range c.videos {
		if video.Id == videoToUpdate.Id {
			c.videos[idx] = videoToUpdate
			context.String(200, "success, video with id %d has been updated", videoToUpdate.Id)
			return
		}
	}
	context.String(400, "bad request cannot find video with %d to update", videoToUpdate.Id)
}

func (c *controller) Create(context *gin.Context) {
	video := models.Video{Id: g.getNextId()}
	if err := context.ShouldBindJSON(&video); err != nil {
		context.String(400, "bad request %v", err)
		return
	}
	c.videos = append(c.videos, video)
	context.String(200, "success, new video id is %d", video.Id)
}

func (c *controller) Delete(context *gin.Context) {
	videoToDelete := models.Video{}
	if err := context.ShouldBindUri(&videoToDelete); err != nil {
		context.String(400, "bad request %v", err)
		return
	}
	for idx, video := range c.videos {
		if video.Id == videoToDelete.Id {
			c.videos = append(c.videos[:idx], c.videos[idx+1:]...)
			context.String(200, "success, delete video id is %d", video.Id)
			return
		}
	}
	context.String(400, "bad request, video id %d not found", videoToDelete.Id)
}
