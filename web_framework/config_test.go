package webframework_test

import (
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func Test_http_config(t *testing.T) {
	router := gin.Default()
	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())
}

func Test_static_resources(t *testing.T) {
	router := gin.Default()
	// 加载静态文件目录
	router.Static("/static", "./static")
	// 加载静态文件目录
	router.StaticFS("/view", http.Dir("view"))
	// 加载静态文件
	router.StaticFile("/favicon", "./static/favicon.ico")

	router.Run(":8080")
}

func CorsMiddle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		origin := ctx.Request.Header.Get("origin")
		if origin != "" {
			// 生产环境中的服务端通常都不会填 *，应当填写指定域名
			ctx.Header("Access-Control-Allow-Origin", origin)
			// 允许使用的HTTP METHOD
			ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			// 允许使用的请求头
			ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			// 允许客户端访问的响应头
			ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			// 是否需要携带认证信息 Credentials 可以是 cookies、authorization headers 或 TLS client certificates
			// 设置为true时，Access-Control-Allow-Origin不能为 *
			ctx.Header("Access-Control-Allow-Credentials", "true")
		}

		// 放行OPTION请求，但不执行后续方法
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		// 放行
		ctx.Next()
	}
}
