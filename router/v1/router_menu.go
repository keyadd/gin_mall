package v1

import (
	"gin_mall/v1/api"
	"github.com/gin-gonic/gin"
)

func InitMenu(Router *gin.RouterGroup) {
	MenuRouter := Router.Group("/menu")
	{

		MenuRouter.POST("/list", api.MenuList)                  //后台菜单列表
		MenuRouter.POST("/update_status", api.MenuUpdateStatus) //修改状态
		MenuRouter.POST("/create", api.MenuCreate)              //创建菜单
		MenuRouter.POST("/edit", api.MenuEdit)                  //编辑菜单
		MenuRouter.POST("/delete", api.MenuDelete)              //删除菜单
	}

}
