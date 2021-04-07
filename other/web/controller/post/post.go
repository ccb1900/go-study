package post

import (
	"github.com/gin-gonic/gin"
	"go-study/other/web/controller"
)

type Post struct {
	controller.Controller
}

func New() *Post {
	return &Post{}
}
func (p *Post) Create(c *gin.Context) {
	p.Println("test")
	p.Test()
	c.JSON(200, gin.H{
		"msg":  "ok",
		"code": 200,
		"data": nil,
	})
}
