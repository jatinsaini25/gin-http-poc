package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jatinsaini25/gin-http-poc/entity"
	"github.com/jatinsaini25/gin-http-poc/services"
	"github.com/jatinsaini25/gin-http-poc/validators"
)

var validate *validator.Validate

type VideoController interface {
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	FindAll() []entity.Video
	ShowAll(ctx *gin.Context)
}

type Controller struct {
	service services.VideoService
}

func New(service services.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &Controller{
		service: service,
	}
}

func (c *Controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *Controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	err = validate.Struct(video)

	if err != nil {
		return err
	}

	c.service.Save(video)
	return nil
}

func (c *Controller) Update(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)

	if err != nil {
		return err
	}

	video.ID = id

	err = validate.Struct(video)

	if err != nil {
		return err
	}

	c.service.Update(video)
	return nil
}

func (c *Controller) Delete(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)

	if err != nil {
		return err
	}

	video.ID = id

	c.service.Delete(video)
	return nil
}

func (c *Controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()

	data := gin.H{
		"Title":  "Video Title",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
