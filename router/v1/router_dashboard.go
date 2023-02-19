package v1

import (
	"gin_mall/v1/api"
	"github.com/gin-gonic/gin"
)

func InitDashboard(Router *gin.RouterGroup) {
	DashboardRouter := Router.Group("/dashboard") //?
	{

		DashboardRouter.POST("/statistics1", api.DashboardStatistics1) //后台分类列表
		DashboardRouter.POST("/statistics2", api.DashboardStatistics2) //修改状态
		DashboardRouter.POST("/statistics3", api.DashboardStatistics3) //创建分类
	}
}
