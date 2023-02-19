package model

import (
	"gin_mall/global"
	"gin_mall/v1/model/request"
	"gin_mall/v1/model/response"
)

func CreateOperationRecord(OperationRecord response.OperationRecord) (err error) {
	err = global.GVA_DB.Table("operation_record").Create(&OperationRecord).Error
	return err
}

func OperationRecordList(r request.OperationRecord) (res []response.OperationRecordList, total int64, err error) {
	limit := r.PageSize
	offset := r.PageSize * (r.Page - 1)

	db := global.GVA_DB.Table("operation_record")
	db = db.Select("operation_record.*, manager.username").Joins("left join manager on operation_record.manager_id = manager.id")
	db = db.Where("operation_record.delete_time = ?", 0)

	if r.Keyword != "" {
		db = db.Where("manager.username like ?", "%"+r.Keyword+"%")
	}
	_ = db.Count(&total).Error
	//db = db.First(&re)
	err = db.Limit(limit).Offset(offset).Find(&res).Error
	return res, total, err
}

func GetOperationRecordCount() (total int64) {
	db := global.GVA_DB.Table("operation_record")
	_ = db.Count(&total).Error
	return
}

func DeleteOperationRecordId(r request.GetById, l response.OperationRecord) (res int64, err error) {
	db := global.GVA_DB.Table("operation_record")
	result := db.Where("id = ?", r.Id).Updates(l)
	return result.RowsAffected, result.Error
}
