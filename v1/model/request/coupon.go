package request

type AddCoupon struct {
	Name       string `json:"name"`
	Order      int64  `json:"order"`
	Type       int64  `json:"type"`
	Value      string `json:"value"`
	Total      int64  `json:"total"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
	MinPrice   string `json:"min_price"`
	StartTime  int64  `json:"start_time"`
	EndTime    int64  `json:"end_time"`
}

type EditCoupon struct {
	GetById
	AddCoupon
}
