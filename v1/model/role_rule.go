package model

import (
	"gin_mall/global"
	"gin_mall/v1/model/response"
)

//type RoleRule struct{}
//
//func (RoleRule) TableName() string {
//	return "role_rule"
//}

func GetRoleRuleList(roleId int64) ([]response.RoleRule, error) {
	var ruleList []response.RoleRule
	db := global.GVA_DB.Table("role_rule")
	err := db.Where("role_id =?", roleId).Find(&ruleList).Error
	return ruleList, err
}

func SetRules(r []response.RoleRule) (res int64, err error) {
	db := global.GVA_DB.Table("role_rule")
	result := db.Create(&r)
	return result.RowsAffected, result.Error
}

func DeleteRules(r []int64) (res int64, err error) {
	roleRule := response.RoleRule{}
	db := global.GVA_DB.Table("role_rule")
	result := db.Delete(&roleRule, r)
	return result.RowsAffected, result.Error
}

func GetRoleRuleIdList(roleId int64) ([]response.RoleRule, error) {
	var ruleList []response.RoleRule
	db := global.GVA_DB.Table("role_rule")
	err := db.Where("role_id =?", roleId).Find(&ruleList).Error
	return ruleList, err
}

func GetRoleRule(roleId int64, ruleId int64) (int64, error) {
	var ruleList response.RoleRule
	db := global.GVA_DB.Table("role_rule")
	err := db.Where("role_id =?", roleId).Where("rule_id =?", ruleId).First(&ruleList).Error
	return ruleList.Id, err
}
