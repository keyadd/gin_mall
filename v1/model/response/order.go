package response

type OrderAddress struct {
	Id           int64  `json:"id" `
	UserId       int64  `json:"user_id"`
	Province     string `json:"province"`
	City         string `json:"city"`
	District     string `json:"district"`
	Address      string `json:"address"`
	Zip          int64  `json:"zip"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	LastUsedTime int64  `json:"last_used_time"`
	CreateTime   string `json:"create_time"`
	UpdateTime   string `json:"update_time"`
}

type OrderAddressUpdate struct {
	Id           int64  `json:"id" `
	UserId       int64  `json:"user_id"`
	Province     string `json:"province"`
	City         string `json:"city"`
	District     string `json:"district"`
	Address      string `json:"address"`
	Zip          int64  `json:"zip"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	LastUsedTime int64  `json:"last_used_time"`
	CreateTime   int64  `json:"create_time"`
	UpdateTime   int64  `json:"update_time"`
	DeleteTime   int64  `json:"deleteTime"`
}

type OrderList struct {
	Id             int64       `json:"id" `
	No             string      `json:"no"`
	UserId         int64       `json:"user_id"`
	OrderAddressId int64       `json:"order_address_id"`
	TotalPrice     string      `json:"total_price"`
	Remark         string      `json:"remark"`
	PaidTime       int64       `json:"paid_time"`
	PaymentMethod  string      `json:"payment_method"`
	PaymentNo      string      `json:"payment_no"`
	RefundStatus   string      `json:"refund_status"`
	RefundNo       string      `json:"refund_no"`
	Closed         string      `json:"closed"`
	ShipStatus     string      `json:"ship_status"`
	ShipData       string      `json:"ship_data"`
	Extra          string      `json:"extra"`
	CreateTime     int64       `json:"create_time"`
	UpdateTime     int64       `json:"update_time"`
	Reviewed       int64       `json:"reviewed"`
	CouponUserId   int64       `json:"coupon_user_id"`
	Username       string      `json:"username" form:"username"`
	Nickname       string      `json:"nickname"`
	OrderItems     []OrderItem `json:"order_items" gorm:"-"`
	Name           string      `json:"name"`
	Phone          string      `json:"phone"`
}

type OrderSend struct {
	ExpressCompany string `json:"express_company"`
	ExpressNo      string `json:"express_no"`
	ExpressTime    string `json:"express_time"`
}
