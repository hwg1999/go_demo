package webframework_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_data_validator(t *testing.T) {
	e := gin.Default()
	e.POST("/register", Register)
	log.Fatalln(e.Run(":8080"))
}

type LoginUser2 struct {
	Username string `binding:"required"  json:"username" form:"username" uri:"username"`
	Password string `binding:"required" json:"password" form:"password" uri:"password"`
}

func Register(ctx *gin.Context) {
	newUser := &LoginUser2{}
	if err := ctx.ShouldBind(newUser); err == nil {
		ctx.String(http.StatusOK, "user%+v", *newUser)
	} else {
		ctx.String(http.StatusBadRequest, "invalid user,%v", err)
	}
}
