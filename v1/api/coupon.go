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
func CouponList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBind(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error("ManagerList with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.GetCouponList(pageInfo)

	utils.ResponseSuccess(c, res)

	return

}

func CouponUpdateStatus(c *gin.Context) {
	var r request.UpdateStatusRequest
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}

	res := service.UpdateStatusCoupon(r)
	utils.ResponseSuccess(c, res)

}

// Create 赠加管理员
func CouponCreate(c *gin.Context) {
	var r request.AddCoupon
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.CreateCoupon(r)
	utils.ResponseSuccess(c, res)

}

func CouponEdit(c *gin.Context) {
	var r request.EditCoupon
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}

	res := service.EditCoupon(r)
	utils.ResponseSuccess(c, res)
}

func CouponDelete(c *gin.Context) {
	var r request.IdsReq
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.DeleteCoupon(r)
	utils.ResponseSuccess(c, res)

}
