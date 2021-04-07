package tag

import (
	"fmt"
	"github.com/gin-gonic/gin"
	app2 "go-study/other/web/framework/app"
)

type Tag struct {
}

func (t *Tag) Index(c *gin.Context) {
	app := app2.GetApp()
	app.Logger.Error("是谁的心啊，孤单的流淌")
}

func NewTag() *Tag {
	return new(Tag)
}

func Index() func(c *gin.Context) {
	fmt.Println("tag index")
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	}
}
