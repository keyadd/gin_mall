package api

import (
	"gin_mall/global"
	"gin_mall/utils"
	"gin_mall/v1/model/request"
	"gin_mall/v1/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GoodsSkusCardValueSet(c *gin.Context) {
	var r request.UpdateStatusRequest
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}

	res := service.SetGoodsSkusCardValue(r)
	utils.ResponseSuccess(c, res)

}

// GoodsSkusCardValueCreate 添加菜单或者权限
func GoodsSkusCardValueCreate(c *gin.Context) {
	var r request.GoodsSkusCardValueRequest
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.CreateGoodsSkusCardValue(r)
	utils.ResponseSuccess(c, res)

}

// GoodsSkusCardValueEdit 编辑菜单或者权限
func GoodsSkusCardValueEdit(c *gin.Context) {
	var r request.EditGoodsSkusCardValueRequest
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}

	res := service.EditGoodsSkusCardValue(r)
	utils.ResponseSuccess(c, res)
}

func GoodsSkusCardValueDelete(c *gin.Context) {
	var r request.GetById
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.DeleteGoodsSkusCardValue(r)
	utils.ResponseSuccess(c, res)

}
