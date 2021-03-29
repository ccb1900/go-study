package app

import (
	"github.com/gin-gonic/gin"
	"go-study/other/web/framework/cache"
	"go-study/other/web/framework/logger"
	"gorm.io/gorm"
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
