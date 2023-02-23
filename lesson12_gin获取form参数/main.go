package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login", func(c *gin.Context) {
		//获取form表单提交的数据
		//方法1
		//username := c.PostForm("username")
		//password := c.PostForm("password")//取到就返回值，取不到返回空字符串

		//方法2
		//username := c.DefaultPostForm("username", "somebody")
		//password := c.DefaultPostForm("password", "***")

		//方法3
		username, ok := c.GetPostForm("username")
		if !ok {
			username = "sb"
		}
		password, _ := c.GetPostForm("password")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})
	})
	r.Run(":9090")
}
