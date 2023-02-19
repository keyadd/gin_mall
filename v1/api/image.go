package api

import (
	"gin_mall/global"
	"gin_mall/utils"
	"gin_mall/v1/model/request"
	"gin_mall/v1/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ImageUpload(c *gin.Context) {
	var imageClassId request.ImageClassIdRequest
	err := c.ShouldBind(&imageClassId)
	if err != nil {
		global.GVA_LOG.Error("ManagerList with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	form, _ := c.MultipartForm()
	files := form.File["img"]

	images := service.AddImage(files, imageClassId)
	utils.ResponseSuccess(c, images)

}

func ImageRename(c *gin.Context) {
	var rename request.ImageRename
	err := c.ShouldBind(&rename)
	if err != nil {
		global.GVA_LOG.Error("ManagerList with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.RenameImage(rename)
	utils.ResponseSuccess(c, res)

}

func ImageDelete(c *gin.Context) {
	var r request.IdsReq
	err := c.ShouldBind(&r)
	if err != nil {
		global.GVA_LOG.Error("c.ShouldBind(&r) with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}
	res := service.DeleteImage(r)
	utils.ResponseSuccess(c, res)
}
