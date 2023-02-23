package main

import (
	"fmt"
	"html/template"
	"net/http"
)
type User struct {
	Name string
	Gender string
	Age int
}
func sayHello(w http.ResponseWriter,r *http.Request){
	//定义模板
	//解析模板
	t,err:=template.ParseFiles("./hello.tmpl")
	if err!=nil{
		fmt.Println("Parse template failed, err:%v",err)
		return
	}
	//渲染模板
	//模板中的.相当于传入的对象
	//模板语法都包含在{{和}}中间，其中{{.}}中的点表示当前对象。
	//当我们传入一个结构体对象时，我们可以根据.来访问结构体的对应字段
	u1:=User{
		Name: "小王子",
		Gender: "男",
		Age:18,
	}
	//字典形式可以在模板中通过键来取值
	//我们传入的变量是map时，也可以在模板文件中通过.根据key来取值。
	m1:=map[string]interface{}{
		"name":"小公主",
		"gender":"女",
		"age":20,
	}
	hobbyList:=[]string{
		"篮球",
		"足球",
		"双色球",
	}
	t.Execute(w,map[string]interface{}{
		"u1":u1,
		"m1":m1,
		"hobby":hobbyList,
	})
}
func main(){
	http.HandleFunc("/",sayHello)
	err:=http.ListenAndServe(":9090",nil)
	if err!=nil{
		fmt.Println("HTTP server start failed, err:%v",err)
		return
	}
}
