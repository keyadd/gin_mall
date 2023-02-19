package v1

import (
	"gin_mall/v1/api"
	"github.com/gin-gonic/gin"
)

func InitImageClass(Router *gin.RouterGroup) {
	ImageClassRouter := Router.Group("/image_class") //？
	{
		ImageClassRouter.POST("/image_list", api.ImageList)    //指定分类id图片列表？
		ImageClassRouter.POST("/list", api.ImageClassList)     //图库分类列表
		ImageClassRouter.POST("/create", api.ImageClassCreate) //创建图库分类
		ImageClassRouter.POST("/edit", api.ImageClassEdit)     //新增图库分类
		ImageClassRouter.POST("/delete", api.ImageClassDelete) //删除图库分类
	}

}
