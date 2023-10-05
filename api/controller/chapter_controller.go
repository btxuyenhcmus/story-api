package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/readtruyen/go-novelstory-api/crawler"
	"github.com/readtruyen/go-novelstory-api/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChapterController struct {
	ChapterUseCase domain.ChapterUseCase
}

func (cc *ChapterController) Create(c *gin.Context) {
	storyID := c.Param("id")
	var createChapter domain.CreateChapter

	err := c.ShouldBind(&createChapter)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	chapter := domain.Chapter{
		Title:         createChapter.Title,
		SourceID:      createChapter.SourceID,
		ChapterNextID: createChapter.ChapterNextID,
		ChapterPrevID: createChapter.ChapterPrevID,
		Content:       createChapter.Content,
		Number:        createChapter.Number,
	}

	chapter.ID = primitive.NewObjectID()
	chapter.StoryID, err = primitive.ObjectIDFromHex(storyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	chapter.ChapterPrev, err = primitive.ObjectIDFromHex(createChapter.ChapterPrev)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = cc.ChapterUseCase.Create(c, &chapter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Chapter created successfully",
	})
}

func (cc *ChapterController) FetchByStoryId(c *gin.Context) {
	storyID := c.Param("id")
	chapters, err := cc.ChapterUseCase.GetChaptersByStoryId(c, storyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.StoryDetailResponse{
		Status:  http.StatusOK,
		Message: "Get list chapter of story successfully",
		Data:    chapters,
	})
}

func (cc *ChapterController) FetchByStoryIdV1(c *gin.Context) {
	storyStr := c.Param("id")
	pageStr := c.Query("page")

	storyID, err := strconv.Atoi(storyStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	response, err := crawler.FetchChapter(storyID, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (cc *ChapterController) DownloadByStoryId(c *gin.Context) {
	storyID := c.Param("id")
	chapters, err := cc.ChapterUseCase.DownloadChapterByStoryId(c, storyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.StoryDetailResponse{
		Status:  http.StatusOK,
		Message: "Get list chapter of story successfully",
		Data:    chapters,
	})
}

func (cc *ChapterController) DownloadByStoryIdV1(c *gin.Context) {
	storyStr := c.Param("id")
	pageStr := c.Query("page")

	storyID, err := strconv.Atoi(storyStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	response, err := crawler.DownloadChapter(storyID, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (cc *ChapterController) FetchByChapterId(c *gin.Context) {
	chapterID := c.Param("id")
	chapter, err := cc.ChapterUseCase.GetChapterById(c, chapterID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.StoryDetailResponse{
		Status:  http.StatusOK,
		Message: "Get detail successfully",
		Data:    chapter,
	})
}
