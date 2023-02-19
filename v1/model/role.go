package model

import (
	"gin_mall/global"
	"gin_mall/v1/model/request"
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

func CreateRole(m response.CreateRole) (response.CreateRole, error) {
	db := global.GVA_DB.Table("role")
	err := db.Create(&m).Error
	return m, err
}

func EditRole(r request.EditRole) (res int64, err error) {
	db := global.GVA_DB.Table("role")
	result := db.Where("id = ?", r.Id).Updates(r.AddRole)
	return result.RowsAffected, result.Error

}

func DeleteRoleId(r request.GetById) (res int64, err error) {
	db := global.GVA_DB.Table("role")
	result := db.Delete(&r)
	return result.RowsAffected, result.Error
}

func UpdateStatusRoleId(r request.UpdateStatusRequest) (res int64, err error) {
	db := global.GVA_DB.Table("role")
	result := db.Where("id = ?", r.Id).Updates(&r)
	return result.RowsAffected, result.Error
}

func GetRoleList(list request.PageInfo) ([]*response.Role, int64, error) {
	var m []*response.Role
	var total int64
	limit := list.PageSize
	offset := list.PageSize * (list.Page - 1)

	db := global.GVA_DB.Table("role")
	_ = db.Count(&total).Error
	err := db.Limit(limit).Offset(offset).Preload("RoleRule").Find(&m).Error
	return m, total, err
}

func GetRoleCount() (total int64) {
	db := global.GVA_DB.Table("role")
	_ = db.Count(&total).Error
	return
}
