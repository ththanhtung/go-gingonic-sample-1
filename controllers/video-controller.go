package controllers

import (
	"example.com/sample1/entity"
	"example.com/sample1/services"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	Save(ctx *gin.Context) entity.Video
	FindAll() []entity.Video
}

type controller struct {
	service services.VideoService
}

func New(service services.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}
