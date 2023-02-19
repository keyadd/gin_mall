package request

type CreateApiRequest struct {
	RuleId     int64  `json:"rule_id" ` //站点id
	Status     int64  `json:"status"`
	Name       string `json:"name"`
	RouterPath string `json:"router_path" `
	Menu       int64  `json:"menu"`
	Order      int64  `json:"order"`
	Method     string `json:"method"`
	ApiGroup   string `json:"api_group"`
}

type ApiRequest struct {
	PageInfo
	Keyword string `json:"keyword" form:"keyword"`
}

type EditApiRequest struct {
	GetById
	CreateApiRequest
}
