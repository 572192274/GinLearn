package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/*
中间件是为应用提供通用服务和功能的软件。数据管理、应用服务、消息传递、身份验证和 API 管理通常都要通过中间件。
在gin框架里，就是我们的所有API接口都要经过我们的中间件，我们可以在中间件做一些拦截处理。
*/
func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "index",
	})
}

// 定义一个中间件m1，统计请求处理函数的耗时
func m1(c *gin.Context) {
	fmt.Println("m1 in...")
	//计时
	start := time.Now()
	c.Next() //调用后续的处理函数
	//c.Abort()//阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("const:%v\n", cost)
	fmt.Println("m1 out...")
}
func m2(c *gin.Context) {
	fmt.Println("m2 in...")
	c.Next() //调用后续的处理函数
	//c.Abort()//阻止调用后续的处理函数
	fmt.Println("m2 out...")
}
func main() {
	r := gin.Default()
	r.Use(m1, m2) //全局注册中间件函数m1
	r.GET("/index", m1, indexHandler)
	r.GET("/shop", m1, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "shop"})
	})
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "user",
		})
	})
	r.Run(":9090")
}
