package api

import (
	"gin_mall/global"
	"gin_mall/utils"
	"gin_mall/v1/model/request"
	"gin_mall/v1/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ImageClassList(c *gin.Context) {
	var r request.ImageClassList
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("ManagerList with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	list := service.GetImageClassList(r)
	utils.ResponseSuccess(c, list)
}

func ImageList(c *gin.Context) {
	var r request.ImageClassId
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("ManagerList with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	result, _ := service.GetIdImageList(r)
	utils.ResponseSuccess(c, result)

}

func ImageClassCreate(c *gin.Context) {

	var r request.CreateImageClass
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("ManagerList with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	result, _ := service.AddImageClass(r)
	utils.ResponseSuccess(c, result)

}
func ImageClassEdit(c *gin.Context) {

}
func ImageClassDelete(c *gin.Context) {

}
