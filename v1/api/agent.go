package api

import (
	"gin_mall/global"
	"gin_mall/utils"
	"gin_mall/v1/model/request"
	"gin_mall/v1/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AgentList(c *gin.Context) {
	var r request.AgentList
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("ManagerList with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	list := service.GetAgentList(r)
	utils.ResponseSuccess(c, list)
}

func AgentStatistics(c *gin.Context) {
	result := service.AgentStatistics()
	utils.ResponseSuccess(c, result)

}
