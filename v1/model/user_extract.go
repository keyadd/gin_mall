package model

import "gin_mall/global"

func GetUserExtractCount() (total int64) {
	db := global.GVA_DB.Table("user_extract")
	_ = db.Where("status = ?", 1).Count(&total).Error
	return
}
