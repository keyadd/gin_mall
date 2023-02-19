package service

import (
	"gin_mall/global"
	"gin_mall/v1/model"
	"go.uber.org/zap"
)

func GetRule(path string, method string, userId int64) (res bool) {

	ruleId, err := model.GetRuleId(path, method)
	if err != nil {
		global.GVA_LOG.Error("model.GetRuleId(path,method)", zap.Any("path", path), zap.Any("method", method), zap.Error(err))
		return
	}
	info, err := model.GetManagerInfo(userId)

	if err != nil {
		global.GVA_LOG.Error("model.GetManagerInfo(userId)", zap.Any("userId", userId), zap.Error(err))
		return
	}

	_, err = model.GetRoleRule(info.RoleId, ruleId)
	if err != nil {
		global.GVA_LOG.Error("model.GetRoleRule(info.RoleId, ruleId))", zap.Any("ruleId", ruleId), zap.Error(err))
		return
	}

	return true
}
