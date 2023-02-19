package v1

import (
	"gin_mall/v1/api"
	"github.com/gin-gonic/gin"
)

func InitAgent(Router *gin.RouterGroup) {
	AgentRouter := Router.Group("/agent") //?
	{
		AgentRouter.POST("/list", api.AgentList)             //公告列表
		AgentRouter.POST("/statistics", api.AgentStatistics) //创建公告
	}

}
