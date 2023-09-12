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
	group.GET("/stories", sc.FetchList)
	group.POST("/story", sc.Create)
	group.GET("/story/:id", sc.FetchByStoryID)
}
