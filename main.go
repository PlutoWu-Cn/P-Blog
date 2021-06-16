package main

import (
	"p-blog/model"
	"p-blog/routes"
)

func main() {
	//引入数据库
	model.InitDb()
	// 引入路由组件
	routes.InitRouter()
}
