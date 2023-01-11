package model

import (
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

func AddImageClass(r request.CreateImageClass) (res response.ImageClass, err error) {
	db := global.GVA_DB.Table("image_class")
	err = db.Create(&res).Error
	return
}
