package model

import (
	"gin_mall/global"
	"gin_mall/v1/model/request"
	"gin_mall/v1/model/response"
)

func CreateSkus(s response.Skus) (response.Skus, error) {
	db := global.GVA_DB.Table("skus")
	err := db.Create(&s).Error
	return s, err
}

func EditSkus(r request.EditSkus) (res int64, err error) {
	db := global.GVA_DB.Table("skus")
	result := db.Where("id = ?", r.Id).Updates(r.AddSkus)
	return result.RowsAffected, result.Error

}

func DeleteSkusId(arr []int) (res int64, err error) {
	db := global.GVA_DB.Table("skus")
	var r response.Skus
	result := db.Delete(r, arr)
	return result.RowsAffected, result.Error
}

func UpdateStatusSkusId(r request.UpdateStatusRequest) (res int64, err error) {
	db := global.GVA_DB.Table("skus")
	result := db.Where("id = ?", r.Id).Updates(&r)
	return result.RowsAffected, result.Error
}

func GetSkusList(list request.PageInfo) ([]*response.SkusList, int64, error) {
	var m []*response.SkusList
	limit := list.PageSize
	offset := list.PageSize * (list.Page - 1)
	var total int64
	db := global.GVA_DB.Table("skus")
	_ = db.Count(&total).Error
	err := db.Limit(limit).Offset(offset).Find(&m).Error
	return m, total, err
}

func GetSkusCount() (total int64) {
	db := global.GVA_DB.Table("skus")
	_ = db.Count(&total).Error
	return
}
