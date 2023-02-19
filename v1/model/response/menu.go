package response

type Menu struct {
	Id         int64   `json:"id" gorm:"id"'`          //id
	RuleId     int64   `json:"rule_id" gorm:"rule_id"` //站点id
	Status     int64   `json:"status" gorm:"status"`
	CreateTime int64   `json:"create_time" gorm:"create_time"`
	UpdateTime int64   `json:"update_time" gorm:"update_time"`
	Name       string  `json:"name" gorm:"name"`
	Desc       string  `json:"desc" gorm:"desc"`
	FrontPath  string  `json:"front_path"`
	Menu       int64   `json:"menu" gorm:"menu"`
	Order      int64   `json:"order" gorm:"order"`
	Icon       string  `json:"icon" gorm:"icon"`
	Child      []*Menu `json:"children" gorm:"-"` //多个子级目录
}

type MenuNames struct {
	RouterPath string `json:"router_path" `
	Method     string `json:"method" gorm:"method"`
}

type MenuResponse struct {
	List []*Menu `json:"list"`
}
