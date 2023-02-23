package main

import (
	"fmt"
	"html/template"
	"net/http"
)
//自定义函数演示
func f1(w http.ResponseWriter,r *http.Request){
	//定义一个函数
	//要么只有一个返回值，要么有两个返回值，第二个返回值必须是error类型
	kua:= func(name string)(string,error) {
		return name+"年轻又帅气！",nil
	}
	//定义模板
	//创建一个名字是f的模板对象，名字一定要与模板的名字能对应上
	t:=template.New("f.tpl")
	//告诉模板引擎，现在多了一个自定义的函数kua
	t.Funcs(template.FuncMap{
		"kua":kua,
	})
	//解析模板

	_,err:=t.ParseFiles("./f.tpl")
	if err!=nil{
		fmt.Println("Parse template failed, err:%v",err)
		return
	}

	//渲染模板
	name:="小王子"
	t.Execute(w,name)

}
//嵌套模板演示
func demo1(w http.ResponseWriter,r *http.Request){
	//定义模板
	//解析模板
	t,err:=template.ParseFiles("./t.tmpl","./ul.tmpl")
	if err!=nil{
		fmt.Println("Parse template failed, err:%v",err)
		return
	}

	//渲染模板
	name:="小公主"
	t.Execute(w,name)
}
func main(){
	http.HandleFunc("/",f1)
	http.HandleFunc("/tmplDemo",demo1)
	err:=http.ListenAndServe(":9090",nil)
	if err!=nil{
		fmt.Println("HTTP server start failed, err:%v",err)
		return
	}
}
