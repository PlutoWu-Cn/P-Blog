package main

import (
	"p-blog/model"
	"p-blog/routes"
)

func main() {

	//引入数据库
	model.InitDb()


	routes.InitRouter()


}
