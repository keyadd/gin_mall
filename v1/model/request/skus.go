package request

type AddSkus struct {
	Name    string `json:"name"`
	Order   int64  `json:"order"`
	Status  int64  `json:"status"`
	Default string `json:"default"`
}

type EditSkus struct {
	GetById
	AddSkus
}

type Skus struct {
	GoodsSkusCardId string `json:"goods_skus_card_id" form:"goods_skus_card_id"`
	Id              string `json:"id" form:"id"`
	Name            string `json:"name" form:"name"`
	Order           int64  `json:"order" form:"order"`
	Text            string `json:"text" form:"text"`
	Value           string `json:"value" form:"value"`
}
