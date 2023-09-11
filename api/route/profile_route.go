package route

import (
	"time"

	"readtruyen-api/api/controller"
	"readtruyen-api/bootstrap"
	"readtruyen-api/domain"
	"readtruyen-api/mongo"
	"readtruyen-api/repository"
	"readtruyen-api/usecase"

	"github.com/gin-gonic/gin"
)

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
	group.PUT("/profile", pc.Update)
}
