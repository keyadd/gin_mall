package v1

import (
	"gin_mall/v1/api"
	"github.com/gin-gonic/gin"
)

func InitAppCategoryItem(Router *gin.RouterGroup) {
	AppCategoryItemRouter := Router.Group("/app_category_item") //?
	{

		AppCategoryItemRouter.POST("/list", api.AppCategoryItemList)     //后台分类列表关联分类商品
		AppCategoryItemRouter.POST("/create", api.AppCategoryItemCreate) //关联分类商品
		AppCategoryItemRouter.POST("/delete", api.AppCategoryItemDelete) //删除关联分类商品
	}

}
