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

func GetNoticeList(r request.NoticeList) (list response.PageResult) {

	res, total, err := model.GetNoticeList(r)
	if err != nil {
		global.GVA_LOG.Error("model.ImageClassList(page)", zap.Any("r", r), zap.Error(err))
		return
	}

	var arr []*response.NoticeList
	for _, j := range res {

		UpdateTime, _ := strconv.Atoi(j.UpdateTime)
		unixUpdateTime := time.Unix(int64(UpdateTime), 0)
		j.UpdateTime = unixUpdateTime.Format("2006-01-02 15:04:05")

		createTime, _ := strconv.Atoi(j.CreateTime)
		unixCreateTime := time.Unix(int64(createTime), 0)
		j.CreateTime = unixCreateTime.Format("2006-01-02 15:04:05")
		arr = append(arr, j)

	}

	result := response.PageResult{
		List:  arr,
		Total: total,
	}
	return result

}

func AddNotice(r request.NoticeRequest) (res response.Notice, err error) {
	notice := response.Notice{
		Title:   r.Title,
		Content: r.Content,
	}
	res, err = model.AddNotice(notice)
	if err != nil {
		global.GVA_LOG.Error("model.AddNotice(notice)", zap.Any("notice", notice), zap.Error(err))
		return
	}
	return

}

func EditNotice(r request.NoticeEdit) bool {
	res, err := model.EditNotice(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditNotice(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}
}

func DeleteNotice(r request.GetById) bool {

	res, err := model.DeleteNoticeId(r)
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
