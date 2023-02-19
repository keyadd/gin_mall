package api

import (
	"gin_mall/global"
	"gin_mall/utils"
	"gin_mall/v1/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetServerInfo(c *gin.Context) {
	if server, err := service.GetServerInfo(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))

		utils.ResponseSuccess(c, "获取失败")
		return
	} else {

		utils.ResponseSuccess(c, server)
	}

}
