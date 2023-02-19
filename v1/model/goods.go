package model

import (
	"gin_mall/global"
	"gin_mall/v1/model/request"
	"gin_mall/v1/model/response"
)

func GetGoodsCount() (total int64) {
	db := global.GVA_DB.Table("goods")
	_ = db.Count(&total).Error
	return
}

func CreateGoods(s response.GoodsCreate) (response.GoodsCreate, error) {
	db := global.GVA_DB.Table("goods")
	err := db.Create(&s).Error
	return s, err
}

func EditGoods(r request.EditGoods) (res int64, err error) {
	db := global.GVA_DB.Table("goods")
	result := db.Where("id = ?", r.Id).Updates(r.AddGoods)
	return result.RowsAffected, result.Error

}

func DeleteAllGoodsId(arr []int) (res int64, err error) {
	db := global.GVA_DB.Table("goods")
	var r response.GoodsList
	result := db.Delete(r, arr)
	return result.RowsAffected, result.Error
}

func DestroyGoodsId(arr []int) (res int64, err error) {
	db := global.GVA_DB.Table("goods")
	var r response.GoodsList
	result := db.Delete(r, arr)
	return result.RowsAffected, result.Error
}

func ChangeStatusGoodsId(r request.StatusIdsRequest) (res int64, err error) {
	db := global.GVA_DB.Table("goods")
	result := db.Where("id in ?", r.Ids).Updates(&r)
	return result.RowsAffected, result.Error
}

func GetGoodsList(list request.GoodsList) ([]*response.GoodsList, int64, error) {
	var m []*response.GoodsList
	var total int64
	limit := list.PageSize
	offset := list.PageSize * (list.Page - 1)

	db := global.GVA_DB.Table("goods")

	if list.Title != "" {
		db = db.Where("title like ?", "%"+list.Title+"%")
	}
	if list.Tab == "all" {
		db = db.Where("delete_time = ?", 0)
	}
	// 审核中
	if list.Tab == "checking" {
		db = db.Where("ischeck = ?", 0)
	}
	if list.Tab == "saling" {
		db = db.Where("ischeck = ?", 1).Where("status = ?", 1)
	}
	// 已下架off
	if list.Tab == "off" {
		db = db.Where("status = ?", 0)
	}
	// 库存预警
	if list.Tab == "min_stock" {
		db = db.Where("stock <= min_stock")
	}
	// 回收站
	if list.Tab == "delete" {
		db = db.Where("delete_time != ?", 0)
	}
	err := db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Category").Preload("GoodsBanner").Preload("GoodsSkus").Preload("GoodsSkusCard").Find(&m).Error

	return m, total, err
}

func GoodsCheck(r request.GoodsCheck) (res int64, err error) {
	db := global.GVA_DB.Table("goods")
	result := db.Where("id  = ?", r.Id).Updates(&r)
	return result.RowsAffected, result.Error
}

func GetGoodsId(r request.GetById) (res response.GoodsRead, err error) {
	db := global.GVA_DB.Table("goods")
	err = db.Where("id = ?", r.Id).First(&res).Error
	return
}

func GoodsUpdateSkus(r request.GoodsUpdateSkus, SkuValue string) (res int64, err error) {
	db := global.GVA_DB.Table("goods")
	result := db.Where("id = ?", r.Id).Update("sku_value", SkuValue)
	return result.RowsAffected, result.Error
}

func GoodsUpdateSkuType(r request.GoodsUpdateSkus) (res int64, err error) {
	db := global.GVA_DB.Table("goods")
	result := db.Where("id  = ?", r.Id).Update("sku_type", r.SkuType)
	return result.RowsAffected, result.Error
}

func GetOrderGoodsInfo(Id int64) (res response.GoodsOrder, err error) {
	db := global.GVA_DB.Table("goods")
	err = db.Where("id = ?", Id).Find(&res).Error
	return
}

//主控台

func GetChecking() (total int64) {
	db := global.GVA_DB.Table("goods")
	_ = db.Where("ischeck = ?", 0).Where("delete_time = ?", 0).Count(&total).Error
	return
}
func GetSaling() (total int64) {
	db := global.GVA_DB.Table("goods")
	_ = db.Where("ischeck = ?", 1).Where("status = ?", 1).Count(&total).Error
	return
}

func GetOff() (total int64) {
	db := global.GVA_DB.Table("goods")
	_ = db.Where("status = ?", 0).Count(&total).Error
	return
}

func GetMinStock() (total int64) {
	db := global.GVA_DB.Table("goods")
	_ = db.Where("status = ?", 0).Where("stock <= min_stock").Count(&total).Error
	return
}
