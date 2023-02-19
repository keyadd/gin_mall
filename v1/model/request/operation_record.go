package request

type OperationRecord struct {
	PageInfo
	Keyword string `json:"keyword" form:"keyword"`
}
