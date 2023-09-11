package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"readtruyen-api/api/controller"
	"readtruyen-api/bootstrap"
	"readtruyen-api/domain"
	"readtruyen-api/mongo"
	"readtruyen-api/repository"
	"readtruyen-api/usecase"
)

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}
