package response

type GoodsSkusCardValue struct {
	Id              int64  `json:"id" gorm:"id"'`       //id
	GoodsSkusCardId int64  `json:"goods_skus_card_id" ` //站点id
	Name            string `json:"name"`
	Order           int64  `json:"order"`
	Value           string `json:"value"`
}

type GoodsSkusCardValueNames struct {
	RouterPath string `json:"router_path" `
	Method     string `json:"method" gorm:"method"`
}
