package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"ppp/other/web/framework/cache"
	"ppp/other/web/framework/logger"
	"sync"
)

type App struct {
	e      *gin.Engine
	Db     *gorm.DB
	Cache  cache.ICache
	Logger logger.ILogger
}

var once sync.Once
var app *App

func GetApp() *App {
	once.Do(func() {
		app = &App{}
	})

	return app
}
