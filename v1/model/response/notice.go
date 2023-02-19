package response

type Notice struct {
	Id         int64  `json:"id" `
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
	Title      string `json:"title" form:"title"`
	Content    string `json:"content" form:"content"`
}

type NoticeList struct {
	Id         int64  `json:"id" `
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	Title      string `json:"title" form:"title"`
	Content    string `json:"content" form:"content"`
	Order      int64  `json:"order"`
}
