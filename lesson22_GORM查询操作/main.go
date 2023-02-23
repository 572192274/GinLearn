package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 定义模型
type User struct {
	ID   int64
	Name string
	Age  int64
}

func main() {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&User{})
	//查询
	var user User //声明模型结构体类型变量user
	//根据主键查询第一条记录
	db.First(&user)
	fmt.Printf("user:%#v\n", user)
	// 随机获取一条记录
	db.Take(&user)
	fmt.Printf("user:%#v\n", user)
	// 根据主键查询最后一条记录
	db.Last(&user)
	//// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	fmt.Printf("user:%#v\n", user)

	// 查询所有的记录
	var users []User
	db.Find(&users)
	fmt.Printf("users:%#v\n", users)

	// 查询指定的某条记录(仅当主键为整型时可用)
	db.First(&user, 10)
	//// SELECT * FROM users WHERE id = 10;
	fmt.Printf("user:%#v\n", user)

}
