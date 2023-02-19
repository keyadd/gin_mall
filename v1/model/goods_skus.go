package model

import (
	"gin_mall/global"
	"gin_mall/v1/model/response"
)

func GetGoodsSkusId(goodsId int64) (res []response.GoodsSkus, err error) {
	db := global.GVA_DB.Table("goods_skus")
	err = db.Where("goods_id = ?", goodsId).Find(&res).Error
	return
}

func GetGoodsSkusReadId(goodsId int64) (res []response.GoodsSkusRead, err error) {
	db := global.GVA_DB.Table("goods_skus")
	err = db.Where("goods_id = ?", goodsId).Find(&res).Error
	return
}

func GetGoodsSkusJson(goodsId int64) (res []response.GoodsSkusJson, err error) {
	db := global.GVA_DB.Table("goods_skus")
	err = db.Where("goods_id = ?", goodsId).Find(&res).Error
	return
}
func GoodsUpdateSkusArr(r []response.UpdateGoodsSkus) (err error) {
	db := global.GVA_DB.Table("goods_skus")
	err = db.Save(&r).Error
	return
}

func GetShopSkusId(shopId int64) (res response.GoodsSkusArr, err error) {
	db := global.GVA_DB.Table("goods_skus")
	err = db.Where("id = ?", shopId).Find(&res).Error
	return
}

func GetGoodsSkus(GoodsId int64) (res []response.GoodsSkusArr, err error) {
	db := global.GVA_DB.Table("goods_skus")
	err = db.Where("goods_id = ?", GoodsId).Find(&res).Error
	return
}
