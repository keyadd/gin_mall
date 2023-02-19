package service

import (
	"encoding/json"
	"gin_mall/global"
	"gin_mall/v1/model"
	"gin_mall/v1/model/request"
	"gin_mall/v1/model/response"
	"go.uber.org/zap"
	"time"
)

// GetOrderList 获取Order 用户列表
func GetOrderList(r request.OrderListRequest) (res response.PageResult) {

	OrderList, total, err := model.GetOrderList(r)
	if err != nil {
		global.GVA_LOG.Error("model.GetOrderList(list)", zap.Error(err))
		return
	}
	var list []*response.OrderList

	for _, v := range OrderList {

		//fmt.Println(addOrder.Id)

		var order []response.OrderItem

		var skus []response.SkusInfo
		order, err = model.OrderItemId(v.Id)
		if err != nil {
			global.GVA_LOG.Error("model.GetOrderList(list)", zap.Error(err))
			//return
		}
		//var ids []int64
		var orderArr []response.OrderItem
		for _, i := range order {

			info, err := model.GetOrderGoodsInfo(i.GoodsId)
			if err != nil {
				global.GVA_LOG.Error("model.GetOrderList(list)", zap.Error(err))
			}
			i.GoodsItem = info

			var goodsSkus response.GoodsSkusArr

			goodsSkus, err = model.GetShopSkusId(i.ShopId)
			if err != nil {
				global.GVA_LOG.Error("model.GetOrderList(list)", zap.Error(err))
				//return
			}

			_ = json.Unmarshal([]byte(goodsSkus.Skus), &skus)

			i.GoodsSkus = skus

			orderArr = append(orderArr, i)

		}
		//fmt.Printf("%v\n", order)

		v.OrderItems = orderArr
		list = append(list, v)

	}

	//total := model.GetOrderCount()
	result := response.PageResult{
		List:  list,
		Total: total,
	}
	return result
}

func DeleteOrder(r request.IdsReq) bool {
	res, err := model.DeleteOrderId(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditOrder(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}

}

func OrderSend(r request.OrderSend) bool {
	send := response.OrderSend{
		ExpressCompany: r.ExpressCompany,
		ExpressNo:      r.ExpressNo,
		ExpressTime:    time.Now().Format("2006-01-02 3:4:5"),
	}
	marshal, _ := json.Marshal(send)
	res, err := model.OrderSend(r.Id, string(marshal))
	if err != nil {
		global.GVA_LOG.Error("model.EditOrder(r)", zap.Any("r", r), zap.Error(err))
		return false
	}

	if res == 1 {
		return true
	} else {
		return false
	}
}

func OrderHandleRefund(r request.OrderHandleRefund) bool {
	//send := response.OrderSend{
	//	ExpressCompany: r.ExpressCompany,
	//	ExpressNo:      r.ExpressNo,
	//	ExpressTime:    time.Now().Format("2006-01-02 3:4:5"),
	//}
	//marshal, _ := json.Marshal(send)
	//res, err := model.OrderSend(r.Id, string(marshal))
	//if err != nil {
	//	global.GVA_LOG.Error("model.EditOrder(r)", zap.Any("r", r), zap.Error(err))
	//	return false
	//}
	res := 1

	if res == 1 {
		return true
	} else {
		return false
	}
}
