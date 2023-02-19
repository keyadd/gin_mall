package v1

import (
	"gin_mall/v1/api"
	"github.com/gin-gonic/gin"
)

func InitUser(Router *gin.RouterGroup) {
	UserRouter := Router.Group("/user") //？
	{

		UserRouter.POST("/list", api.UserList)                  //后台分类列表
		UserRouter.POST("/update_status", api.UserUpdateStatus) //修改状态
		UserRouter.POST("/create", api.UserCreate)              //创建分类
		UserRouter.POST("/edit", api.UserEdit)                  //新增分类
		UserRouter.POST("/delete", api.UserDelete)              //删除分类
	}
}
