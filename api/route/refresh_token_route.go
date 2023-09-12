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

func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, timeout),
		Env:                 env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}
