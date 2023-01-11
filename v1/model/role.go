package model

import (
	"gin_mall/global"
	"gin_mall/v1/model/response"
)

func InfoRole(role_id int64) (*response.InfoRole, error) {
	var info_role *response.InfoRole
	db := global.GVA_DB.Table("role")
	err := db.Where("id = ?", role_id).First(&info_role).Error

	return info_role, err
}

func RoleList() ([]*response.InfoRole, error) {
	var info_role []*response.InfoRole
	db := global.GVA_DB.Table("role")
	err := db.Find(&info_role).Error
	return info_role, err
}
