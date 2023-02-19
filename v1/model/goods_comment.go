package model

import (
	"gin_mall/global"
	"gin_mall/v1/model/request"
	"gin_mall/v1/model/response"
)

func GetGoodsCommentCount() (total int64) {
	db := global.GVA_DB.Table("order_item")
	_ = db.Count(&total).Error
	return
}
func GetGoodsCommentList(list request.ManagerList) ([]*response.GoodsComment, int64, error) {
	var m []*response.GoodsComment

	limit := list.PageSize
	offset := list.PageSize * (list.Page - 1)
	var total int64
	db := global.GVA_DB.Table("order_item")
	db = db.Select("order_item.*,user.username,user.nickname,user.avatar,goods.title,goods.cover")
	db = db.Joins("join user on order_item.user_id = user.id")
	db = db.Joins("join goods on order_item.goods_id = goods.id")
	//gdb = db.Where("GoodsCommentname like ?", "%"+list.Keyword+"%")
	_ = db.Count(&total).Error
	err := db.Limit(limit).Offset(offset).Find(&m).Error
	return m, total, err

}

func CreateGoodsComment(m response.CreateGoodsComment) (response.CreateGoodsComment, error) {
	db := global.GVA_DB.Table("order_item")
	err := db.Create(&m).Error
	return m, err
}

func EditGoodsComment(r request.GoodsCommentEditRequest) (res int64, err error) {
	db := global.GVA_DB.Table("order_item")
	result := db.Where("id = ?", r.Id).Updates(r.GoodsCommentRequest)
	return result.RowsAffected, result.Error

}

func DeleteGoodsCommentId(r request.GetById) (res int64, err error) {
	db := global.GVA_DB.Table("order_item")
	result := db.Delete(&r)
	return result.RowsAffected, result.Error
}

func UpdateStatusGoodsCommentId(r request.UpdateStatusRequest) (res int64, err error) {
	db := global.GVA_DB.Table("order_item")
	result := db.Where("id = ?", r.Id).Updates(&r)
	return result.RowsAffected, result.Error
}

func OrderItemId(Id int64) (res []response.OrderItem, err error) {
	db := global.GVA_DB.Table("order_item")
	err = db.Where("order_id = ?", Id).Find(&res).Error
	return
}
