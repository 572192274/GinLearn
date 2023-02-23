package main

import (
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
	//删除
	//警告 删除记录时，请确保主键字段有值，GORM 会通过主键去删除记录，如果主键为空，GORM 会删除该 model 的所有记录。
	var u User
	u.ID = 1
	db.Debug().Delete(&u)
	//删除全部匹配的记录
	db.Debug().Where("name=?", "hello").Delete(User{})
	db.Delete(User{}, "age=?", 24)

	// Unscoped 方法可以物理删除记录
	db.Unscoped().Delete(&u)

	//如果一个 model 有 DeletedAt 字段，他将自动获得软删除的功能！ 当调用 Delete 方法时， 记录不会真正的从数据库中被删除， 只会将DeletedAt 字段的值会被设置为当前时间
	//如果不需要软删除，结构体不需要嵌套gorm.Model
}
