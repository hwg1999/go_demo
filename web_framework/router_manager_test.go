package webframework_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_router_group(t *testing.T) {
	e := gin.Default()
	v1 := e.Group("v1")
	{
		v1.GET("/hello", func(ctx *gin.Context) {})
		v1.GET("/login", func(ctx *gin.Context) {})
	}
	v2 := e.Group("v2")
	{
		v2.POST("/update", func(ctx *gin.Context) {})
		v2.DELETE("/delete", func(ctx *gin.Context) {})
	}
}

func Test_404_router(t *testing.T) {
	e := gin.Default()
	e.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "<h1>404 Page Not Found</h1>")
	})
	log.Fatalln(e.Run(":8080"))
}

func Test_405_notallowed(t *testing.T) {
	e := gin.Default()
	e.HandleMethodNotAllowed = true
	e.NoRoute(func(context *gin.Context) {
		context.String(http.StatusNotFound, "<h1>404 Page Not Found</h1>")
	})
	// 注册处理器
	e.NoMethod(func(context *gin.Context) {
		context.String(http.StatusMethodNotAllowed, "method not allowed")
	})
	log.Fatalln(e.Run(":8080"))
}

func Test_redirect(t *testing.T) {
	e := gin.Default()
	e.GET("/", IndexPermanently)
	e.GET("/hello", Hello)
	log.Fatalln(e.Run(":8080"))
}

func IndexPermanently(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/hello")
}

func Hello(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}
