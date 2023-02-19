package model

import (
	"gin_mall/global"
	"gin_mall/v1/model/request"
	"gin_mall/v1/model/response"
)

func GetRuleMenu(rule_id []int64) (Menu []*response.Menu, err error) {

	var menu []*response.Menu
	db := global.GVA_DB.Table("rule")
	err = db.Where("id in ?", rule_id).Where("menu=?", 1).Find(&menu).Error
	return menu, err
}

func CreateRule(rule response.Menu) (response.Menu, error) {
	db := global.GVA_DB.Table("rule")
	err := db.Create(&rule).Error
	return rule, err
}

func EditRule(rule request.EditMenuRequest) (res int64, err error) {
	db := global.GVA_DB.Table("rule")
	result := db.Where("id = ?", rule.Id).Updates(rule.MenuRequest)
	return result.RowsAffected, result.Error

}

func DeleteRuleId(r request.GetById) (res int64, err error) {
	db := global.GVA_DB.Table("rule")
	result := db.Delete(&r)
	return result.RowsAffected, result.Error
}

func UpdateStatusRuleId(r request.UpdateStatusRequest) (res int64, err error) {
	db := global.GVA_DB.Table("rule")
	result := db.Where("id = ?", r.Id).Updates(&r)
	return result.RowsAffected, result.Error
}

func GetRuleMenuList() (Menu []*response.Menu, err error) {

	var menu []*response.Menu
	db := global.GVA_DB.Table("rule")
	err = db.Table("rule").Where("menu=?", 1).Find(&menu).Error
	return menu, err
}

//api管理

func CreateRuleApi(rule response.Api) (response.Api, error) {
	db := global.GVA_DB.Table("rule")
	err := db.Create(&rule).Error
	return rule, err
}

func EditRuleApi(rule request.EditApiRequest) (res int64, err error) {
	db := global.GVA_DB.Table("rule")
	result := db.Where("id = ?", rule.Id).Updates(rule.CreateApiRequest)
	return result.RowsAffected, result.Error

}

func GetRuleApiList(list request.ApiRequest) (Menu []*response.ApiList, total int64, err error) {
	limit := list.PageSize
	offset := list.PageSize * (list.Page - 1)

	var api []*response.ApiList
	db := global.GVA_DB.Table("rule")

	if list.Keyword != "" {
		db = db.Where("router_path like ?", "%"+list.Keyword+"%")
	}
	_ = db.Where("menu=?", 0).Count(&total)

	err = db.Limit(limit).Offset(offset).Where("menu=?", 0).Find(&api).Error
	return api, total, err
}

func GetApiGroup() (Menu []*response.Api, err error) {
	var api []*response.Api
	db := global.GVA_DB.Table("rule")
	err = db.Where("menu=?", 0).Where("api_group !=?", "").Find(&api).Error
	return api, err
}

func GetRuleId(path string, method string) (ruleId int64, err error) {
	var api response.Api
	db := global.GVA_DB.Table("rule")
	err = db.Select("id").Where("menu=?", 0).Where("router_path =?", path).Where("method =?", method).First(&api).Error
	return api.Id, err
}

func GetRuleName(path string, method string) (name string, err error) {
	var api response.Api
	db := global.GVA_DB.Table("rule")
	err = db.Select("name").Where("menu=?", 0).Where("router_path =?", path).Where("method =?", method).First(&api).Error
	return api.Name, err
}
