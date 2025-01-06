package webframework_test

import (
	"fmt"
	"io"
	"log"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func Test_log_file(t *testing.T) {
	e := gin.Default()
	// 关掉控制台颜色
	gin.DisableConsoleColor()
	// 创建两个日志文件
	log1, _ := os.Create("info1.log")
	log2, _ := os.Create("info2.log")
	// 同时记录进两个日志文件
	gin.DefaultWriter = io.MultiWriter(log1, log2)
	e.GET("/hello", Hello)
	log.Fatalln(e.Run(":8080"))
}

func Test_log_file2(t *testing.T) {
	router := gin.New()
	// LoggerWithFormatter 中间件会写入日志到 gin.DefaultWriter
	// 默认 gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// 输出自定义格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.Run(":8080")
}

func Test_router_log(t *testing.T) {
	e := gin.Default()
	gin.SetMode(gin.DebugMode)
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		if gin.IsDebugging() {
			log.Printf("路由 %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
		}
	}
	e.GET("/hello", Hello)
	log.Fatalln(e.Run(":8080"))
}
