package response

type Coupon struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Order      int64  `json:"order"`
	Type       int64  `json:"type"`
	Value      string `json:"value"`
	Total      int64  `json:"total"`
	Used       int64  `json:"used"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
	MinPrice   string `json:"min_price"`
	StartTime  int64  `json:"start_time"`
	EndTime    int64  `json:"end_time"`
	Status     int64  `json:"status"`
	Desc       string `json:"desc"`
}

type CouponList struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Order      int64  `json:"order"`
	Type       int64  `json:"type"`
	Value      string `json:"value"`
	Total      int64  `json:"total"`
	Used       int64  `json:"used"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	MinPrice   string `json:"min_price"`
	StartTime  int64  `json:"start_time"`
	EndTime    int64  `json:"end_time"`
	Status     int64  `json:"status"`
	Desc       string `json:"desc"`
}
