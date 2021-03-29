package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-study/other/web/controller"
	"go-study/other/web/framework/cache"
	"go-study/other/web/framework/logger"
	"go-study/other/web/middleware"
	"gorm.io/gorm"
)

type App struct {
	e      *gin.Engine
	Db     *gorm.DB
	Cache  cache.ICache
	Logger logger.ILogger
}

func main() {
	g := gin.Default()
	middlewareGroup(g)
	route(g)
	g.Run(":8082")

	app := App{
		e:     nil,
		Db:    nil,
		Cache: cache.NewCache(),
	}

	fmt.Println(app)
}

func middlewareGroup(g *gin.Engine) {
	g.Use(middleware.Logger())
	g.Use(middleware.Demo())
}

func route(g *gin.Engine) {
	g.GET("/", test)
	g.GET("tags", controller.NewTag().Index)
	g.GET("tags/:id", showTag)
	g.DELETE("tags/:id", test)
	g.PUT("tags/:id", test)
	g.POST("tags", test)
	g.POST("upload", upload)
}

func test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

func showTag(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
		"id":      c.Param("id"),
		"qid":     c.Query("id"),
	})
}

func upload(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "upload",
	})
}
