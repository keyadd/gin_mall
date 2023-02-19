package response

type Statistics struct {
	Color string `json:"color"`
	Label string `json:"label"`
	Value string `json:"value"`
}

type AgentList struct {
	Id             int64    `json:"id" `
	Username       string   `json:"username" form:"username" binding:"required"`
	Avatar         string   `json:"avatar" form:"avatar"`
	NickName       string   `json:"nickname"`
	Email          string   `json:"email"`
	Phone          string   `json:"phone"`
	CreateTime     string   `json:"create_time"`
	Status         int64    `json:"status" form:"status" binding:"required"`
	ShareNum       int64    `json:"share_num"`
	ShareOrderNum  int64    `json:"share_order_num"`
	OrderPrice     string   `json:"order_price"`
	Commission     string   `json:"commission"`
	CashOutPrice   string   `json:"cash_out_price"`
	CashOutTime    string   `json:"cash_out_time"`
	NoCashOutPrice string   `json:"no_cash_out_price"`
	P1             int64    `json:"p1"`
	P2             int64    `json:"p2"`
	UserInfo       UserInfo `json:"user_info" gorm:"foreignKey:UserId;"`
}
