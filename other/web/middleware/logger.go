package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", "12354")
		log.Print("一个请求抵达了。。。", c.Request.URL.RawQuery)
		c.Next()
		latency := time.Since(t)
		log.Print("啊，请求来了。。", latency)
		status := c.Writer.Status()
		log.Println("啊，请求状态", status)
	}
}

func Demo() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", "12354")
		log.Print("demo。。。", c.Request.URL.RawQuery)
		c.Next()
		latency := time.Since(t)
		log.Print("后置", latency)
		status := c.Writer.Status()
		log.Println("后置o", status)
	}
}
