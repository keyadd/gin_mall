package service

import (
	"gin_mall/global"
	"gin_mall/v1/model"
	"gin_mall/v1/model/request"
	"gin_mall/v1/model/response"
	"go.uber.org/zap"
)

// CreateManager 创建manager用户
func CreateGoodsSkusCard(r request.GoodsSkusCardRequest) (res response.GoodsSkusCard) {

	GoodsSkusCard := response.GoodsSkusCard{
		GoodsId: r.GoodsId,
		Name:    r.Name,
		Type:    r.Type,
		Order:   r.Order,
	}

	res, err := model.CreateGoodsSkusCard(GoodsSkusCard)

	if err != nil {
		global.GVA_LOG.Error("model.CreateGoodsSkusCard(GoodsSkusCard)", zap.Any("GoodsSkusCard", GoodsSkusCard), zap.Error(err))
		return
	}
	return
}

func EditGoodsSkusCard(r request.EditGoodsSkusCardRequest) bool {
	res, err := model.EditGoodsSkusCard(r)
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

func DeleteGoodsSkusCard(r request.GetById) bool {
	res, err := model.DeleteGoodsSkusCardId(r)
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

func GoodsSkusCardUpdateOrder(r request.UpdateStatusRequest) bool {
	res, err := model.GoodsSkusCardUpdateOrder(r)
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
