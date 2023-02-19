package service

import (
	"gin_mall/global"
	"gin_mall/v1/model"
	"gin_mall/v1/model/request"
	"gin_mall/v1/model/response"
	"go.uber.org/zap"
)

func GetImageClassList(r request.ImageClassList) (list response.PageResult) {

	res, err := model.ImageClassList(r)
	if err != nil {
		global.GVA_LOG.Error("model.ImageClassList(page)", zap.Any("r", r), zap.Error(err))
		return
	}
	var resList []*response.ImageClassRes
	for _, v := range res {
		count := model.GetImageCount(v.Id)
		classRes := response.ImageClassRes{
			Id:          v.Id,
			Name:        v.Name,
			Order:       v.Order,
			ImagesCount: count,
		}

		resList = append(resList, &classRes)
	}
	total := model.GetImageClassCount()
	result := response.PageResult{
		List:  resList,
		Total: total,
	}
	return result

}

func GetIdImageList(r request.ImageClassId) (result response.PageResult, err error) {
	count := model.GetImageCount(r.Id)
	res, err := model.GetImageList(r)
	if err != nil {
		global.GVA_LOG.Error("model.ImageClassList(page)", zap.Any("r", r), zap.Error(err))
		return
	}
	result = response.PageResult{
		List:  res,
		Total: count,
	}
	//fmt.Println(res)
	return
}

func AddImageClass(r request.CreateImageClass) (res response.ImageClass, err error) {

	imageClass := response.ImageClass{
		Name:  r.Name,
		Order: r.Order,
	}
	res, err = model.AddImageClass(imageClass)
	if err != nil {
		global.GVA_LOG.Error("model.AddImageClass(r)", zap.Any("r", r), zap.Error(err))
		return
	}
	return

}

func EditImageClass(r request.EditImageClass) bool {
	res, err := model.EditImageClass(r)
	if err != nil {
		global.GVA_LOG.Error("model.AddImageClass(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}
}

func DeleteImageClass(r request.GetById) bool {

	res, err := model.DeleteImageClassId(r)
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
