package model

import (
	"gin_mall/global"
	"gin_mall/v1/model/request"
	"gin_mall/v1/model/response"
)

func GetImageCount(imageClassId int64) (imageCount int64) {
	db := global.GVA_DB.Table("image")
	_ = db.Where("image_class_id = ?", imageClassId).Count(&imageCount).Error
	return
}

func GetImageList(r request.ImageClassId) (res []*response.Image, err error) {

	limit := r.PageSize
	offset := r.PageSize * (r.Page - 1)
	db := global.GVA_DB.Table("image")
	db = db.Where("image_class_id = ?", r.Id)
	err = db.Limit(limit).Offset(offset).Find(&res).Error
	return
}
