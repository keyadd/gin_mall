package service

import (
	"gin_mall/global"
	"gin_mall/v1/model"
	"gin_mall/v1/model/request"
	"gin_mall/v1/model/response"
	"go.uber.org/zap"
)

// GetUserLevelList 获取UserLevel 用户列表
func GetUserLevelList(r request.PageInfo) (res response.PageResult) {
	UserLevelList, total, err := model.GetUserLevelList(r)
	if err != nil {
		global.GVA_LOG.Error("model.GetUserLevelList(list)", zap.Error(err))
		return
	}
	result := response.PageResult{
		List:  UserLevelList,
		Total: total,
	}

	return result
}

// CreateUserLevel 创建UserLevel用户
func CreateUserLevel(r request.UserLevelRequest) (res response.CreateUserLevel) {

	m := response.CreateUserLevel{
		Name:     r.Name,
		Level:    r.Level,
		Status:   r.Status,
		Discount: r.Discount,
		MaxPrice: r.MaxPrice,
		MaxTimes: r.MaxTimes,
	}

	u, err := model.CreateUserLevel(m)

	if err != nil {
		global.GVA_LOG.Error("model.CreateUserLevel(m)", zap.Any("m", m), zap.Error(err))
		return
	}
	return u
}

func EditUserLevel(r request.UserLevelEditRequest) bool {
	res, err := model.EditUserLevel(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditUserLevel(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}
}

func DeleteUserLevel(r request.GetById) bool {
	res, err := model.DeleteUserLevelId(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditUserLevel(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}

}

func UpdateStatusUserLevel(r request.UpdateStatusRequest) bool {
	res, err := model.UpdateStatusUserLevelId(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditUserLevel(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}
}
