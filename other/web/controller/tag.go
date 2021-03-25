package controller

import (
	"github.com/gin-gonic/gin"
	app2 "ppp/other/web/framework/app"
)

type Tag struct {
}

func (t Tag) Index(c *gin.Context) {
	app := app2.GetApp()
	app.Logger.Error("是谁的心啊，孤单的流淌")
}

func NewTag() *Tag {
	return new(Tag)
}
