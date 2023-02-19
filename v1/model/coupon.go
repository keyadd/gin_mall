package model

import (
	"gin_mall/global"
	"gin_mall/v1/model/request"
	"gin_mall/v1/model/response"
)

func CreateCoupon(s response.Coupon) (response.Coupon, error) {
	db := global.GVA_DB.Table("coupon")
	err := db.Create(&s).Error
	return s, err
}

func EditCoupon(r request.EditCoupon) (res int64, err error) {
	db := global.GVA_DB.Table("coupon")
	result := db.Where("id = ?", r.Id).Updates(r.AddCoupon)
	return result.RowsAffected, result.Error

}

func DeleteCouponId(arr []int) (res int64, err error) {
	db := global.GVA_DB.Table("coupon")
	var r response.Coupon
	result := db.Delete(r, arr)
	return result.RowsAffected, result.Error
}

func UpdateStatusCouponId(r request.UpdateStatusRequest) (res int64, err error) {
	db := global.GVA_DB.Table("coupon")
	result := db.Where("id = ?", r.Id).Updates(&r)
	return result.RowsAffected, result.Error
}

func GetCouponList(list request.PageInfo) ([]*response.CouponList, error) {
	var m []*response.CouponList
	limit := list.PageSize
	offset := list.PageSize * (list.Page - 1)

	db := global.GVA_DB.Table("coupon")
	err := db.Limit(limit).Offset(offset).Find(&m).Error
	return m, err
}
