package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/readtruyen/go-novelstory-api/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CarouselController struct {
	CarouselUseCase domain.CarouselUseCase
}

func (cc *CarouselController) Create(c *gin.Context) {
	var carousel domain.Carousel

	err := c.ShouldBind(&carousel)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	carousel.ID = primitive.NewObjectID()
	err = cc.CarouselUseCase.Create(c, &carousel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Carousel created successfully",
	})
}

func (cc *CarouselController) FetchList(c *gin.Context) {
	carousels, err := cc.CarouselUseCase.GetCarousels(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.StoryDetailResponse{
		Status:  http.StatusOK,
		Message: "Get list carousel successfully",
		Data:    carousels,
	})
}
