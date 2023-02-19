package model

import "gin_mall/global"

func GetUserBillCount() (total int64) {
	db := global.GVA_DB.Table("user_bill")
	_ = db.Where("status = ?", 1).Count(&total).Error
	return
}
