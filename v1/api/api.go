package api

import (
	"gin_mall/global"
	"gin_mall/utils"
	"gin_mall/v1/model/request"
	"gin_mall/v1/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ApiList List 获取管理员列表
func ApiList(c *gin.Context) {
	var r request.ApiRequest
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("ManagerList with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.GetApiList(r)
	utils.ResponseSuccess(c, res)
	return
}
func ApiListTree(c *gin.Context) {
	res := service.GetApiListTree()
	utils.ResponseSuccess(c, res)
	return
}

func ApiUpdateStatus(c *gin.Context) {
	var r request.UpdateStatusRequest
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}

	res := service.UpdateStatusApi(r)
	utils.ResponseSuccess(c, res)

}

// ApiCreate 添加菜单或者权限
func ApiCreate(c *gin.Context) {
	var r request.CreateApiRequest
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.CreateApi(r)
	utils.ResponseSuccess(c, res)

}

// ApiEdit 编辑菜单或者权限
func ApiEdit(c *gin.Context) {
	var r request.EditApiRequest
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}

	res := service.EditApi(r)
	utils.ResponseSuccess(c, res)
}

func ApiDelete(c *gin.Context) {
	var r request.GetById
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.DeleteApi(r)
	utils.ResponseSuccess(c, res)

}
