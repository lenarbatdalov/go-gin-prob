package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lenarbatdalov/learn/controller"
	"github.com/lenarbatdalov/learn/middlewares"
	"github.com/lenarbatdalov/learn/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	server := gin.New()

	server.Static("/css", "/templates/css")
	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK!!!",
		})
	})

	apiRoutes := server.Group("/api")
	{

		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Video Input is Valid!",
				})
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	server.Run(":8081")
}
