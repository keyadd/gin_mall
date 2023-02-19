package service

import (
	"gin_mall/global"
	"gin_mall/utils"
	"gin_mall/v1/model"
	"gin_mall/v1/model/request"
	"gin_mall/v1/model/response"
	"go.uber.org/zap"
	"strconv"
	"time"
)

// GetUserList 获取User 用户列表
func GetUserList(r request.UserListRequest) (res response.UserListRes) {
	UserList, err := model.GetUserList(r)
	if err != nil {
		global.GVA_LOG.Error("model.GetUserList(list)", zap.Error(err))
		return
	}

	var userLevelIds []int64
	for _, v := range UserList {
		userLevelIds = append(userLevelIds, v.UserLevelId)
	}

	levelIds := model.GetUserLevelIds(userLevelIds)
	var arr []*response.User
	for _, i := range UserList {
		for _, j := range levelIds {

			if i.UserLevelId == j.Id {
				i.UserLevel = j
				createTime, _ := strconv.Atoi(i.CreateTime)
				unix := time.Unix(int64(createTime), 0)
				i.CreateTime = unix.Format("2006-01-02 15:04:05")
				arr = append(arr, i)
			}

		}
	}
	level, err := model.GetUserLevel()
	if err != nil {
		global.GVA_LOG.Error("model.GetUserList(list)", zap.Error(err))
		return
	}

	total := model.GetUserCount()
	result := response.UserListRes{
		UserList:  arr,
		Total:     total,
		UserLevel: level,
	}
	return result
}

// CreateUser 创建User用户
func CreateUser(r request.UserRequest) (res response.CreateUser) {
	password := utils.MD5V([]byte(r.Password))
	m := response.CreateUser{
		Username:    r.Username,
		CreateTime:  time.Now().Unix(),
		Status:      r.Status,
		Password:    password,
		Phone:       r.Phone,
		Email:       r.Email,
		Avatar:      r.Avatar,
		UserLevelId: r.UserLevelId,
	}

	u, err := model.CreateUser(m)

	if err != nil {
		global.GVA_LOG.Error("model.CreateUser(m)", zap.Any("m", m), zap.Error(err))
		return
	}
	return u
}

func EditUser(r request.UserEditRequest) bool {
	r.Password = utils.MD5V([]byte(r.Password))
	res, err := model.EditUser(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditUser(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}
}

func DeleteUser(r request.GetById) bool {
	res, err := model.DeleteUserId(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditUser(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}

}

func UpdateStatusUser(r request.UpdateStatusRequest) bool {
	res, err := model.UpdateStatusUserId(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditUser(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}
}
