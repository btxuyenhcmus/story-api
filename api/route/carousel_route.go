package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/readtruyen/go-novelstory-api/api/controller"
	"github.com/readtruyen/go-novelstory-api/bootstrap"
	"github.com/readtruyen/go-novelstory-api/domain"
	"github.com/readtruyen/go-novelstory-api/mongo"
	"github.com/readtruyen/go-novelstory-api/repository"
	"github.com/readtruyen/go-novelstory-api/usecase"
)

func NewCarouselRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	cr := repository.NewCarouselRepository(db, domain.CollectionCarousel)
	cc := &controller.CarouselController{
		CarouselUseCase: usecase.NewCarouselUseCase(cr, timeout),
	}

	// Final version router
	group.GET("/carousels", cc.FetchList)
	group.POST("/carousel", cc.Create)
}
