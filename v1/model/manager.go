package model

import (
	"fmt"
	"gin_mall/global"
	"gin_mall/v1/model/request"
	"gin_mall/v1/model/response"
)

func GetManger(user *request.Login) (m *response.Manager, err error) {

	db := global.GVA_DB.Table("manager")
	manager := new(response.Manager)
	db.Where("username=?", user.Username).First(manager)
	return manager, err

}

func GetManagerInfo(id int64) (UserInfo *response.Manager, err error) {
	db := global.GVA_DB.Table("manager")
	user := new(response.Manager)
	db.Where("id=?", id).First(user)
	return user, nil
}

func GetManagerList(list request.ManagerList) ([]*response.Manager, error) {
	var m []*response.Manager
	limit := list.PageSize
	offset := list.PageSize * (list.Page - 1)

	db := global.GVA_DB.Table("manager")
	db = db.Where("username like ?", "%"+list.Keyword+"%")
	err := db.Limit(limit).Offset(offset).Find(&m).Error
	return m, err
}

func GetManagerCount() (total int64) {
	db := global.GVA_DB.Table("manager")
	_ = db.Count(&total).Error
	return
}

func CreateManager(m response.Manager) (response.Manager, error) {
	db := global.GVA_DB.Table("manager")
	err := db.Create(&m).Error
	return m, err
}

func EditManager(r request.EditRequest) (res int64, err error) {
	db := global.GVA_DB.Table("manager")
	result := db.Where("id = ?", r.Id).Updates(r.ManagerRequest)
	return result.RowsAffected, result.Error

}

func DeleteManagerId(r request.GetById) (res int64, err error) {
	db := global.GVA_DB.Table("manager")
	result := db.Delete(&r)
	return result.RowsAffected, result.Error
}

func UpdateStatusManagerId(r request.UpdateStatusRequest) (res int64, err error) {
	db := global.GVA_DB.Table("manager")
	result := db.Where("id = ?", r.Id).Updates(&r)
	return result.RowsAffected, result.Error
}

func UpdatePassword(i int64, r request.UpdatePasswordRequest) (res int64, err error) {
	fmt.Println(r)
	db := global.GVA_DB.Table("manager")
	result := db.Where("id = ?", i).Where("password = ?", r.OldPassword).Update("password", r.Password)

	return result.RowsAffected, result.Error
}
