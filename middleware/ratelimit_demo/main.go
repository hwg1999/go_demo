package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	ratelimit2 "github.com/juju/ratelimit" // 令牌桶
	ratelimit1 "go.uber.org/ratelimit"     // 漏桶
)

func main() {
	r := gin.Default()
	r.GET("/ping", rateLimit1(), func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})
	r.GET("/hei", rateLimit2(2*time.Second, 1), func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ha")
	})
}

// 基于漏桶的限流中间件
func rateLimit1() func(ctx *gin.Context) {
	// 生成一个限流器
	rl := ratelimit1.New(100)
	return func(ctx *gin.Context) {
		// 取水滴
		if time.Until(rl.Take()) > 0 {
			//time.Sleep(rl.Take().Sub(time.Now())) // 需要等这么长时间下一滴水才会滴下来
			ctx.String(http.StatusOK, "rate limit...")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

// 基于令牌桶的限流中间件2
func rateLimit2(fillInterval time.Duration, cap int64) func(ctx *gin.Context) {
	rl := ratelimit2.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		// rl.Take()          // 这一次可以欠账
		// rl.TakeAvailable // 有令牌的时候才会取出令牌
		if rl.TakeAvailable(1) == 1 { // 此次取到令牌
			c.Next()
			return
		}
		c.String(http.StatusOK, "rate limit....")
		c.Abort()
	}
}
