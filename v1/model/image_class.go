package model

import (
	"fmt"
	"gin_mall/global"
	"gin_mall/v1/model/request"
	"gin_mall/v1/model/response"
)

func ImageClassList(page request.ImageClassList) (res []*response.ImageClass, err error) {
	limit := page.PageSize
	offset := page.PageSize * (page.Page - 1)
	db := global.GVA_DB.Table("image_class")
	err = db.Limit(limit).Offset(offset).Find(&res).Error
	return res, err
}

func GetImageClassCount() (total int64) {
	db := global.GVA_DB.Table("image_class")
	_ = db.Count(&total).Error
	return
}

func AddImageClass(r response.ImageClass) (res response.ImageClass, err error) {
	fmt.Println(r)
	db := global.GVA_DB.Table("image_class")
	err = db.Create(&r).Error
	fmt.Println(r.Name)
	return r, err
}

func EditImageClass(r request.EditImageClass) (res int64, err error) {
	db := global.GVA_DB.Table("image_class")
	result := db.Where("id = ?", r.Id).Updates(r)
	return result.RowsAffected, result.Error
}

func DeleteImageClassId(r request.GetById) (res int64, err error) {
	db := global.GVA_DB.Table("image_class")
	result := db.Delete(&r)
	return result.RowsAffected, result.Error
}
