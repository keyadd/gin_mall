package model

import (
	"gin_mall/global"
	"gin_mall/v1/model/request"
	"gin_mall/v1/model/response"
)

type UserLevelLevel struct {
	Id       int64  `json:"id" `
	Name     string `json:"name"`
	Level    int64  `json:"level"`
	Status   int64  `json:"status"`
	Discount int64  `json:"discount"`
	MaxPrice int64  `json:"max_price"`
	MaxTimes int64  `json:"max_times"`
}

func GetUserLevelIds(ids []int64) (res []response.UserLevelList) {
	db := global.GVA_DB.Table("user_level")
	_ = db.Where("id in ?", ids).Find(&res).Error
	return
}

func GetUserLevel() ([]*response.UserLevelList, error) {
	var m []*response.UserLevelList

	db := global.GVA_DB.Table("user_level")
	err := db.Find(&m).Error
	return m, err
}
func GetUserLevelList(list request.PageInfo) ([]*response.UserLevel, int64, error) {
	var m []*response.UserLevel

	limit := list.PageSize
	offset := list.PageSize * (list.Page - 1)
	var total int64
	db := global.GVA_DB.Table("user_level")
	_ = db.Count(&total).Error
	err := db.Limit(limit).Offset(offset).Find(&m).Error
	return m, total, err
}

func CreateUserLevel(m response.CreateUserLevel) (response.CreateUserLevel, error) {
	db := global.GVA_DB.Table("user_level")
	err := db.Create(&m).Error
	return m, err
}

func EditUserLevel(r request.UserLevelEditRequest) (res int64, err error) {
	db := global.GVA_DB.Table("user_level")
	result := db.Where("id = ?", r.Id).Updates(r.UserLevelRequest)
	return result.RowsAffected, result.Error

}

func DeleteUserLevelId(r request.GetById) (res int64, err error) {
	db := global.GVA_DB.Table("user_level")
	result := db.Delete(&r)
	return result.RowsAffected, result.Error
}

func UpdateStatusUserLevelId(r request.UpdateStatusRequest) (res int64, err error) {
	db := global.GVA_DB.Table("user_level")
	result := db.Where("id = ?", r.Id).Updates(&r)
	return result.RowsAffected, result.Error
}
