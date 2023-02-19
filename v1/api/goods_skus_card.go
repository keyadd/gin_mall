package api

import (
	"gin_mall/global"
	"gin_mall/utils"
	"gin_mall/v1/model/request"
	"gin_mall/v1/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GoodsSkusCardUpdateOrder(c *gin.Context) {
	var r request.UpdateStatusRequest
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}

	res := service.GoodsSkusCardUpdateOrder(r)
	utils.ResponseSuccess(c, res)

}

// GoodsSkusCardCreate 添加菜单或者权限
func GoodsSkusCardCreate(c *gin.Context) {
	var r request.GoodsSkusCardRequest
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.CreateGoodsSkusCard(r)
	utils.ResponseSuccess(c, res)

}

// GoodsSkusCardEdit 编辑菜单或者权限
func GoodsSkusCardEdit(c *gin.Context) {
	var r request.EditGoodsSkusCardRequest
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}

	res := service.EditGoodsSkusCard(r)
	utils.ResponseSuccess(c, res)
}

func GoodsSkusCardDelete(c *gin.Context) {
	var r request.GetById
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.DeleteGoodsSkusCard(r)
	utils.ResponseSuccess(c, res)

}
