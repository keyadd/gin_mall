package service

import (
	"gin_mall/global"
	"gin_mall/v1/model"
	"gin_mall/v1/model/request"
	"gin_mall/v1/model/response"
	"go.uber.org/zap"
	"time"
)

// GetAppCategoryItemList 获取AppCategoryItem 用户列表
func GetAppCategoryItemList(r request.AppCategoryItemRequest) (res response.PageResult) {
	AppCategoryItemList, err := model.GetAppCategoryItemList(r)
	if err != nil {
		global.GVA_LOG.Error("model.GetAppCategoryItemList(list)", zap.Error(err))
		return
	}
	result := response.PageResult{List: AppCategoryItemList}

	return result
}

// CreateAppCategoryItem 创建AppCategoryItem用户
func CreateAppCategoryItem(r request.AppCategoryItemCreateRequest) (res response.AppCategoryItem) {

	m := response.AppCategoryItem{
		//Name:       r.Name,
		CreateTime: time.Now().Unix(),
	}

	u, err := model.CreateAppCategoryItem(m)

	if err != nil {
		global.GVA_LOG.Error("model.CreateAppCategoryItem(m)", zap.Any("m", m), zap.Error(err))
		return
	}
	return u
}

func DeleteAppCategoryItem(r request.GetById) bool {
	res, err := model.DeleteAppCategoryItemId(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditAppCategoryItem(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}

}
