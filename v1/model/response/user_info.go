package response

type UserInfo struct {
	Id       int64  `json:"id"`
	UserId   int64  `json:"user_id"`
	Sex      int64  `json:"sex"`
	Birthday string `json:"birthday"`
	Name     string `json:"name"`
}
