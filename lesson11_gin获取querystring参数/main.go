package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// querystring指的是URL中?后面携带的参数，例如：/user/search?username=小王子&address=沙河。
func main() {
	r := gin.Default()

	//querystring
	//GET请求URL ?后面是querystring参数
	//格式为key=value，多个key-value用 & 连接
	//eq： /web?query=小王子&age=19
	r.GET("/web", func(c *gin.Context) {
		//获取浏览器那边发请求携带的query string 参数
		//方式1
		//name := c.Query("query") //通过Query获取请求中携带的querystring参数

		//方式2
		//name := c.DefaultQuery("query", "somebody") //取不到就用指定的默认值

		//方式3
		//name, ok := c.GetQuery("query")
		//if !ok {
		//	//取不到query
		//	name = "somebody"
		//}
		//获取URL后多个参数
		name := c.Query("query")
		age := c.Query("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	r.Run(":9090")
}
