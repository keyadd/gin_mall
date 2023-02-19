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

// GetSkusList 获取Skus 用户列表
func GetSkusList(list request.PageInfo) (res response.PageResult) {
	SkusList, total, err := model.GetSkusList(list)
	if err != nil {
		global.GVA_LOG.Error("model.GetSkusList(list)", zap.Any("list", list), zap.Error(err))
		return
	}
	var arr []*response.SkusList
	for _, j := range SkusList {

		UpdateTime, _ := strconv.Atoi(j.UpdateTime)
		unixUpdateTime := time.Unix(int64(UpdateTime), 0)
		j.UpdateTime = unixUpdateTime.Format("2006-01-02 15:04:05")

		createTime, _ := strconv.Atoi(j.CreateTime)
		unixCreateTime := time.Unix(int64(createTime), 0)
		j.CreateTime = unixCreateTime.Format("2006-01-02 15:04:05")
		arr = append(arr, j)

	}
	result := response.PageResult{
		List:  SkusList,
		Total: total,
	}
	return result
}

// CreateSkus 创建Skus用户
func CreateSkus(r request.AddSkus) (res response.Skus) {
	skus := response.Skus{
		Name:    r.Name,
		Order:   r.Order,
		Status:  r.Status,
		Default: r.Default,
	}

	u, err := model.CreateSkus(skus)

	if err != nil {
		global.GVA_LOG.Error("model.CreateSkus(skus)", zap.Any("skus", skus), zap.Error(err))
		return
	}
	return u
}

func EditSkus(r request.EditSkus) bool {

	res, err := model.EditSkus(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditSkus(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}
}

func DeleteSkus(r request.IdsReq) bool {
	var arr []int
	for _, v := range r.Ids {
		arr = append(arr, v)
	}
	res, err := model.DeleteSkusId(arr)
	if err != nil {
		global.GVA_LOG.Error("model.EditSkus(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}

}

func UpdateStatusSkus(r request.UpdateStatusRequest) bool {
	res, err := model.UpdateStatusSkusId(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditSkus(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}
}
