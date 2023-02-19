package model

import (
	"gin_mall/global"
	"gin_mall/v1/model/request"
	"gin_mall/v1/model/response"
)

func GetAppCategoryItemList(r request.AppCategoryItemRequest) ([]*response.AppCategoryItem, error) {
	var m []*response.AppCategoryItem

	db := global.GVA_DB.Table("app_category_item")
	db = db.Where("category_id = ?", r.Id)

	err := db.Find(&m).Error
	return m, err
}

func CreateAppCategoryItem(m response.AppCategoryItem) (response.AppCategoryItem, error) {
	db := global.GVA_DB.Table("app_category_item")
	err := db.Create(&m).Error
	return m, err
}

func DeleteAppCategoryItemId(r request.GetById) (res int64, err error) {
	db := global.GVA_DB.Table("app_category_item")
	result := db.Delete(&r)
	return result.RowsAffected, result.Error
}
