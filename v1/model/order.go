package model

import (
	"gin_mall/global"
	"gin_mall/v1/model/request"
	"gin_mall/v1/model/response"
)

type Order struct{}

func (Order) TableName() string {
	return "order"
}

func UpdateOrder(r response.OrderAddressUpdate) (res int64, err error) {
	db := global.GVA_DB.Table("order_address")
	result := db.Where("id = ?", r.Id).Updates(r)
	return result.RowsAffected, result.Error
}

func GetOrderAddress() (res []response.OrderAddressUpdate, err error) {
	db := global.GVA_DB.Table("order_address")
	err = db.Find(&res).Error
	return res, err
}

func GetOrderCount() (total int64) {
	db := global.GVA_DB.Table("order")
	_ = db.Count(&total).Error
	return
}
func GetOrderList(list request.OrderListRequest) ([]*response.OrderList, int64, error) {
	var m []*response.OrderList
	//var u response.UserOrder
	var total int64
	limit := list.PageSize
	offset := list.PageSize * (list.Page - 1)

	db := global.GVA_DB.Table("order")
	db = db.Select("order.id,order.no,order.user_id,order.order_address_id,order.total_price,order.remark,order.paid_time,order.payment_method,order.payment_no," +
		"order.refund_status,order.refund_no,order.closed,order.ship_status,order.ship_data,order.extra,order.create_time,order.update_time,order.reviewed,order.coupon_user_id," +
		"user.username,user.nickname,order_address.name,order_address.phone")
	db = db.Joins("left join user on order.user_id = user.id")
	db = db.Joins("join order_address on order.order_address_id = order_address.id")
	if list.No != "" {
		db = db.Where("order.no = ?", list.No)
	}
	if list.Phone != "" {
		db = db.Where("order_address.phone = ?", list.Phone)
	}
	if list.Name != "" {
		db = db.Where("order_address.name = ?", list.Name)
	}
	// 待付款
	if list.Tab == "nopay" {
		db = db.Where("order.closed = ?", 0).Where("order.payment_method = ", "")
	}
	// 待发货
	if list.Tab == "noship" {
		db = db.Where("order.closed = ?", 0).Where("order.payment_method != ", "").Where("order.ship_status = ?", "pending").Where("order.refund_status = ?", "pending")
	}
	// 已发货
	if list.Tab == "shiped" {
		db = db.Where("order.closed = ?", 0).Where("order.payment_method != ", "").Where("order.ship_status = ?", "delivered").Where("order.refund_status = ?", "pending")
	}
	// 已收货
	if list.Tab == "received" {
		db = db.Where("order.closed = ?", 0).Where("order.payment_method != ", "").Where("order.ship_status = ?", "received").Where("order.refund_status = ?", "pending")
	}
	// 已完成
	if list.Tab == "finish" {
		db = db.Where("order.closed = ?", 0).Where("order.payment_method != ", "").Where("order.ship_status = ?", "received").Where("order.refund_status = ?", "pending")
	}
	// 已关闭
	if list.Tab == "closed" {
		db = db.Where("order.closed = ?", 1)
	}
	// 退款中
	if list.Tab == "refunding" {
		db = db.Where("order.closed = ?", 0).Where("order.refund_status = ?", "applied")
	}

	_ = db.Count(&total).Error

	err := db.Limit(limit).Offset(offset).Find(&m).Error
	return m, total, err

}

func GetTotalSale() (res float64) {

	db := global.GVA_DB.Table("order")
	_ = db.Where("payment_method !=?", " ").Pluck("sum(total_price) as total", &res)
	return

}

func GetOrderPayCount() (total int64) {
	db := global.GVA_DB.Table("order")
	_ = db.Where("payment_method !=?", " ").Count(&total).Error
	return
}

func DeleteOrderId(r request.IdsReq) (res int64, err error) {
	db := global.GVA_DB.Model(&Order{})
	result := db.Delete(&r)
	return result.RowsAffected, result.Error
}

func OrderSend(id int64, s string) (res int64, err error) {
	db := global.GVA_DB.Table("order")
	result := db.Where("id = ?", id).Select("ShipData").Updates(response.OrderList{ShipData: s})
	return result.RowsAffected, result.Error
}

//主控台

func GetOrderNopayCount() (total int64) {
	db := global.GVA_DB.Table("order")
	_ = db.Where("closed =?", 0).Where("payment_method =?", " ").Count(&total).Error
	return
}

func GetOrderNoshipCount() (total int64) {
	db := global.GVA_DB.Table("order")
	_ = db.Where("closed =?", 0).Where("payment_method !=?", " ").Where("ship_status =?", "pending").Where("refund_status =?", "pending").Count(&total).Error
	return
}

func GetOrderShipedCount() (total int64) {
	db := global.GVA_DB.Table("order")
	_ = db.Where("closed =?", 0).Where("payment_method !=?", " ").Where("ship_status =?", "delivered").Where("refund_status =?", "pending").Count(&total).Error
	return
}

func GetOrderRefundingCount() (total int64) {
	db := global.GVA_DB.Table("order")
	_ = db.Where("closed =?", 0).Where("refund_status =?", "applied").Count(&total).Error
	return
}

func GetOrder(startTime int64, endTime int64, dateStr string) (res []response.Res, err error) {
	var u []response.Res
	db := global.GVA_DB
	Query1 := db.Model(&Order{}).Where("create_time>= ?", startTime).Where("create_time<= ?", endTime)
	err = db.Select("FROM_UNIXTIME(a.create_time,?) as time,COUNT(a.id) AS num", dateStr).Table("(?) as a", Query1).Group("time").Find(&u).Error

	//fmt.Println(u)
	return u, err
}
