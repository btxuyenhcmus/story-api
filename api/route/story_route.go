package route

import (
	"time"

	"github.com/readtruyen/go-novelstory-api/api/controller"
	"github.com/readtruyen/go-novelstory-api/bootstrap"
	"github.com/readtruyen/go-novelstory-api/domain"
	"github.com/readtruyen/go-novelstory-api/mongo"
	"github.com/readtruyen/go-novelstory-api/repository"
	"github.com/readtruyen/go-novelstory-api/usecase"

	"github.com/gin-gonic/gin"
)

func NewStoryRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	sr := repository.NewStoryRepository(db, domain.CollectionStory)
	sc := &controller.StoryController{
		StoryUseCase: usecase.NewStoryUseCase(sr, timeout),
	}

	// Final version router
	group.GET("/stories", sc.FetchList)

	// Version 1 router
	group.GET("/v1/stories", sc.FetchListV1)
	group.GET("/v1/story/:id/related", sc.FetchRelatedByStoryIdV1)
}

func NewUserStoryRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	sr := repository.NewStoryRepository(db, domain.CollectionStory)
	sc := &controller.StoryController{
		StoryUseCase: usecase.NewStoryUseCase(sr, timeout),
	}

	// Final version router
	group.GET("/story/:id", sc.FetchByStoryId)

	// Version 1 router
	group.GET("/v1/story/:id", sc.FetchByStoryIdV1)
}

func NewPrivateStoryRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	sr := repository.NewStoryRepository(db, domain.CollectionStory)
	sc := &controller.StoryController{
		StoryUseCase: usecase.NewStoryUseCase(sr, timeout),
	}

	// Final version
	group.POST("/story", sc.Create)
}
