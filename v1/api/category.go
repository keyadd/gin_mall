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
func CategoryList(c *gin.Context) {

	res := service.GetCategoryList()

	utils.ResponseSuccess(c, res)

	return

}

func CategoryUpdateStatus(c *gin.Context) {
	var r request.UpdateStatusRequest
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}

	res := service.UpdateStatusCategory(r)
	utils.ResponseSuccess(c, res)

}

// Create 赠加管理员
func CategoryCreate(c *gin.Context) {
	var r request.CategoryRequest
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.CreateCategory(r)
	utils.ResponseSuccess(c, res)

}

func CategoryEdit(c *gin.Context) {
	var r request.CategoryEditRequest
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}

	res := service.EditCategory(r)
	utils.ResponseSuccess(c, res)
}

func CategoryDelete(c *gin.Context) {
	var r request.GetById
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.DeleteCategory(r)
	utils.ResponseSuccess(c, res)

}
