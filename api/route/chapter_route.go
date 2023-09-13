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

func NewChapterRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	cr := repository.NewChapterRepository(db, domain.CollectionChapter)
	cc := &controller.ChapterController{
		ChapterUseCase: usecase.NewChapterUseCase(cr, timeout),
	}
	group.POST("/story/:id/chapter", cc.Create)
	group.GET("/story/:id/chapters", cc.FetchByStoryId)
	group.GET("/story/:id/download", cc.DownloadByStoryId)
	group.GET("/chapter/:id", cc.FetchByChapterId)
}