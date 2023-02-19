package response

type Category struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Status     int64  `json:"status"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
	CategoryId int64  `json:"category_id"`
	Order      int64  `json:"order"`
	//Child      []*Category `json:"children" gorm:"-"` //多个子级目录
}

type CategoryList struct {
	Id         int64           `json:"id"`
	Name       string          `json:"name"`
	Status     int64           `json:"status"`
	CreateTime string          `json:"create_time"`
	UpdateTime string          `json:"update_time"`
	CategoryId int64           `json:"category_id"`
	Order      int64           `json:"order"`
	Child      []*CategoryList `json:"children" gorm:"-"` //多个子级目录
}

type CategoryResList struct {
	CategoryList []*CategoryList `json:"list"`
	TotalCount   int64           `json:"totalCount"`
	InfoRole     []*InfoRole     `json:"roles"`
}

type CategoryInfo struct {
	Id        int64     `json:"id" `
	Username  string    `json:"username" `
	Avatar    string    `json:"avatar"`
	Super     int64     `json:"super"`
	Role      *InfoRole `json:"role"`
	Menus     []*Menu   `json:"menus"`
	RuleNames []string  `json:"ruleNames"`
}
type CreateCategory struct {
	Id         int64  `json:"id" `
	Status     int64  `json:"status"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
	Username   string `json:"username"`
	RoleId     int64  `json:"role_id"`
	Avatar     string `json:"avatar"`
	Super      int64  `json:"super"`
}
