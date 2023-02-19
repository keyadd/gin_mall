package model

import (
	"gin_mall/global"
	"gin_mall/v1/model/response"
)

func GetBannerId(goodsId int64) (res []response.GoodsBanner, err error) {
	db := global.GVA_DB.Table("goods_banner")
	err = db.Where("goods_id = ?", goodsId).Find(&res).Error
	return
}

func GetBannerIdsList(goodsId []int64) (res []response.GoodsBanner, err error) {
	db := global.GVA_DB.Table("goods_banner")
	err = db.Where("goods_id in ?", goodsId).Find(&res).Error
	return
}

func GoodsBannerSet(arr []response.GoodsBanner) (res []response.GoodsBanner, err error) {
	db := global.GVA_DB.Table("goods_banner")
	err = db.Create(&arr).Error
	return arr, err
}
