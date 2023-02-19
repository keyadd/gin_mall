package request

type AddGoods struct {
	Title        string `json:"title"`
	CategoryId   int64  `json:"category_id"`
	Cover        string `json:"cover"`
	Desc         string `json:"desc"`
	Unit         string `json:"unit"`
	Stock        int64  `json:"stock"`
	MinStock     int64  `json:"min_stock"`
	Status       int64  `json:"status"`
	StockDisplay int64  `json:"stock_display"`
	MinPrice     string `json:"min_price"`
	MinOprice    string `json:"min_oprice"`
	CreateTime   int64  `json:"create_time"`
	UpdateTime   int64  `json:"update_time"`
}

type EditGoods struct {
	GetById
	AddGoods
	Content string `json:"content"`
}

type GoodsList struct {
	PageInfo
	Tab        string `json:"tab"`
	Title      string `json:"title"`
	CategoryId string `json:"category_id"`
}

type GoodsCheck struct {
	GetById
	IsCheck int64 `json:"isCheck"`
}
