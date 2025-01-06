package webframework_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_router_parameter(t *testing.T) {
	e := gin.Default()
	e.GET("/findUser/:username/:userid", FindUser)
	e.GET("/downloadFile/*filepath", UserPage)

	log.Fatalln(e.Run(":8080"))
}

// 命名参数示例
func FindUser(c *gin.Context) {
	username := c.Param("username")
	userid := c.Param("userid")
	c.String(http.StatusOK, "username is %s\n userid is %s", username, userid)
}

// 路径参数示例
func UserPage(c *gin.Context) {
	filepath := c.Param("filepath")
	c.String(http.StatusOK, "filepath is %s", filepath)
}

func Test_url_parameter(t *testing.T) {
	e := gin.Default()
	e.GET("/findUser2", FindUser2)
	log.Fatalln(e.Run(":8080"))
}

func FindUser2(c *gin.Context) {
	username := c.DefaultQuery("username", "defaultUser")
	userid := c.Query("userid")
	c.String(http.StatusOK, "username is %s\nuserid is %s", username, userid)
}

func Test_form_parameter(t *testing.T) {
	e := gin.Default()
	e.POST("/register", RegisterUser)
	e.POST("/update", UpdateUser)
	e.Run(":8080")
}

func RegisterUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	c.String(http.StatusOK, "successfully registered,your username is [%s],password is [%s]", username, password)
}

func UpdateUser(c *gin.Context) {
	var form map[string]string
	c.ShouldBind(&form)
	c.String(http.StatusOK, "successfully update,your username is [%s],password is [%s]", form["username"], form["password"])
}
