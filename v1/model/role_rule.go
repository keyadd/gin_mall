package model

import (
	"gin_mall/global"
	"gin_mall/v1/model/response"
)

func GetManagerRuleID(role_id int64) ([]response.RoleRule, error) {
	var ruleList []response.RoleRule
	db := global.GVA_DB.Table("role_rule")
	err := db.Select("rule_id").Where("role_id =?", role_id).Find(&ruleList).Error
	return ruleList, err
}
