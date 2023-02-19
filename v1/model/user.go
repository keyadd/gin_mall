package model

import (
	"gin_mall/global"
	"gin_mall/v1/model/request"
	"gin_mall/v1/model/response"
)

type User struct{}

func (User) TableName() string {
	return "user"
}

func GetUserCount() (total int64) {
	db := global.GVA_DB.Table("user")
	_ = db.Count(&total).Error
	return
}
func GetUserList(list request.UserListRequest) ([]*response.User, error) {
	var m []*response.User

	limit := list.PageSize
	offset := list.PageSize * (list.Page - 1)

	db := global.GVA_DB.Table("user")
	db = db.Where("username like ?", "%"+list.Keyword+"%")
	err := db.Limit(limit).Offset(offset).Find(&m).Error
	return m, err

}

func CreateUser(m response.CreateUser) (response.CreateUser, error) {
	db := global.GVA_DB.Table("user")
	err := db.Create(&m).Error
	return m, err
}

func EditUser(r request.UserEditRequest) (res int64, err error) {
	db := global.GVA_DB.Table("user")
	result := db.Where("id = ?", r.Id).Updates(r.UserRequest)
	return result.RowsAffected, result.Error

}

func DeleteUserId(r request.GetById) (res int64, err error) {
	db := global.GVA_DB.Table("user")
	result := db.Delete(&r)
	return result.RowsAffected, result.Error
}

func UpdateStatusUserId(r request.UpdateStatusRequest) (res int64, err error) {
	db := global.GVA_DB.Table("user")
	result := db.Where("id = ?", r.Id).Updates(&r)
	return result.RowsAffected, result.Error
}

func GetUserOrderPrice() (res float64) {
	db := global.GVA_DB.Table("user")
	_ = db.Pluck("sum(order_price) as price", &res)
	return
}
func GetUserAgentList(list request.AgentList) ([]*response.AgentList, error) {
	var m []*response.AgentList

	limit := list.PageSize
	offset := list.PageSize * (list.Page - 1)

	db := global.GVA_DB.Table("user")
	//db = db.Where("username like ?", "%"+list.Keyword+"%")
	err := db.Limit(limit).Offset(offset).Preload("UserInfo").Find(&m).Error
	return m, err

}
