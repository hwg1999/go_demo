package webframework_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_data_response(t *testing.T) {
	e := gin.Default()
	// 加载HTML文件，也可以使用Engine.LoadHTMLGlob()
	e.LoadHTMLFiles("index.html")
	e.GET("/", Index)
	log.Fatalln(e.Run(":8080"))
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
