package webframework_test

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func Test_global_middleware(t *testing.T) {
	e := gin.Default()
	// 注册全局中间件
	e.Use(GlobalMiddleware())
	v1 := e.Group("/v1")
	{
		v1.GET("/hello", Hello2)
		v1.GET("/login", Login2)
	}
	v2 := e.Group("/v2")
	{
		v2.POST("/update", Update2)
		v2.DELETE("/delete", Delete2)
	}
	log.Fatalln(e.Run(":8080"))
}

func GlobalMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("全局中间件被执行")
	}
}

func Hello2(ctx *gin.Context) {
	fmt.Println("hello2")
}

func Login2(ctx *gin.Context) {
	fmt.Println("login2")
}
func Update2(ctx *gin.Context) {
	fmt.Println("update2")
}
func Delete2(ctx *gin.Context) {
	fmt.Println("delete2")
}

func Test_local_middleware(t *testing.T) {
	e := gin.Default()
	// 注册全局中间件
	e.Use(GlobalMiddleware())
	// 注册路由组局部中间件
	v1 := e.Group("/v1", LocalMiddleware())
	{
		v1.GET("/hello", Hello2)
		v1.GET("/login", Login2)
	}
	v2 := e.Group("/v2")
	{
		// 注册单个路由局部中间件
		v2.POST("/update", LocalMiddleware(), Update2)
		v2.DELETE("/delete", Delete2)
	}
	log.Fatalln(e.Run(":8080"))
}

func LocalMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("局部中间件被执行")
	}
}

func Test_requrest_time_middleware(t *testing.T) {
	e := gin.Default()
	// 注册全局中间件，计时中间件
	e.Use(GlobalMiddleware(), TimeMiddleware())
	// 注册路由组局部中间件
	v1 := e.Group("/v1", LocalMiddleware())
	{
		v1.GET("/hello", Hello2)
		v1.GET("/login", Login2)
	}
	v2 := e.Group("/v2")
	{
		// 注册单个路由局部中间件
		v2.POST("/update", LocalMiddleware(), Update2)
		v2.DELETE("/delete", Delete2)
	}
	log.Fatalln(e.Run(":8080"))
}

func TimeMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 记录开始时间
		start := time.Now()
		// 执行后续调用链
		ctx.Next()
		// 计算时间间隔
		duration := time.Since(start)
		// 输出纳秒，以便观测结果
		fmt.Println("请求用时: ", duration.Nanoseconds())
	}
}
