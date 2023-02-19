package request

type AddRole struct {
	Name   string `json:"name"`
	Status int64  `json:"status"`
	Desc   string `json:"desc"`
}

type EditRole struct {
	GetById
	AddRole
}

type SetRules struct {
	RoleId  int64   `json:"role_id"`
	RuleIds []int64 `json:"rule_ids"`
}
