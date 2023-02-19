package service

import (
	"errors"
	"gin_mall/global"
	"gin_mall/initialize"
	"gin_mall/utils"
	"gin_mall/v1/model"
	"gin_mall/v1/model/request"
	"gin_mall/v1/model/response"
	"go.uber.org/zap"
	"strconv"
	"time"
)

var (
	ErrorUserExist        = errors.New("用户已存在")
	ErrorUserNotExist     = errors.New("not is username")
	ErrorPasswordNotExist = errors.New("not is password")
)

func Login(p *request.Login) (string, error) {
	oPassword := p.Password
	m, err := model.GetManger(p)

	if err != nil {
		return "", err
	}
	password := utils.MD5V([]byte(oPassword))

	if m.Password != password {
		return "", ErrorPasswordNotExist
	}

	//生成token
	return initialize.GenToken(m.Id, m.Username)

}

func GetManagerInfo(id int64) (res response.ManagerInfo, err error) {

	info, err := model.GetManagerInfo(id)
	if err != nil {
		global.GVA_LOG.Error("model.model.GetManagerRoleId(id) failed", zap.Int64("manager_id", id), zap.Error(err))
		return
	}
	ruleIdList, err := model.GetRoleRuleList(info.RoleId)
	if err != nil {
		global.GVA_LOG.Error("model.GetManagerRuleID(info.RoleId) failed", zap.Int64("info.RoleId", info.RoleId), zap.Error(err))
		return
	}
	var ruleIds []int64
	for _, v := range ruleIdList {
		ruleIds = append(ruleIds, v.RuleId)
	}

	menu, err := model.GetRuleMenu(ruleIds)
	if err != nil {
		global.GVA_LOG.Error("model.GetRuleMenu(rule_id) failed", zap.Any("rule_id[]", ruleIds), zap.Error(err))
		return
	}

	menus := utils.MenuTree(menu, 0)
	managerInfo := response.ManagerInfo{
		Id:       info.Id,
		Username: info.Username,
		Avatar:   info.Avatar,
		Super:    info.Super,
		Menus:    menus,
	}

	return managerInfo, err
}

// GetManagerList 获取manager 用户列表
func GetManagerList(list request.ManagerList) (res response.ResList, err error) {
	managerList, total, err := model.GetManagerList(list)
	if err != nil {
		global.GVA_LOG.Error("model.GetManagerList(list)", zap.Any("list", list), zap.Error(err))
		return
	}
	var listArr []*response.ManagerList
	//var err error
	for _, v := range managerList {
		//var role *response.InfoRole
		//role, err = model.InfoRole(v.RoleId)
		//if err != nil {
		//	global.GVA_LOG.Error("model.InfoRole(v.RoleId))", zap.Any("v.RoleId", v.RoleId), zap.Error(err))
		//	return
		//}

		UpdateTime, _ := strconv.Atoi(v.UpdateTime)
		unixUpdateTime := time.Unix(int64(UpdateTime), 0)
		v.UpdateTime = unixUpdateTime.Format("2006-01-02 15:04:05")

		createTime, _ := strconv.Atoi(v.CreateTime)
		unixCreateTime := time.Unix(int64(createTime), 0)
		v.CreateTime = unixCreateTime.Format("2006-01-02 15:04:05")
		r := response.ManagerList{
			Id:         v.Id,
			Username:   v.Username,
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
			Status:     v.Status,
			Avatar:     v.Avatar,
			RoleId:     v.RoleId,
			Super:      v.Super,
			Role:       v.Role,
		}
		listArr = append(listArr, &r)
	}

	roleList, err := model.RoleList()
	if err != nil {
		global.GVA_LOG.Error("model.RoleList()", zap.Error(err))
		return
	}

	resList := response.ResList{
		ManagerList: listArr,
		Total:       total,
		InfoRole:    roleList,
	}
	return resList, err
}

// CreateManager 创建manager用户
func CreateManager(r request.ManagerRequest) (res response.CreateManager, err error) {
	password := utils.MD5V([]byte(r.Password))
	m := response.Manager{
		Status:     r.Status,
		CreateTime: time.Now().Unix(),
		Username:   r.Username,
		Password:   password,
		Avatar:     r.Avatar,
		RoleId:     r.RoleId,
	}

	u, err := model.CreateManager(m)
	res = response.CreateManager{
		Id:         u.Id,
		Status:     u.Status,
		CreateTime: u.CreateTime,
		UpdateTime: u.UpdateTime,
		Username:   u.Username,
		RoleId:     u.RoleId,
		Avatar:     u.Avatar,
		Super:      u.Super,
	}

	if err != nil {
		global.GVA_LOG.Error("model.CreateManager(m)", zap.Any("m", m), zap.Error(err))
		return
	}
	return
}

func EditManager(r request.EditRequest) bool {
	r.ManagerRequest.Password = utils.MD5V([]byte(r.ManagerRequest.Password))
	res, err := model.EditManager(r)
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

func DeleteManager(r request.GetById) bool {
	res, err := model.DeleteManagerId(r)
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

func UpdateStatusManager(r request.UpdateStatusRequest) bool {
	res, err := model.UpdateStatusManagerId(r)
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

func LogoutManager(i int64) (res bool) {
	info, err := model.GetManagerInfo(i)
	if err != nil {
		global.GVA_LOG.Error("model.EditManager(r)", zap.Any("i", i), zap.Error(err))
		return false
	}

	_, err = initialize.GenToken(i, info.Username)
	if err != nil {
		global.GVA_LOG.Error("initialize.GenToken(i, info.Username)", zap.Any("i, info.Username", info.Username), zap.Error(err))
		return false
	}

	return true

}

func ChangePassword(i int64, r request.UpdatePasswordRequest) bool {
	r.OldPassword = utils.MD5V([]byte(r.OldPassword))
	r.Password = utils.MD5V([]byte(r.Password))

	res, err := model.UpdatePassword(i, r)
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
