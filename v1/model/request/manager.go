package request

// Login 登录请求参数
type Login struct {
	Username  string `json:"username" form:"username" binding:"required"` //用户名
	Password  string `json:"password" form:"password" binding:"required"` //用户密码
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}

type ManagerList struct {
	PageInfo
	Keyword string `json:"keyword" form:"keyword"`
}

type ManagerRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	RoleId   int64  `json:"role_id" form:"role_id" binding:"required"`
	Status   int64  `json:"status" form:"status" binding:"required"`
	Avatar   string `json:"avatar" form:"avatar"`
}

type EditRequest struct {
	GetById
	ManagerRequest
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password" form:"old_password" binding:"required"`
	Password    string `json:"password" form:"password" binding:"required"`
}
