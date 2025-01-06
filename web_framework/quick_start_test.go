package webframework_test

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_quick_start(t *testing.T) {
	engine := gin.Default()
	engine.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	engine.Run()
}
