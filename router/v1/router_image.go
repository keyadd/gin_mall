package v1

import (
	"gin_mall/v1/api"
	"github.com/gin-gonic/gin"
)

func InitImage(Router *gin.RouterGroup) {
	ImageRouter := Router.Group("/image") //？
	{

		ImageRouter.POST("/upload", api.ImageUpload) //上传图片
		ImageRouter.POST("/rename", api.ImageRename) //修改图片名
		ImageRouter.POST("/delete", api.ImageDelete) //删除图片
	}

}
