package model

import (
	"gin_mall/global"
	"gin_mall/v1/model/response"
)

func GetRuleMenu(rule_id []int64) (Menu []*response.Menu, err error) {

	var menu []*response.Menu
	db := global.GVA_DB.Table("rule")
	err = db.Table("rule").Where("rule_id in ?", rule_id).Where("menu=?", 1).Find(&menu).Error
	return menu, err
}

func GetRuleMenuItem(menu_id int64) ([]*response.Menu, error) {
	//定义子节点目录
	var nodes []*response.Menu
	db := global.GVA_DB.Table("rule")
	err := db.Where("rule_id = ?", menu_id).Find(&nodes).Error
	return nodes, err
}

func GetRuleNames(rule_id []int64) ([]*response.RuleNames, error) {
	var res []*response.RuleNames
	db := global.GVA_DB.Table("rule")
	err := db.Table("rule").Where("rule_id in ?", rule_id).Where("menu=?", 0).Find(&res).Error
	return res, err

}
