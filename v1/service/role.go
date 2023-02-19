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

// GetRoleList 获取Role 用户列表
func GetRoleList(list request.PageInfo) (res response.PageResult) {
	RoleList, total, err := model.GetRoleList(list)
	if err != nil {
		global.GVA_LOG.Error("model.GetRoleList(list)", zap.Any("list", list), zap.Error(err))
		return
	}
	var result []*response.Role
	for _, j := range RoleList {
		UpdateTime, _ := strconv.Atoi(j.UpdateTime)
		unixUpdateTime := time.Unix(int64(UpdateTime), 0)
		j.UpdateTime = unixUpdateTime.Format("2006-01-02 15:04:05")

		createTime, _ := strconv.Atoi(j.CreateTime)
		unixCreateTime := time.Unix(int64(createTime), 0)
		j.CreateTime = unixCreateTime.Format("2006-01-02 15:04:05")

		result = append(result, j)

	}
	//total := model.GetRoleCount()
	pageResult := response.PageResult{
		List:  result,
		Total: total,
	}
	return pageResult
}

// CreateRole 创建Role用户
func CreateRole(r request.AddRole) (res response.CreateRole) {

	m := response.CreateRole{
		Status:     r.Status,
		CreateTime: time.Now().Unix(),
		Desc:       r.Desc,
		Name:       r.Name,
	}

	u, err := model.CreateRole(m)

	if err != nil {
		global.GVA_LOG.Error("model.CreateRole(m)", zap.Any("m", m), zap.Error(err))
		return
	}
	return u
}

func EditRole(r request.EditRole) bool {

	res, err := model.EditRole(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditRole(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}
}

func DeleteRole(r request.GetById) bool {
	res, err := model.DeleteRoleId(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditRole(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}

}

func UpdateStatusRole(r request.UpdateStatusRequest) bool {
	res, err := model.UpdateStatusRoleId(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditRole(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}
}

func SetRules(r request.SetRules) (b bool) {
	ruleIdList, err := model.GetRoleRuleIdList(r.RoleId)
	if err != nil {
		global.GVA_LOG.Error("model.GetManagerRuleID(info.RoleId) failed", zap.Int64("info.RoleId", r.RoleId), zap.Error(err))
		return
	}

	var RuleIdArr []int64
	for _, v := range ruleIdList {
		RuleIdArr = append(RuleIdArr, v.RuleId)
	}

	addIds := difference(r.RuleIds, RuleIdArr)

	delIds := difference(RuleIdArr, r.RuleIds)

	if len(addIds) > 0 { //添加
		var roleRuleArr []response.RoleRule
		for _, v := range addIds {
			rule := response.RoleRule{RoleId: r.RoleId, RuleId: v}
			roleRuleArr = append(roleRuleArr, rule)
		}
		res, err := model.SetRules(roleRuleArr)
		if err != nil {
			global.GVA_LOG.Error("model.EditRole(r)", zap.Any("r", r), zap.Error(err))
			return false
		}
		if res == 1 {
			return true
		} else {
			return false
		}

	} else if len(addIds) == 0 && len(delIds) == 0 { //没有修改
		return true
	} else if len(delIds) > 0 { //删除
		var ids []int64
		for _, v := range delIds {
			for _, j := range ruleIdList {
				if v == j.RuleId {
					ids = append(ids, j.Id)
				}
			}
		}
		res, err := model.DeleteRules(ids)
		if err != nil {
			global.GVA_LOG.Error("model.EditRole(r)", zap.Any("r", r), zap.Error(err))
			return false
		}
		if res == 1 {
			return true
		} else {
			return false
		}
	}
	return true
}

// difference returns the elements in `a` that aren't in `b`.
func difference(a, b []int64) []int64 {
	mb := make(map[int64]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []int64
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}
