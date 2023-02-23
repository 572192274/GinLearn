package main

import (
	"fmt"
	"net/http"
	"os"
)
func sayHello(w http.ResponseWriter, r *http.Request){
	//_,_ = fmt.Fprintln(w,"<h1>Hello Golang!</h1><h2>how are you!</h2>")
	//hello.txt相当于html，浏览器会自动解析其中的内容
	b,_:=os.ReadFile("./hello.txt")
	_,_ = fmt.Fprintln(w,string(b))
}
func main(){
	//HandleFunc注册一个处理器函数handler和对应的模式pattern（注册到DefaultServeMux）。
	//ServeMux的文档解释了模式的匹配机制。
	http.HandleFunc("/hello",sayHello)

	//ListenAndServe监听TCP地址addr，并且会使用handler参数调用Serve函数处理接收到的连接。
	//handler参数一般会设为nil，此时会使用DefaultServeMux。
	err:=http.ListenAndServe(":9090",nil)
	if err!=nil{
		fmt.Printf("http serve failed, err:%v\n",err)
		return
	}
}

