package v1

import (
	"gin_mall/v1/api"
	"github.com/gin-gonic/gin"
)

func InitApi(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("/api")
	{
		ApiRouter.POST("/list", api.ApiList)                  //后台菜单列表
		ApiRouter.POST("/update_status", api.ApiUpdateStatus) //修改状态
		ApiRouter.POST("/create", api.ApiCreate)              //创建菜单
		ApiRouter.POST("/edit", api.ApiEdit)                  //新增菜单
		ApiRouter.POST("/delete", api.ApiDelete)              //删除菜单
		ApiRouter.POST("/list_tree", api.ApiListTree)
	}

}
