package v1

import (
	"gin_mall/v1/api"
	"github.com/gin-gonic/gin"
)

func InitCategory(Router *gin.RouterGroup) {
	CategoryRouter := Router.Group("/category") //?
	{

		CategoryRouter.POST("/list", api.CategoryList)                  //后台分类列表
		CategoryRouter.POST("/update_status", api.CategoryUpdateStatus) //修改状态
		CategoryRouter.POST("/create", api.CategoryCreate)              //创建分类
		CategoryRouter.POST("/edit", api.CategoryEdit)                  //新增分类
		CategoryRouter.POST("/delete", api.CategoryDelete)              //删除分类
	}

}
