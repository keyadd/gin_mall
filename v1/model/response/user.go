package response

type User struct {
	Id            int64       `json:"id" `
	Username      string      `json:"username" form:"username" binding:"required"`
	Status        int64       `json:"status" form:"status" binding:"required"`
	Avatar        string      `json:"avatar" form:"avatar"`
	NickName      string      `json:"nickname"`
	Email         string      `json:"email"`
	Phone         string      `json:"phone"`
	UserLevelId   int64       `json:"user_level_id"`
	UserLevel     interface{} `json:"user_level" gorm:"-"`
	CreateTime    string      `json:"create_time"`
	UpdateTime    int64       `json:"update_time"`
	LastLoginTime int64       `json:"last_login_time"`
}

type UserList struct {
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

type UserListRes struct {
	UserList  interface{} `json:"list"`
	Total     int64       `json:"total"`
	UserLevel interface{} `json:"user_level"`
}

type UserInfoList struct {
	Id        int64     `json:"id" `
	Username  string    `json:"username" `
	Avatar    string    `json:"avatar"`
	Super     int64     `json:"super"`
	Role      *InfoRole `json:"role"`
	Menus     []*Menu   `json:"menus"`
	RuleNames []string  `json:"ruleNames"`
}
type CreateUser struct {
	Id          int64  `json:"id" `
	Username    string `json:"username" form:"username" binding:"required"`
	Password    string `json:"password" form:"password" binding:"required"`
	Status      int64  `json:"status" form:"status" binding:"required"`
	Avatar      string `json:"avatar" form:"avatar"`
	NickName    string `json:"nickname"`
	Email       string `json:"email"`
	Phone       int64  `json:"phone"`
	UserLevelId int64  `json:"user_level_id"`
	CreateTime  int64  `json:"create_time"`
	UpdateTime  int64  `json:"update_time"`
}

type UserOrder struct {
	Id       int64  `json:"id" gorm:"type:int(11);autoIncrement;primaryKey;column:id;"`
	Username string `json:"username" form:"username"`
	Avatar   string `json:"avatar" form:"avatar"`
	Nickname string `json:"nickname"`
}
