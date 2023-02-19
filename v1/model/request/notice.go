package request

type NoticeRequest struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}

type NoticeEdit struct {
	GetById
	NoticeRequest
}

type NoticeList struct {
	PageInfo
}
