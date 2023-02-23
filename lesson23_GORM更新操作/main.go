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
	//更新
	var user User
	db.First(&user)
	user.Name = "七米"
	//更新所有字段
	db.Save(&user)
	//如果你只希望更新指定字段，可以使用Update或者Updates
	db.Model(&user).Update("name", "小王子")
	// 使用 map 更新多个属性，只会更新其中有变化的属性
	db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18})
	// 使用 struct 更新多个属性，只会更新其中有变化且为非零值的字段
	db.Model(&user).Updates(User{Name: "hello", Age: 18})
	//如果你想更新或忽略某些字段，你可以使用 Select，Omit
	db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 32}) //只更新age
	db.Model(&user).Omit("name").Updates(map[string]interface{}{"name": "hello", "age": 18})   //忽略name的更新
	//无Hooks更新
	//上面的更新操作会自动运行 model 的 BeforeUpdate, AfterUpdate 方法，更新 UpdatedAt 时间戳, 在更新时保存其 Associations, 如果你不想调用这些方法，你可以使用 UpdateColumn， UpdateColumns
	// 更新单个属性，类似于 `Update`
	db.Model(&user).UpdateColumn("name", "hello")
	//// UPDATE users SET name='hello' WHERE id = 111;
	// 更新多个属性，类似于 `Updates`
	db.Model(&user).UpdateColumns(User{Name: "hello", Age: 18})
	//// UPDATE users SET name='hello', age=18 WHERE id = 111;

	//批量更新时Hooks（钩子函数）不会运行。
	// 使用 `RowsAffected` 获取更新记录总数
	row := db.Model(User{}).Updates(User{Name: "hello", Age: 18}).RowsAffected
	fmt.Println(row)
	//让users表中所有用户的年龄在原来的基础上+2
	db.Model(&User{}).Update("age", gorm.Expr("age+?", 2))
}
