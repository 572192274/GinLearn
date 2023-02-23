package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{
		//	"status": "ok",
		//})
		//重定向
		c.Redirect(http.StatusMovedPermanently, "http://www.sogo.com")
	})
	//内部跳转
	r.GET("/a", func(c *gin.Context) {
		//跳转到 /b对应的路由处理函数
		c.Request.URL.Path = "/b" //把请求的URL修改
		r.HandleContext(c)        //继续后续的处理
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "b",
		})
	})
	r.Run(":9090")
}
