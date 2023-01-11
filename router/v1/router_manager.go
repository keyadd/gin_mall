package v1

import (
	"gin_mall/v1/api"
	"github.com/gin-gonic/gin"
)

func InitManager(Router *gin.RouterGroup) {
	ManagerRouter := Router.Group("/manager")
	{
		ManagerRouter.POST("/get_info", api.GetManagerInfo)        //获取用户详细信息
		ManagerRouter.POST("/update_password", api.ChangePassword) //修改密码
		ManagerRouter.POST("/logout", api.Logout)                  //退出登录
		ManagerRouter.POST("/list", api.List)                      //后台用户列表
		ManagerRouter.POST("/update_status", api.UpdateStatus)     //修改状态
		ManagerRouter.POST("/create", api.Create)                  //创建用户
		ManagerRouter.POST("/edit", api.Edit)                      //新增用户
		ManagerRouter.POST("/delete", api.Delete)                  //删除用户
	}

}
