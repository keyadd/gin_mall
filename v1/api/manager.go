package api

import (
	"errors"
	"gin_mall/global"
	"gin_mall/utils"
	"gin_mall/v1/model/request"
	"gin_mall/v1/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// Login 后台登录
func Login(c *gin.Context) {
	p := &request.Login{}
	if err := c.ShouldBind(p); err != nil {
		global.GVA_LOG.Error("Login with invalid param", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			utils.ResponseError(c, global.CodeInvalidParam)
			return
		}
		utils.ResponseErrorWithMsgValid(c, global.CodeInvalidParam, errs.Translate(global.TRANS))
		return
	}

	token, err := service.Login(p)

	if nil != err {
		global.GVA_LOG.Error("service.Login(p) failed", zap.Error(err))

		if errors.Is(err, service.ErrorUserNotExist) {
			utils.ResponseError(c, global.CodeUserNotExist)
			return
		}
		utils.ResponseError(c, global.CodeInvalidPassword)
		return
	}
	//返回响应
	utils.ResponseSuccess(c, token)

}

// GetManagerInfo 获取指定用户的详细信息
func GetManagerInfo(c *gin.Context) {

	id, _ := c.Get("userID")
	i := id.(int64)
	res, _ := service.GetManagerInfo(i)
	utils.ResponseSuccess(c, res)

}

func ChangePassword(c *gin.Context) {
	var r request.UpdatePasswordRequest
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("ManagerList with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	id, _ := c.Get("userID")
	i := id.(int64)
	res := service.ChangePassword(i, r)
	utils.ResponseSuccess(c, res)

}

func Logout(c *gin.Context) {
	id, _ := c.Get("userID")
	i := id.(int64)
	res := service.LogoutManager(i)
	utils.ResponseSuccess(c, res)

}

// List 获取管理员列表
func List(c *gin.Context) {
	var pageInfo request.ManagerList
	err := c.ShouldBind(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error("ManagerList with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res, _ := service.GetManagerList(pageInfo)

	utils.ResponseSuccess(c, res)

	return

}

func UpdateStatus(c *gin.Context) {
	var r request.UpdateStatusRequest
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}

	res := service.UpdateStatusManager(r)
	utils.ResponseSuccess(c, res)

}

// Create 赠加管理员
func Create(c *gin.Context) {
	var r request.ManagerRequest
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res, _ := service.CreateManager(r)
	utils.ResponseSuccess(c, res)

}

func Edit(c *gin.Context) {
	var r request.EditRequest
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}

	res := service.EditManager(r)
	utils.ResponseSuccess(c, res)
}

func Delete(c *gin.Context) {
	var r request.GetById
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.DeleteManager(r)
	utils.ResponseSuccess(c, res)

}
