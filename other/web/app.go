package main

import (
	"github.com/gin-gonic/gin"
	"go-study/other/web/framework/cache"
	"go-study/other/web/framework/logger"
	"go-study/other/web/middleware"
	"go-study/other/web/model"
	"go-study/other/web/route"
	"gorm.io/gorm"
)

type App struct {
	e      *gin.Engine
	Db     *gorm.DB
	Cache  cache.ICache
	Logger logger.ILogger
}

func NewApp() *App {
	model.Run()
	g := gin.Default()
	middleware.Use(g)
	route.New(g).Use()
	g.Run(":8082")

	app := &App{
		e:     nil,
		Db:    nil,
		Cache: cache.NewCache(),
	}

	return app
}
