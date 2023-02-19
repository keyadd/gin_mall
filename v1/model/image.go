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

func UploadImage(image response.Image) (res response.Image, err error) {
	db := global.GVA_DB.Table("image")
	err = db.Create(&image).Error
	return image, err
}

func RenameImage(r request.ImageRename) (res int64, err error) {
	db := global.GVA_DB.Table("image")
	result := db.Where("id = ?", r.Id).Updates(r)
	return result.RowsAffected, result.Error
}

func DeleteImage(arr []int) (res int64, err error) {
	var r response.Image
	db := global.GVA_DB.Table("image")
	result := db.Delete(r, arr)
	return result.RowsAffected, result.Error
}
