package response

type InfoRole struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Role struct {
	Id         int64      `json:"id"`
	Name       string     `json:"name"`
	CreateTime string     `json:"create_time"`
	UpdateTime string     `json:"update_time"`
	Status     int64      `json:"status"`
	Desc       string     `json:"desc"`
	RoleRule   []RoleRule `json:"rules" gorm:"foreignKey:RoleId;references:Id"`
}

type CreateRole struct {
	Id         int64      `json:"id"`
	Name       string     `json:"name"`
	CreateTime int64      `json:"create_time"`
	UpdateTime int64      `json:"update_time"`
	Status     int64      `json:"status"`
	Desc       string     `json:"desc"`
	RoleRule   []RoleRule `json:"rules" gorm:"foreignKey:RoleId;references:Id"`
}
