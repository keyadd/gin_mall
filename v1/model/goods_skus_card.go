package model

import (
	"gin_mall/global"
	"gin_mall/v1/model/request"
	"gin_mall/v1/model/response"
)

func GetGoodsSkusCardId(goodsId int64) (res []response.GoodsSkusCard, err error) {
	db := global.GVA_DB.Table("goods_skus_card")
	err = db.Where("goods_id = ?", goodsId).Find(&res).Error
	return
}

func CreateGoodsSkusCard(GoodsSkusCard response.GoodsSkusCard) (response.GoodsSkusCard, error) {
	db := global.GVA_DB.Table("goods_skus_card")
	err := db.Create(&GoodsSkusCard).Error
	return GoodsSkusCard, err

}

func EditGoodsSkusCard(GoodsSkusCard request.EditGoodsSkusCardRequest) (res int64, err error) {
	db := global.GVA_DB.Table("goods_skus_card")
	result := db.Where("id = ?", GoodsSkusCard.Id).Updates(GoodsSkusCard.GoodsSkusCardRequest)
	return result.RowsAffected, result.Error

}

func DeleteGoodsSkusCardId(r request.GetById) (res int64, err error) {
	db := global.GVA_DB.Table("goods_skus_card")
	result := db.Delete(&r)
	return result.RowsAffected, result.Error
}

func GoodsSkusCardUpdateOrder(r request.UpdateStatusRequest) (res int64, err error) {
	db := global.GVA_DB.Table("goods_skus_card")
	result := db.Where("id = ?", r.Id).Updates(&r)
	return result.RowsAffected, result.Error
}
