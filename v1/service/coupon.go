package service

import (
	"gin_mall/global"
	"gin_mall/v1/model"
	"gin_mall/v1/model/request"
	"gin_mall/v1/model/response"
	"go.uber.org/zap"
	"strconv"
	"time"
)

// GetCouponList 获取Coupon 用户列表
func GetCouponList(list request.PageInfo) (res response.PageResult) {
	CouponList, err := model.GetCouponList(list)
	if err != nil {
		global.GVA_LOG.Error("model.GetCouponList(list)", zap.Any("list", list), zap.Error(err))
		return
	}

	var arr []*response.CouponList
	for _, j := range CouponList {

		UpdateTime, _ := strconv.Atoi(j.UpdateTime)
		unixUpdateTime := time.Unix(int64(UpdateTime), 0)
		j.UpdateTime = unixUpdateTime.Format("2006-01-02 15:04:05")

		createTime, _ := strconv.Atoi(j.CreateTime)
		unixCreateTime := time.Unix(int64(createTime), 0)
		j.CreateTime = unixCreateTime.Format("2006-01-02 15:04:05")
		arr = append(arr, j)

	}

	result := response.PageResult{List: arr}
	return result
}

// CreateCoupon 创建Coupon用户
func CreateCoupon(r request.AddCoupon) (res response.Coupon) {
	coupon := response.Coupon{
		Name:       r.Name,
		Order:      r.Order,
		CreateTime: time.Now().Unix(),
		Type:       r.Type,
		Value:      r.Value,
		Total:      r.Total,
		MinPrice:   r.MinPrice,
		StartTime:  r.StartTime,
		EndTime:    r.EndTime,
	}

	u, err := model.CreateCoupon(coupon)

	if err != nil {
		global.GVA_LOG.Error("model.CreateCoupon(Coupon)", zap.Any("Coupon", coupon), zap.Error(err))
		return
	}
	return u
}

func EditCoupon(r request.EditCoupon) bool {
	r.UpdateTime = time.Now().Unix()

	res, err := model.EditCoupon(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditCoupon(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}
}

func DeleteCoupon(r request.IdsReq) bool {
	var arr []int
	for _, v := range r.Ids {
		arr = append(arr, v)
	}
	res, err := model.DeleteCouponId(arr)
	if err != nil {
		global.GVA_LOG.Error("model.EditCoupon(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}

}

func UpdateStatusCoupon(r request.UpdateStatusRequest) bool {
	res, err := model.UpdateStatusCouponId(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditCoupon(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}
}
