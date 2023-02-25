package main

import (
	"net/http"

	"example.com/sample1/controllers"
	"example.com/sample1/middlewares"
	"example.com/sample1/services"
	"github.com/gin-gonic/gin"
)

var (
	videoService services.VideoService = services.New()
	VideoController controllers.VideoController = controllers.New(videoService)
)

func main() {
	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context){ 
		err := VideoController.Save(ctx)

		if (err != nil){
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}else{
			ctx.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
		}
	})

	server.Run(":3000")
}