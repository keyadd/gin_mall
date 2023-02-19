package api

import (
	"gin_mall/global"
	"gin_mall/utils"
	"gin_mall/v1/model/request"
	"gin_mall/v1/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GoodsList 获取管理员列表
func GoodsList(c *gin.Context) {
	var l request.GoodsList
	err := c.ShouldBind(&l)
	if err != nil {
		global.GVA_LOG.Error("ManagerList with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.GetGoodsList(l)

	utils.ResponseSuccess(c, res)

	return

}

func GoodsChangeStatus(c *gin.Context) {
	var r request.StatusIdsRequest
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}

	res := service.ChangeStatusGoods(r)
	utils.ResponseSuccess(c, res)

}

// Create 赠加管理员
func GoodsCreate(c *gin.Context) {
	var r request.AddGoods
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.CreateGoods(r)
	utils.ResponseSuccess(c, res)

}

func GoodsEdit(c *gin.Context) {
	var r request.EditGoods
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}

	res := service.EditGoods(r)
	utils.ResponseSuccess(c, res)
}

func GoodsDeleteAll(c *gin.Context) {
	var r request.IdsReq
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.DeleteAllGoods(r)
	utils.ResponseSuccess(c, res)

}

func GoodsDestroyAll(c *gin.Context) {
	var r request.IdsReq
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.DestroyAllGoods(r)
	utils.ResponseSuccess(c, res)

}

func GoodsCheck(c *gin.Context) {
	var r request.GoodsCheck
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.GoodsCheck(r)
	utils.ResponseSuccess(c, res)
}

func GoodsRead(c *gin.Context) {
	var r request.GetById
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.GoodsRead(r)
	utils.ResponseSuccess(c, res)

}

func GoodsBannerSet(c *gin.Context) {
	var r request.GoodsBannerSet
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.GoodsBannerSet(r)
	utils.ResponseSuccess(c, res)

}
func GoodsUpdateSkus(c *gin.Context) {
	var r request.GoodsUpdateSkus
	err := c.ShouldBindJSON(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.GoodsUpdateSkus(r)
	utils.ResponseSuccess(c, res)

}
