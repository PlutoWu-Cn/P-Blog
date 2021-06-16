package routes

import (
	"github.com/gin-gonic/gin"
	"p-blog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	/*
		后台管理路由接口
	*/
	auth := r.Group("api/v1")
	{
		// 用户模块的路由接口

		//修改密码

		// 分类模块的路由接口

		// 文章模块的路由接口

		// 上传文件

		// 更新个人设置

		// 评论模块

	}

	r.Run(utils.HttpPort)

}