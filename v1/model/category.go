package model

import (
	"gin_mall/global"
	"gin_mall/v1/model/request"
	"gin_mall/v1/model/response"
)

func GetCategoryList() ([]*response.CategoryList, error) {
	var m []*response.CategoryList
	db := global.GVA_DB.Table("category")
	err := db.Find(&m).Error
	return m, err
}

func GetCategoryArr(ids []int64) ([]*response.Category, error) {
	var m []*response.Category
	db := global.GVA_DB.Table("category")
	err := db.Where("id in ?", ids).Find(&m).Error
	return m, err
}

func CreateCategory(m response.Category) (response.Category, error) {
	db := global.GVA_DB.Table("category")
	err := db.Create(&m).Error
	return m, err
}

func EditCategory(r request.CategoryEditRequest) (res int64, err error) {
	db := global.GVA_DB.Table("category")
	result := db.Where("id = ?", r.Id).Updates(r.CategoryRequest)
	return result.RowsAffected, result.Error

}

func DeleteCategoryId(r request.GetById) (res int64, err error) {
	db := global.GVA_DB.Table("category")
	result := db.Delete(&r)
	return result.RowsAffected, result.Error
}

func UpdateStatusCategoryId(r request.UpdateStatusRequest) (res int64, err error) {
	db := global.GVA_DB.Table("category")
	result := db.Where("id = ?", r.Id).Updates(&r)
	return result.RowsAffected, result.Error
}
