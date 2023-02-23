package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// 未使用模板继承
func index(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:%v", err)
		return
	}
	//渲染模板
	msg := "这是index页面"
	t.Execute(w, msg)

}
func home(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./home.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:%v", err)
		return
	}
	//渲染模板
	msg := "这是home页面"
	t.Execute(w, msg)

}

// 使用模板继承
func index2(w http.ResponseWriter, r *http.Request) {
	//定义模板(模板继承的方式)
	//解析模板
	t, err := template.ParseFiles("./template/base.tmpl", "./template/index.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:%v", err)
		return
	}
	//渲染模板
	msg := "这是index2页面"
	t.ExecuteTemplate(w, "index.tmpl", msg)

}
func home2(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./template/base.tmpl", "./template/home.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:%v", err)
		return
	}
	//渲染模板
	msg := "这是home2页面"
	t.ExecuteTemplate(w, "home.tmpl", msg)

}

// 使用另一种方式完成模板继承
func base(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("template/*.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	err = tmpl.ExecuteTemplate(w, "index.tmpl", "index2.tmpl")
	if err != nil {
		fmt.Println("render template failed, err:", err)
		return
	}
}
func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)
	http.HandleFunc("/base", base)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP server start failed, err:%v", err)
		return
	}
}
