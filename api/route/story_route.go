package route

import (
	"readtruyen-api/api/controller"
	"readtruyen-api/bootstrap"
	"readtruyen-api/domain"
	"readtruyen-api/mongo"
	"readtruyen-api/repository"
	"readtruyen-api/usecase"
	"time"

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
