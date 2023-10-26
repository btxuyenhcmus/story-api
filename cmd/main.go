package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	route "github.com/readtruyen/go-novelstory-api/api/route"
	"github.com/readtruyen/go-novelstory-api/bootstrap"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()
	gin.Use(cors.Default())

	route.Setup(env, timeout, db, gin)

	gin.Run(env.ServerAddress)
}
