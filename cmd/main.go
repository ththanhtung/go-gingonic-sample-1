package main

import (
	"example.com/sample1/controllers"
	"example.com/sample1/services"
	"github.com/gin-gonic/gin"
)

var (
	videoService services.VideoService = services.New()
	VideoController controllers.VideoController = controllers.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context){
		ctx.JSON(200, VideoController.Save(ctx))
	})

	server.Run(":3000")
}