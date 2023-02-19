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

// GetManagerList 获取manager 用户列表
func GetMenuList() (res response.MenuResponse) {
	menu, err := model.GetRuleMenuList()
	if err != nil {
		global.GVA_LOG.Error("model.GetRuleMenu(rule_id) failed", zap.Error(err))
		return
	}
	menuTree := utils.MenuTree(menu, 0)

	menuResponse := response.MenuResponse{
		List: menuTree,
	}
	return menuResponse
}

// CreateManager 创建manager用户
func CreateMenu(r request.MenuRequest) (res response.Menu) {

	menu := response.Menu{
		Status:     r.Status,
		CreateTime: time.Now().Unix(),
		RuleId:     r.RuleId,
		Name:       r.Name,
		FrontPath:  r.FrontPath,
		Menu:       r.Menu,
		Order:      r.Order,
		Icon:       r.Icon,
	}

	res, err := model.CreateRule(menu)

	if err != nil {
		global.GVA_LOG.Error("model.CreateRule(rule)", zap.Any("rule", menu), zap.Error(err))
		return
	}
	return
}

func EditMenu(r request.EditMenuRequest) bool {
	res, err := model.EditRule(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditManager(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}
}

func DeleteMenu(r request.GetById) bool {
	res, err := model.DeleteRuleId(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditManager(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}

}

func UpdateStatusMenu(r request.UpdateStatusRequest) bool {
	res, err := model.UpdateStatusRuleId(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditManager(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}
}
