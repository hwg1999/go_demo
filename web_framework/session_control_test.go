package webframework_test

import (
	"fmt"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Test_Cookie(t *testing.T) {
	router := gin.Default()
	router.GET("/cookie", func(ctx *gin.Context) {
		// 获取对应的cookie
		cookie, err := ctx.Cookie("gin_cookie")
		if err != nil {
			cookie = "NotSet"
			// 设置cookie 参数：key，val，存在时间，目录，域名，是否允许他人通过js访问cookie，仅http
			ctx.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}
		fmt.Printf("Cookie value: %s \n", cookie)
	})

	router.Run()
}

func Test_session(t *testing.T) {
	r := gin.Default()
	// 创建基于Cookie的存储引擎
	store := cookie.NewStore([]byte("secret"))
	// 设置Session中间件，mysession即session名称，也是cookie的名称
	r.Use(sessions.Sessions("mysession", store))
	r.GET("/incr", func(ctx *gin.Context) {
		// 初始化session
		session := sessions.Default(ctx)
		var count int
		// 获取值
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		// 设置
		session.Set("count", count)
		// 保存
		session.Save()
		ctx.JSON(200, gin.H{"count": count})
	})
}
