package request

type UserRequest struct {
	Username    string `json:"username" form:"username" binding:"required"`
	Password    string `json:"password" form:"password" binding:"required"`
	Status      int64  `json:"status" form:"status" binding:"required"`
	Avatar      string `json:"avatar" form:"avatar"`
	NickName    string `json:"nickname"`
	Email       string `json:"email"`
	Phone       int64  `json:"phone"`
	UserLevelId int64  `json:"user_level_id"`
}

type UserListRequest struct {
	PageInfo
	Keyword     string `json:"keyword" form:"keyword"`
	UserLevelId int64  `json:"user_level_id" form:"user_level_id"`
}

type UserEditRequest struct {
	GetById
	UserRequest
}
