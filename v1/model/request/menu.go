package request

type MenuRequest struct {
	RuleId    int64  `json:"rule_id" form:"rule_id"` //站点id
	Status    int64  `json:"status" form:"status"`
	Name      string `json:"name" form:"name"`
	FrontPath string `json:"front_path" form:"front_path"`
	Menu      int64  `json:"menu" form:"menu"`
	Order     int64  `json:"order" form:"order"`
	Icon      string `json:"icon" form:"icon"`
}

type EditMenuRequest struct {
	GetById
	MenuRequest
}
