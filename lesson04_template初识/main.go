package main

import (
	"fmt"
	"html/template"
	"net/http"
)
func sayHello(w http.ResponseWriter,r *http.Request){
	//解析模板
	t,err:=template.ParseFiles("./hello.tmpl")
	if err!=nil{
		fmt.Println("Parse template failed, err:%v",err)
		return
	}
	//渲染模板
	//替换模板中的内容
	err=t.Execute(w,"小王子")
	if err!=nil{
		fmt.Println("render template failed, err:%v",err)
		return
	}
}
func main(){
	http.HandleFunc("/",sayHello)
	err:=http.ListenAndServe(":9000",nil)
	if err!=nil{
		fmt.Println("HTTP server start failed, err:%v",err)
		return
	}
}
