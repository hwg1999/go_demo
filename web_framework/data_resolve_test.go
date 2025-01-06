package webframework_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Test_data_resolve(t *testing.T) {
	e := gin.Default()
	e.POST("/loginWithJSON", Login)
	e.POST("/loginWithForm", Login)
	e.GET("/loginWithQuery/:username/:password", Login)
	e.POST("/someHandler", SomeHandler)
	e.Run(":8080")
}

type LoginUser struct {
	Username string `bind:"required" json:"username" form:"username" uri:"username"`
	Password string `bind:"required" json:"password" form:"password" uri:"password"`
}

func Login(c *gin.Context) {
	var login LoginUser
	// 使用ShouldBind来让gin自动推断
	if c.ShouldBind(&login) == nil && login.Password != "" && login.Username != "" {
		c.String(http.StatusOK, "login successfully !")
	} else if c.ShouldBindUri(&login) == nil && login.Password != "" && login.Username != "" {
		c.String(http.StatusOK, "login successfully !")
	} else {
		c.String(http.StatusBadRequest, "login failed !")
	}
	fmt.Println(login)
}

func SomeHandler(c *gin.Context) {
	objA := struct{}{}
	objB := struct{}{}
	// 读取 c.Request.Body 并将结果存入上下文。
	if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
		// 这时, 复用存储在上下文中的 body。
	}
	if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
		c.String(http.StatusOK, `the body should be formB JSON`)
		// 可以接受其他格式
	}
	if errB2 := c.ShouldBindBodyWith(&objB, binding.XML); errB2 == nil {
		c.String(http.StatusOK, `the body should be formB XML`)
	}
}
