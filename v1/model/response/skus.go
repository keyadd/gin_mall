package response

type Skus struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
	Status     int64  `json:"status"`
	Default    string `json:"default"`
	Order      int64  `json:"order" `
}
type SkusList struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	Status     int64  `json:"status"`
	Default    string `json:"default"`
	Order      int64  `json:"order" `
}

type SkusInfo struct {
	GoodsSkusCardId string `json:"goods_skus_card_id" form:"goods_skus_card_id"`
	Id              string `json:"id" form:"id"`
	Name            string `json:"name" form:"name"`
	Order           int64  `json:"order" form:"order"`
	Text            string `json:"text" form:"text"`
	Value           string `json:"value" form:"value"`
}
