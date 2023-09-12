package controller

import (
	"net/http"
	"readtruyen-api/domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StoryController struct {
	StoryUseCase domain.StoryUseCase
}

func (sc *StoryController) Create(c *gin.Context) {
	var story domain.Story

	err := c.ShouldBind(&story)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	story.ID = primitive.NewObjectID()

	err = sc.StoryUseCase.Create(c, &story)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Story created successfully",
	})
}

func (sc *StoryController) FetchList(c *gin.Context) {
	stories, err := sc.StoryUseCase.GetStories(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.StorySuccessResponse{
		Status:  http.StatusOK,
		Message: "Get list story successfully",
		Data:    stories,
	})
}

func (sc *StoryController) FetchByStoryID(c *gin.Context) {
	storyID := c.Param("id")
	story, err := sc.StoryUseCase.GetStoryById(c, storyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.StorySuccessResponse{
		Status:  http.StatusOK,
		Message: "Get detail successfully",
		Data:    story,
	})
}
