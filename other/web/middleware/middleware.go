package middleware

import "github.com/gin-gonic/gin"

func Use(g *gin.Engine) {
	g.Use(Logger())
	g.Use(Demo())
}
