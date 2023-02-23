package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 单个文件上传
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/upload", func(c *gin.Context) {
		//从请求中读取文件
		f, err := c.FormFile("f1") //从请求中获取携带的参数一样
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			//将读取到的文件保持在本地（服务端本地）
			dst := fmt.Sprintf("./%s", f.Filename)
			c.SaveUploadedFile(f, dst) //上次到指定的文件夹
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		}
	})
	r.Run(":9090")
}
