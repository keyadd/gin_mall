package model

import (
	"gin_mall/global"
	"gin_mall/v1/model/request"
	"gin_mall/v1/model/response"
)

func CreateGoodsSkusCardValue(GoodsSkusCardValue response.GoodsSkusCardValue) (response.GoodsSkusCardValue, error) {
	db := global.GVA_DB.Table("goods_skus_card_value")
	err := db.Create(&GoodsSkusCardValue).Error
	return GoodsSkusCardValue, err

}

func EditGoodsSkusCardValue(GoodsSkusCardValue request.EditGoodsSkusCardValueRequest) (res int64, err error) {
	db := global.GVA_DB.Table("goods_skus_card_value")
	result := db.Where("id = ?", GoodsSkusCardValue.Id).Updates(GoodsSkusCardValue.GoodsSkusCardValueRequest)
	return result.RowsAffected, result.Error

}

func DeleteGoodsSkusCardValueId(r request.GetById) (res int64, err error) {
	db := global.GVA_DB.Table("goods_skus_card_value")
	result := db.Delete(&r)
	return result.RowsAffected, result.Error
}

func SetGoodsSkusCardValue(r request.UpdateStatusRequest) (res int64, err error) {
	db := global.GVA_DB.Table("goods_skus_card_value")
	result := db.Where("id = ?", r.Id).Updates(&r)
	return result.RowsAffected, result.Error
}

func GetGoodsSkusCardValueId(goodsSkusCardValueId int64) (res []response.GoodsSkusCardValue, err error) {
	db := global.GVA_DB.Table("goods_skus_card_value")
	err = db.Where("goods_skus_card_id = ?", goodsSkusCardValueId).Find(&res).Error
	return
}
