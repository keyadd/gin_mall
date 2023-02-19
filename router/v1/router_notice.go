package v1

import (
	"gin_mall/v1/api"
	"github.com/gin-gonic/gin"
)

func InitNotice(Router *gin.RouterGroup) {
	NoticeRouter := Router.Group("/notice") //?
	{
		NoticeRouter.POST("/list", api.NoticeList)     //公告列表
		NoticeRouter.POST("/create", api.NoticeCreate) //创建公告
		NoticeRouter.POST("/edit", api.NoticeEdit)     //新增公告
		NoticeRouter.POST("/delete", api.NoticeDelete) //删除公告
	}

}
