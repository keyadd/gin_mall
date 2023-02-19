package response

type GoodsBanner struct {
	Id         int64  `json:"id"`
	GoodsId    int64  `json:"goods_id"`
	Url        string `json:"url"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}
