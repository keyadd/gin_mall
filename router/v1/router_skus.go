package v1

import (
	"gin_mall/v1/api"
	"github.com/gin-gonic/gin"
)

func InitSkus(Router *gin.RouterGroup) {
	SkusRouter := Router.Group("/skus") //？
	{
		SkusRouter.POST("/list", api.SkusList)                  //后台用户列表
		SkusRouter.POST("/update_status", api.SkusUpdateStatus) //修改状态
		SkusRouter.POST("/create", api.SkusCreate)              //创建用户
		SkusRouter.POST("/edit", api.SkusEdit)                  //新增用户
		SkusRouter.POST("/delete_all", api.SkusDelete)          //删除用户
	}

}
