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

func GetAgentList(r request.AgentList) (res response.PageResult) {
	list, err := model.GetUserAgentList(r)
	if err != nil {
		global.GVA_LOG.Error("model.GetUserAgentList(r)", zap.Any("r", r), zap.Error(err))
		return
	}

	var arr []*response.AgentList
	for _, j := range list {
		//info, err := model.GetUserInfo(j.Id)
		//if err != nil {
		//	global.GVA_LOG.Error("model.GetUserAgentList(r)", zap.Any("j.Id", j.Id), zap.Error(err))
		//	return
		//}
		//j.UserInfo = info

		//UpdateTime, _ := strconv.Atoi(j.UpdateTime)
		//unixUpdateTime := time.Unix(int64(UpdateTime), 0)
		//j.UpdateTime = unixUpdateTime.Format("2006-01-02 15:04:05")

		createTime, _ := strconv.Atoi(j.CreateTime)
		unixCreateTime := time.Unix(int64(createTime), 0)
		j.CreateTime = unixCreateTime.Format("2006-01-02 15:04:05")
		arr = append(arr, j)

	}
	total := model.GetUserCount()
	result := response.PageResult{List: arr, Total: total}
	return result

}

func AgentStatistics() (res response.PageResult) {
	// 总分销人数
	total := model.GetUserCount()
	// 总推广订单数
	billCount := model.GetUserBillCount()

	// 总推广订单金额
	totalOrderPrice := model.GetUserOrderPrice()

	userExtractCount := model.GetUserExtractCount()

	totalStr := strconv.FormatInt(total, 10)
	billCountStr := strconv.FormatInt(billCount, 10)
	userExtractCountStr := strconv.FormatInt(userExtractCount, 10)
	totalOrderPriceStr := strconv.FormatFloat(totalOrderPrice, 'f', 10, 32)
	statistics := []response.Statistics{
		{
			Color: "bg-blue-400",
			Label: "分销员人数(人)",
			Value: totalStr,
		},
		{
			Color: "bg-orange-400",
			Label: "订单数(单)",
			Value: billCountStr,
		},
		{
			Color: "bg-green-400",
			Label: "订单金额(元)",
			Value: totalOrderPriceStr,
		},
		{
			Color: "bg-indigo-400",
			Label: "提现次数(次)",
			Value: userExtractCountStr,
		},
	}
	result := response.PageResult{List: statistics}
	return result

}
