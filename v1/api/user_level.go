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
func UserLevelList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBind(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error("ManagerList with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.GetUserLevelList(pageInfo)

	utils.ResponseSuccess(c, res)

	return

}

func UserLevelUpdateStatus(c *gin.Context) {
	var r request.UpdateStatusRequest
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}

	res := service.UpdateStatusUserLevel(r)
	utils.ResponseSuccess(c, res)

}

// Create 赠加管理员
func UserLevelCreate(c *gin.Context) {
	var r request.UserLevelRequest
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.CreateUserLevel(r)
	utils.ResponseSuccess(c, res)

}

func UserLevelEdit(c *gin.Context) {
	var r request.UserLevelEditRequest
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}

	res := service.EditUserLevel(r)
	utils.ResponseSuccess(c, res)
}

func UserLevelDelete(c *gin.Context) {
	var r request.GetById
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.DeleteUserLevel(r)
	utils.ResponseSuccess(c, res)

}
