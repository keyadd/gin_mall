package api

import (
	"gin_mall/global"
	"gin_mall/utils"
	"gin_mall/v1/model/request"
	"gin_mall/v1/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// List 获取管理员列表
func DashboardStatistics1(c *gin.Context) {
	res := service.GetStatistics1()

	utils.ResponseSuccess(c, res)

	return

}

func DashboardStatistics2(c *gin.Context) {

	res := service.GetStatistics2()
	utils.ResponseSuccess(c, res)

}

// Create 赠加管理员
func DashboardStatistics3(c *gin.Context) {
	var r request.Statistics3
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("ManagerList with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.GetStatistics3(r)
	utils.ResponseSuccess(c, res)

}
