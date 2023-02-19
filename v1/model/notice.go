package model

import (
	"gin_mall/global"
	"gin_mall/v1/model/request"
	"gin_mall/v1/model/response"
)

func GetNoticeList(page request.NoticeList) (res []*response.NoticeList, total int64, err error) {
	limit := page.PageSize
	offset := page.PageSize * (page.Page - 1)
	db := global.GVA_DB.Table("notice")
	_ = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&res).Error
	return res, total, err
}

func GetNoticeCount() (total int64) {
	db := global.GVA_DB.Table("notice")
	_ = db.Count(&total).Error
	return
}

func AddNotice(r response.Notice) (res response.Notice, err error) {
	db := global.GVA_DB.Table("notice")
	err = db.Create(&r).Error
	return r, err
}

func EditNotice(r request.NoticeEdit) (res int64, err error) {
	db := global.GVA_DB.Table("notice")
	result := db.Where("id = ?", r.Id).Updates(r)
	return result.RowsAffected, result.Error
}

func DeleteNoticeId(r request.GetById) (res int64, err error) {
	db := global.GVA_DB.Table("notice")
	result := db.Delete(&r)
	return result.RowsAffected, result.Error
}
