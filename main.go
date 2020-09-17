package main

import (
	"go_yangzhou/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("resource/template/*.html")
	//默认请求,返回html
	r.GET("", func(c *gin.Context) {

		c.HTML(http.StatusOK, "welcome.html", gin.H{
			"host":  c.ClientIP(),
			"local": c.DefaultQuery("local", "en"),
		})
	})
	//检测请求，返回json
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//合成登录
	r.GET("/video/compose", api.VideoCompose)
	//用户接口
	user := r.Group("/user")
	{
		user.POST("/login", api.UserLogin)
		user.POST("/logout", api.UserLogout)
		user.POST("/newUser", api.NewUser)
		user.GET("/records",api.GetRecords)
	}

	r.Run(":9090") // listen and serve on 0.0.0.0:8080
}
