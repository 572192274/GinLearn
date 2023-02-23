package main

import (
	"database/sql"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 定义模型
type User struct {
	ID   int64
	Name string `gorm:"default:'xiaowangzi'"` //默认值
	Age  int64
}

// 使用指针
type User2 struct {
	ID   int64
	Name *string `gorm:"default:'小王子'"`
	Age  int64
}

// 使用 Scanner/Valuer
type User3 struct {
	ID   int64
	Name sql.NullString `gorm:"default:'小王子'"` // sql.NullString 实现了Scanner/Valuer接口
	Age  int64
}

func main() {
	//1连接MySQL数据库
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//2把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})
	db.AutoMigrate(&User2{})
	db.AutoMigrate(&User3{})
	//3创建
	//所有字段的零值, 比如0, "",false或者其它零值，都不会保存到数据库内，但会使用他们的默认值。
	//如果你想避免这种情况，可以考虑使用指针或实现 Scanner/Valuer接口
	u := User{Name: "", Age: 18} //在代码层面创建一个User对象
	u2 := User2{Name: new(string), Age: 20}
	u3 := User3{Name: sql.NullString{"", true}, Age: 21}
	db.NewRecord(&u) //判断主键是否为空
	db.NewRecord(&u2)
	db.NewRecord(&u3)
	db.Debug().Create(&u)
	db.Debug().Create(&u2)
	db.Debug().Create(&u3)
}
