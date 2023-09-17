package controller

import (
	"net/http"
	"strconv"

	"github.com/readtruyen/go-novelstory-api/crawler"
	"github.com/readtruyen/go-novelstory-api/domain"

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
	pageStr := c.Query("page")
	perPageStr := c.Query("per_page")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage <= 0 || perPage >= 100 {
		perPage = 10
	}

	stories, pagination, err := sc.StoryUseCase.GetStoriesPagination(c, page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.StoryListResponse{
		Status:  http.StatusOK,
		Message: "Get list story successfully",
		Data:    stories,
		Meta: domain.Meta{
			Pagination: pagination,
		},
	})
}

func (sc *StoryController) FetchListV1(c *gin.Context) {
	typeStr := c.Query("type")
	pageStr := c.Query("page")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	response, err := crawler.GetStoriesPagination(typeStr, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (sc *StoryController) FetchByStoryId(c *gin.Context) {
	storyID := c.Param("id")
	story, err := sc.StoryUseCase.GetStoryById(c, storyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.StoryDetailResponse{
		Status:  http.StatusOK,
		Message: "Get detail successfully",
		Data:    story,
	})
}
