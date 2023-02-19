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

func CreateOperationRecord(record response.OperationRecord) (err error) {
	err = model.CreateOperationRecord(record)
	return
}

func GetRuleNameOperationRecord(path string, method string) (name string) {
	ruleName, err := model.GetRuleName(path, method)
	if err != nil {
		global.GVA_LOG.Error("model.GetRuleName(path, method))", zap.Any("path", path), zap.Error(err))
		return
	}
	return ruleName
}

func GetOperationRecordList(r request.OperationRecord) (result response.PageResult) {

	res, total, err := model.OperationRecordList(r)
	if err != nil {
		global.GVA_LOG.Error("model.OperationRecordList(r)", zap.Any("r", r), zap.Error(err))
		return
	}

	var arr []response.OperationRecordList
	for _, j := range res {

		UpdateTime, _ := strconv.Atoi(j.UpdateTime)
		unixUpdateTime := time.Unix(int64(UpdateTime), 0)
		j.UpdateTime = unixUpdateTime.Format("2006-01-02 15:04:05")

		createTime, _ := strconv.Atoi(j.CreateTime)
		unixCreateTime := time.Unix(int64(createTime), 0)
		j.CreateTime = unixCreateTime.Format("2006-01-02 15:04:05")
		arr = append(arr, j)

	}

	result = response.PageResult{
		List:  arr,
		Total: total,
	}
	return result
}

func DeleteOperationRecord(r request.GetById) bool {
	list := response.OperationRecord{
		DeleteTime: time.Now().Unix(),
	}

	res, err := model.DeleteOperationRecordId(r, list)
	if err != nil {
		global.GVA_LOG.Error("model.DeleteNoticeId(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}

}
