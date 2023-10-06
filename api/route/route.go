package route

import (
	"time"

	"github.com/readtruyen/go-novelstory-api/api/middleware"
	"github.com/readtruyen/go-novelstory-api/bootstrap"
	"github.com/readtruyen/go-novelstory-api/mongo"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)
	NewStoryRouter(env, timeout, db, publicRouter)
	NewChapterRouter(env, timeout, db, publicRouter)
	NewCarouselRouter(env, timeout, db, publicRouter)

	assignedRouter := gin.Group("")
	// Middleware to assign AccessToken
	assignedRouter.Use(middleware.JwtAssignAuthMiddleware(env.AccessTokenSecret))
	NewUserStoryRouter(env, timeout, db, assignedRouter)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewPrivateCarouselRouter(env, timeout, db, protectedRouter)
	NewPrivateStoryRouter(env, timeout, db, protectedRouter)
	NewProfileRouter(env, timeout, db, protectedRouter)
	NewTaskRouter(env, timeout, db, protectedRouter)
}
