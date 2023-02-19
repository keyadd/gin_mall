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

// GetManagerList 获取manager 用户列表
func GetApiList(r request.ApiRequest) (res response.PageResult) {
	Api, total, err := model.GetRuleApiList(r)
	if err != nil {
		global.GVA_LOG.Error("model.GetRuleApi(rule_id) failed", zap.Error(err))
		return
	}

	var arr []*response.ApiList
	for _, j := range Api {

		UpdateTime, _ := strconv.Atoi(j.UpdateTime)
		unixUpdateTime := time.Unix(int64(UpdateTime), 0)
		j.UpdateTime = unixUpdateTime.Format("2006-01-02 15:04:05")

		createTime, _ := strconv.Atoi(j.CreateTime)
		unixCreateTime := time.Unix(int64(createTime), 0)
		j.CreateTime = unixCreateTime.Format("2006-01-02 15:04:05")
		arr = append(arr, j)

	}

	apiRes := response.PageResult{
		List:  arr,
		Total: total,
	}
	return apiRes
}

func GetApiListTree() (res response.PageResult) {

	api, err := model.GetApiGroup()
	if err != nil {
		global.GVA_LOG.Error("model.GetRuleApi(rule_id) failed", zap.Error(err))
		return
	}
	//var arr []string
	//
	//for _, v := range api {
	//
	//	arr = append(arr, v.ApiGroup)
	//}
	//duplication := utils.RemoveDuplication(arr)
	//
	//var apiTreeResponse []response.ApiResponse
	//for i, j := range duplication {
	//	//id, err := model.GetRuleId(j)
	//	//if err != nil {
	//	//	global.GVA_LOG.Error("model.GetRuleApi(rule_id) failed", zap.Error(err))
	//	//	continue
	//	//}
	//
	//	menu, err := model.GetApiListTree(j)
	//	if err != nil {
	//		global.GVA_LOG.Error("model.GetRuleApi(rule_id) failed", zap.Error(err))
	//		continue
	//	}
	//	apiResponse := response.ApiResponse{
	//		Id:       (i + 1) - (i+1)*2, //取一个正数的负数 0除外
	//		Name:     j,
	//		Children: menu,
	//	}
	//	apiTreeResponse = append(apiTreeResponse, apiResponse)
	//}
	apiRes := response.PageResult{
		List: api,
	}

	return apiRes

}

// CreateManager 创建manager用户
func CreateApi(r request.CreateApiRequest) (res response.Api) {

	Api := response.Api{
		Status:     r.Status,
		CreateTime: time.Now().Unix(),
		RuleId:     r.RuleId,
		Name:       r.Name,
		RouterPath: r.RouterPath,
		Menu:       r.Menu,
		Order:      r.Order,
		Method:     r.Method,
		ApiGroup:   r.ApiGroup,
	}

	res, err := model.CreateRuleApi(Api)

	if err != nil {
		global.GVA_LOG.Error("model.CreateRule(rule)", zap.Any("rule", Api), zap.Error(err))
		return
	}
	return
}

func EditApi(r request.EditApiRequest) bool {
	res, err := model.EditRuleApi(r)
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

func DeleteApi(r request.GetById) bool {
	res, err := model.DeleteRuleId(r)
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

func UpdateStatusApi(r request.UpdateStatusRequest) bool {
	res, err := model.UpdateStatusRuleId(r)
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
