package response

type Manager struct {
	Id         int64  `json:"id" `
	Status     int64  `json:"status"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
	Username   string `json:"username"`
	Password   string `json:"password" `
	RoleId     int64  `json:"role_id"`
	Avatar     string `json:"avatar"`
	Super      int64  `json:"super"`
}

type ManagerList struct {
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

type ResList struct {
	ManagerList []*ManagerList `json:"list"`
	TotalCount  int64          `json:"totalCount"`
	InfoRole    []*InfoRole    `json:"roles"`
}

type ManagerInfo struct {
	Id        int64     `json:"id" `
	Username  string    `json:"username" `
	Avatar    string    `json:"avatar"`
	Super     int64     `json:"super"`
	Role      *InfoRole `json:"role"`
	Menus     []*Menu   `json:"menus"`
	RuleNames []string  `json:"ruleNames"`
}
type CreateManager struct {
	Id         int64  `json:"id" `
	Status     int64  `json:"status"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
	Username   string `json:"username"`
	RoleId     int64  `json:"role_id"`
	Avatar     string `json:"avatar"`
	Super      int64  `json:"super"`
}
