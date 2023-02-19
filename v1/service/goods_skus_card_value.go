package service

import (
	"gin_mall/global"
	"gin_mall/v1/model"
	"gin_mall/v1/model/request"
	"gin_mall/v1/model/response"
	"go.uber.org/zap"
)

// CreateManager 创建manager用户
func CreateGoodsSkusCardValue(r request.GoodsSkusCardValueRequest) (res response.GoodsSkusCardValue) {

	GoodsSkusCardValue := response.GoodsSkusCardValue{
		Name:            r.Name,
		Order:           r.Order,
		GoodsSkusCardId: r.GoodsSkusCardId,
		Value:           r.Value,
	}

	res, err := model.CreateGoodsSkusCardValue(GoodsSkusCardValue)

	if err != nil {
		global.GVA_LOG.Error("model.CreateGoodsSkusCardValue(GoodsSkusCardValue)", zap.Any("GoodsSkusCardValue", GoodsSkusCardValue), zap.Error(err))
		return
	}
	return res
}

func EditGoodsSkusCardValue(r request.EditGoodsSkusCardValueRequest) bool {
	res, err := model.EditGoodsSkusCardValue(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditManager(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}
}

func DeleteGoodsSkusCardValue(r request.GetById) bool {
	res, err := model.DeleteGoodsSkusCardValueId(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditManager(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}

}

func SetGoodsSkusCardValue(r request.UpdateStatusRequest) bool {
	res, err := model.SetGoodsSkusCardValue(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditManager(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}
}
