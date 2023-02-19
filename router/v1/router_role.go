package v1

import (
	"gin_mall/v1/api"
	"github.com/gin-gonic/gin"
)

func InitRole(Router *gin.RouterGroup) {
	RoleRouter := Router.Group("/role") //?
	{
		RoleRouter.POST("/list", api.RoleList)                  //后台角色列表
		RoleRouter.POST("/update_status", api.RoleUpdateStatus) //修改状态
		RoleRouter.POST("/create", api.RoleCreate)              //创建角色
		RoleRouter.POST("/edit", api.RoleEdit)                  //新增角色
		RoleRouter.POST("/delete", api.RoleDelete)              //删除角色
		RoleRouter.POST("/set_rules", api.RoleSetRules)         //设置角色
	}

}
