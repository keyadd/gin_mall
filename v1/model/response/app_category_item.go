package response

type AppCategoryItem struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Cover      string `json:"cover"`
	CategoryId int64  `json:"category_id"`
	GoodsId    int64  `json:"goods_id"`
	Order      int64  `json:"order"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

type AppCategoryItemList struct {
	Id         int64     `json:"id" `
	Status     int64     `json:"status"`
	CreateTime int64     `json:"create_time"`
	UpdateTime int64     `json:"update_time"`
	Username   string    `json:"username"`
	Avatar     string    `json:"avatar"`
	RoleId     int64     `json:"role_id"`
	Super      int64     `json:"super"`
	Role       *InfoRole `json:"role"`
}

type AppCategoryItemResList struct {
	AppCategoryItemList []*AppCategoryItemList `json:"list"`
	TotalCount          int64                  `json:"totalCount"`
	InfoRole            []*InfoRole            `json:"roles"`
}

type AppCategoryItemInfo struct {
	Id        int64     `json:"id" `
	Username  string    `json:"username" `
	Avatar    string    `json:"avatar"`
	Super     int64     `json:"super"`
	Role      *InfoRole `json:"role"`
	Menus     []*Menu   `json:"menus"`
	RuleNames []string  `json:"ruleNames"`
}
type CreateAppCategoryItem struct {
	Id         int64  `json:"id" `
	Status     int64  `json:"status"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
	Username   string `json:"username"`
	RoleId     int64  `json:"role_id"`
	Avatar     string `json:"avatar"`
	Super      int64  `json:"super"`
}
