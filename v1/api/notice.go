package api

import (
	"gin_mall/global"
	"gin_mall/utils"
	"gin_mall/v1/model/request"
	"gin_mall/v1/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NoticeList(c *gin.Context) {
	var r request.NoticeList
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("ManagerList with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	list := service.GetNoticeList(r)
	utils.ResponseSuccess(c, list)
}

func NoticeCreate(c *gin.Context) {

	var r request.NoticeRequest
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("ManagerList with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	result, _ := service.AddNotice(r)
	utils.ResponseSuccess(c, result)

}
func NoticeEdit(c *gin.Context) {
	var r request.NoticeEdit
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("ManagerList with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	result := service.EditNotice(r)
	utils.ResponseSuccess(c, result)

}
func NoticeDelete(c *gin.Context) {
	var r request.GetById
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.DeleteNotice(r)
	utils.ResponseSuccess(c, res)

}
