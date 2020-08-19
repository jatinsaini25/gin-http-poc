package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jatinsaini25/gin-http-poc/controllers"
	"github.com/jatinsaini25/gin-http-poc/middlewares"
	"github.com/jatinsaini25/gin-http-poc/repository"
	"github.com/jatinsaini25/gin-http-poc/services"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoRepo       repository.VideoRepository  = repository.NewVideoRepository()
	videoService    services.VideoService       = services.New(videoRepo)
	videoController controllers.VideoController = controllers.New(videoService)
)

func setLogOutput() {
	f, err := os.Create("gin.log")
	if err != nil {
		return
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	defer videoRepo.Close()
	server := gin.New()

	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	setLogOutput()

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)

			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video Input Object is valid!"})
			}
		})

		apiRoutes.PUT("/videos/:id", func(ctx *gin.Context) {
			err := videoController.Update(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video Input Object is valid!"})
			}
		})

		apiRoutes.DELETE("/videos/:id", func(ctx *gin.Context) {
			err := videoController.Delete(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video Input Object is valid!"})
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	server.Run(":8000")
}
