package response

type Menu struct {
	Id         int64   `json:"id" gorm:"id"'`          //id
	RuleId     int64   `json:"rule_id" gorm:"rule_id"` //站点id
	Status     int64   `json:"status" gorm:"status"`
	CreateTime string  `json:"create_time" gorm:"create_time"`
	UpdateTime string  `json:"update_time" gorm:"update_time"`
	Name       string  `json:"name" gorm:"name"`
	Desc       string  `json:"desc" gorm:"desc"`
	FrontPath  string  `json:"frontpath" gorm:"frontpath"`
	Condition  string  `json:"condition" gorm:"condition"`
	Menu       int64   `json:"menu" gorm:"menu"`
	Order      int64   `json:"order" gorm:"order"`
	Icon       string  `json:"icon" gorm:"icon"`
	Method     string  `json:"method" gorm:"method"`
	Child      []*Menu `json:"child" gorm:"-"` //多个子级目录
}

type RuleNames struct {
	Condition string `json:"condition" gorm:"condition"`
	Method    string `json:"method" gorm:"method"`
}
