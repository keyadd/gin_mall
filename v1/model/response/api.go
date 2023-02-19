package response

type Api struct {
	Id         int64  `json:"id" gorm:"id"'`          //id
	RuleId     int64  `json:"rule_id" gorm:"rule_id"` //站点id
	Status     int64  `json:"status" gorm:"status"`
	CreateTime int64  `json:"create_time" gorm:"create_time"`
	UpdateTime int64  `json:"update_time" gorm:"update_time"`
	Name       string `json:"name" gorm:"name"`
	Desc       string `json:"desc" gorm:"desc"`
	RouterPath string `json:"router_path" `
	Menu       int64  `json:"menu" gorm:"menu"`
	Order      int64  `json:"order" gorm:"order"`
	ApiGroup   string `json:"api_group"`
	Method     string `json:"method" gorm:"method"`
}

type ApiList struct {
	Id         int64  `json:"id" gorm:"id"'`          //id
	RuleId     int64  `json:"rule_id" gorm:"rule_id"` //站点id
	Status     int64  `json:"status" gorm:"status"`
	CreateTime string `json:"create_time" gorm:"create_time"`
	UpdateTime string `json:"update_time" gorm:"update_time"`
	Name       string `json:"name" gorm:"name"`
	Desc       string `json:"desc" gorm:"desc"`
	RouterPath string `json:"router_path" `
	Menu       int64  `json:"menu" gorm:"menu"`
	Order      int64  `json:"order" gorm:"order"`
	ApiGroup   string `json:"api_group"`
	Method     string `json:"method" gorm:"method"`
}

type ApiNames struct {
	RouterPath string `json:"router_path" `
	Method     string `json:"method" gorm:"method"`
}

type ApiResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Children []*Api `json:"children"`
}

type ApiTreeResponse struct {
}
