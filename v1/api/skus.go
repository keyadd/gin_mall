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
func SkusList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBind(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error("ManagerList with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.GetSkusList(pageInfo)

	utils.ResponseSuccess(c, res)

	return

}

func SkusUpdateStatus(c *gin.Context) {
	var r request.UpdateStatusRequest
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}

	res := service.UpdateStatusSkus(r)
	utils.ResponseSuccess(c, res)

}

// Create 赠加管理员
func SkusCreate(c *gin.Context) {
	var r request.AddSkus
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.CreateSkus(r)
	utils.ResponseSuccess(c, res)

}

func SkusEdit(c *gin.Context) {
	var r request.EditSkus
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}

	res := service.EditSkus(r)
	utils.ResponseSuccess(c, res)
}

func SkusDelete(c *gin.Context) {
	var r request.IdsReq
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.DeleteSkus(r)
	utils.ResponseSuccess(c, res)

}
