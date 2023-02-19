package service

import (
	"fmt"
	"gin_mall/global"
	"gin_mall/utils"
	"gin_mall/v1/model"
	"gin_mall/v1/model/request"
	"gin_mall/v1/model/response"
	"go.uber.org/zap"
	"strconv"
	"time"
)

// GetCategoryList 获取Category 用户列表
func GetCategoryList() (res response.PageResult) {
	CategoryList, err := model.GetCategoryList()
	if err != nil {
		global.GVA_LOG.Error("model.GetCategoryList(list)", zap.Error(err))
		return
	}
	var arr []*response.CategoryList
	for _, j := range CategoryList {

		UpdateTime, _ := strconv.Atoi(j.UpdateTime)
		unixUpdateTime := time.Unix(int64(UpdateTime), 0)
		j.UpdateTime = unixUpdateTime.Format("2006-01-02 15:04:05")

		createTime, _ := strconv.Atoi(j.CreateTime)
		unixCreateTime := time.Unix(int64(createTime), 0)
		j.CreateTime = unixCreateTime.Format("2006-01-02 15:04:05")
		arr = append(arr, j)

	}

	categoryTree := utils.CategoryTree(arr, 0)
	result := response.PageResult{
		List: categoryTree,
	}

	return result
}

// CreateCategory 创建Category用户
func CreateCategory(r request.CategoryRequest) (res response.Category) {

	m := response.Category{
		Name:       r.Name,
		CreateTime: time.Now().Unix(),
		CategoryId: r.CategoryId,
		Status:     r.Status,
		Order:      r.Order,
	}
	fmt.Println(m)

	u, err := model.CreateCategory(m)

	if err != nil {
		global.GVA_LOG.Error("model.CreateCategory(m)", zap.Any("m", m), zap.Error(err))
		return
	}
	fmt.Println(u)
	return u
}

func EditCategory(r request.CategoryEditRequest) bool {

	res, err := model.EditCategory(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditCategory(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}
}

func DeleteCategory(r request.GetById) bool {
	res, err := model.DeleteCategoryId(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditCategory(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}

}

func UpdateStatusCategory(r request.UpdateStatusRequest) bool {
	res, err := model.UpdateStatusCategoryId(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditCategory(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}
}
