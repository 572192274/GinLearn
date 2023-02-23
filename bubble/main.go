package main

import (
	"gin_learn/bubble/dao"
	"gin_learn/bubble/models"
	"gin_learn/bubble/routers"
)

func main() {
	//创建数据库
	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})
	r := routers.SetupRouter()
	r.Run(":9090")
}
