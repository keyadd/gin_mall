package response

type GoodsSkusCard struct {
	Id                 int64                `json:"id"` //id
	GoodsId            int64                `json:"goods_id"`
	Name               string               `json:"name"`
	Order              int64                `json:"order"`
	Type               int64                `json:"type"`
	GoodsSkusCardValue []GoodsSkusCardValue `json:"goods_skus_card_value" gorm:"-"`
}

type GoodsSkusCardNames struct {
	RouterPath string `json:"router_path" `
	Method     string `json:"method"`
}
