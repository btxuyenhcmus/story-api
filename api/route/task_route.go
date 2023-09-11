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

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
}
