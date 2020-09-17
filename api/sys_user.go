package api

import (
	"go_yangzhou/model"

	"github.com/gin-gonic/gin"
)

/**
用户api
*/

// User登录
func UserLogin(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err == nil {
	}
}

// User退出
func UserLogout(c *gin.Context) {

}

// User检测是否登录
func NewUser(c *gin.Context) {
	model.BashWithMap()
}

func GetRecords(c *gin.Context)  {
	model.GetRecords()
}
