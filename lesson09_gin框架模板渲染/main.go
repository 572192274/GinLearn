package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 静态文件：
// html页面上用到的样式文件 css js文件 图片
func main() {
	r := gin.Default()
	//加载静态文件，需要在解析模板之前
	//这里相当于将/xxx指向了./statics，从而在模板中使用/xxx/index.css路径
	r.Static("/xxx", "./statics")
	//gin框架中给模板添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	//Gin框架中使用LoadHTMLGlob()或者LoadHTMLFiles()方法进行HTML模板渲染。模板解析
	r.LoadHTMLFiles("templates/index.tmpl")

	//加载多个文件
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/index", func(c *gin.Context) {
		//HTTP请求
		c.HTML(http.StatusOK, "index.tmpl", gin.H{ //模板渲染
			"title": "liwenzhou.com",
		})
	})
	//通过name指定渲染的是哪个模板，name跟模板中的define定义的名字保持一致
	r.GET("/posts/index", func(c *gin.Context) {
		//HTTP请求
		c.HTML(http.StatusOK, "posts/index.html", gin.H{ //模板渲染
			"title": "posts.com",
		})
	})
	//自定义函数保证渲染的内容不被转义
	r.GET("/users/index", func(c *gin.Context) {
		//HTTP请求
		c.HTML(http.StatusOK, "users/index.html", gin.H{ //模板渲染
			"title": "<a href='http://liwenzhou.com'>博客</a>",
		})
	})
	r.Run(":9090") //启动Server
}
