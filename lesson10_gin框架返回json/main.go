package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		//方法1：使用map
		//data := map[string]interface{}{
		//	"name":    "小王子",
		//	"message": "hello world",
		//	"age":     18,
		//}
		data := gin.H{"name": "小王子", "message": "hello world", "age": 18}

		c.JSON(http.StatusOK, data)
	})
	//方法二：结构体
	//灵活使用tag来对结构体字段做定制化操作
	type msg struct {
		Name    string `json:"name"`
		Message string
		Age     int
	}
	r.GET("/another_json", func(c *gin.Context) {
		data := msg{
			"小公主",
			"Hello world",
			20,
		}
		c.JSON(http.StatusOK, data) //json序列化
	})
	r.Run(":9090")
}
