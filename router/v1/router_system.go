package v1

import (
	"gin_mall/v1/api"
	"github.com/gin-gonic/gin"
)

func InitSystemRouter(Router *gin.RouterGroup) {
	SystemRouter := Router.Group("/system")
	{
		SystemRouter.POST("/system_info", api.GetServerInfo) // 获取服务器信息
	}
}
