package request

type OrderListRequest struct {
	PageInfo
	Tab       string `json:"tab"`
	No        string `json:"no"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
}

type OrderEditRequest struct {
	GetById
}

type OrderSend struct {
	GetById
	ExpressCompany string `json:"express_company" form:"express_company"`
	ExpressNo      string `json:"express_no" form:"express_no"`
}

type OrderHandleRefund struct {
	GetById
	Agree          int64  `json:"agree"`
	DisagreeReason string `json:"disagree_reason"`
}
