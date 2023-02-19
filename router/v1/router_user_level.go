package v1

import (
	"gin_mall/v1/api"
	"github.com/gin-gonic/gin"
)

func InitUserLevel(Router *gin.RouterGroup) {
	UserLevelRouter := Router.Group("/user_level") //？
	{

		UserLevelRouter.POST("/list", api.UserLevelList)                  //后台分类列表
		UserLevelRouter.POST("/update_status", api.UserLevelUpdateStatus) //修改状态
		UserLevelRouter.POST("/create", api.UserLevelCreate)              //创建分类
		UserLevelRouter.POST("/edit", api.UserLevelEdit)                  //新增分类
		UserLevelRouter.POST("/delete", api.UserLevelDelete)              //删除分类
	}
}
