package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo ---> 数据库表
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//创建表 自动迁移 （把结构体和数据表进行对应）
	db.AutoMigrate(&UserInfo{})
	//创建数据行
	//u1 := UserInfo{1, "七米", "男", "蛙泳"}
	//db.Create(u1)

	//查询
	var u UserInfo
	db.First(&u) //查询表中第一条数据保持到u中
	fmt.Printf("u:%#v\n", u)
	//更新
	db.Model(&u).Update("hobby", "双色球")
	//删除
	db.Delete(&u)
}
