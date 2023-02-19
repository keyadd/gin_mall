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
func OrderList(c *gin.Context) {
	var orderList request.OrderListRequest
	err := c.ShouldBind(&orderList)
	if err != nil {
		global.GVA_LOG.Error("ManagerList with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.GetOrderList(orderList)

	utils.ResponseSuccess(c, res)

	return

}

func OrderHandleRefund(c *gin.Context) {
	var r request.OrderHandleRefund
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}

	res := service.OrderHandleRefund(r)
	utils.ResponseSuccess(c, res)

}

func OrderSend(c *gin.Context) {
	var r request.OrderSend
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}

	res := service.OrderSend(r)
	utils.ResponseSuccess(c, res)

}

func OrderDelete(c *gin.Context) {
	var r request.IdsReq
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.DeleteOrder(r)
	utils.ResponseSuccess(c, res)

}
