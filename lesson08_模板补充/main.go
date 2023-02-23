package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// 修改模板引擎的标识符
func index(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	//在解析之前修改模板引擎的标识符
	t, err := template.New("index.tmpl").Delims("{[", "]}").ParseFiles("index.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	//渲染模板
	t.Execute(w, "小王子")
}

// html/template针对的是需要返回HTML内容的场景，在模板渲染过程中会对一些有风险的内容进行转义，以此来防范跨站脚本攻击。
func xss(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	//解析模板之前定义一个自定义的函数safe，为了区分恶意脚本和善意的
	t, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	}).ParseFiles("./xss.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	//渲染模板
	str1 := "<script>alert(123);</script>"
	str2 := "<a href='http://liwenzhou.com'>liwenzhou博客</a>"

	t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})
}
func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP server start failed, err:%v", err)
		return
	}
}
