package service

import (
	"gin_mall/global"
	"gin_mall/utils"
	"gin_mall/v1/model"
	"gin_mall/v1/model/request"
	"gin_mall/v1/model/response"
	"go.uber.org/zap"
	"time"
)

// GetGoodsCommentList 获取GoodsComment 用户列表
func GetGoodsCommentList(r request.ManagerList) (res response.PageResult) {
	GoodsCommentList, total, err := model.GetGoodsCommentList(r)
	if err != nil {
		global.GVA_LOG.Error("model.GetGoodsCommentList(list)", zap.Error(err))
		return
	}

	result := response.PageResult{
		List:  GoodsCommentList,
		Total: total,
	}
	return result
}

// CreateGoodsComment 创建GoodsComment用户
func CreateGoodsComment(r request.GoodsCommentRequest) (res response.CreateGoodsComment) {
	password := utils.MD5V([]byte(r.Password))
	m := response.CreateGoodsComment{
		GoodsCommentname:    r.GoodsCommentname,
		CreateTime:          time.Now().Unix(),
		Status:              r.Status,
		Password:            password,
		Phone:               r.Phone,
		Email:               r.Email,
		Avatar:              r.Avatar,
		GoodsCommentLevelId: r.GoodsCommentLevelId,
	}

	u, err := model.CreateGoodsComment(m)

	if err != nil {
		global.GVA_LOG.Error("model.CreateGoodsComment(m)", zap.Any("m", m), zap.Error(err))
		return
	}
	return u
}

func EditGoodsComment(r request.GoodsCommentEditRequest) bool {
	r.Password = utils.MD5V([]byte(r.Password))
	res, err := model.EditGoodsComment(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditGoodsComment(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}
}

func DeleteGoodsComment(r request.GetById) bool {
	res, err := model.DeleteGoodsCommentId(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditGoodsComment(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}

}

func UpdateStatusGoodsComment(r request.UpdateStatusRequest) bool {
	res, err := model.UpdateStatusGoodsCommentId(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditGoodsComment(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}
}
