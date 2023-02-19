package request

type GoodsSkusCardValueRequest struct {
	GoodsSkusCardId int64  `json:"goods_skus_card_id" ` //站点id
	Name            string `json:"name"`
	Order           int64  `json:"order"`
	Value           string `json:"value"`
}

type EditGoodsSkusCardValueRequest struct {
	GetById
	GoodsSkusCardValueRequest
}
