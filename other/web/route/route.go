package route

import (
	"github.com/gin-gonic/gin"
	"go-study/other/web/controller/post"
	"go-study/other/web/controller/tag"
)

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

type Route struct {
	g *gin.Engine
}

func New(g *gin.Engine) *Route {
	return &Route{g: g}
}
func (r *Route) Use() {
	r.g.GET("/", test)
	r.g.GET("tags", tag.Index())
	r.g.GET("tags/:id", showTag)
	r.g.DELETE("tags/:id", test)
	r.g.PUT("tags/:id", test)
	r.g.POST("tags", test)
	r.g.POST("upload", upload)
	r.postCate()
}

func (r *Route) postCate() {
	r.g.GET("post", post.New().Create)
}
