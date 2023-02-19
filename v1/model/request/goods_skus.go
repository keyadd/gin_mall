package request

type SkuValue struct {
	Pprice float64 `json:"pprice"`
	Oprice float64 `json:"oprice"`
	Cprice float64 `json:"cprice"`

	Volume float64 `json:"volume"`
	Weight float64 `json:"weight"`
}

type GoodsSkus struct {
	Id      int64   `json:"id" form:"id"`
	GoodsId int64   `json:"goods_id" form:"goods_id"`
	Image   string  `json:"image" form:"image"`
	Pprice  string  `json:"pprice" form:"pprice"`
	Oprice  string  `json:"oprice" form:"oprice"`
	Cprice  string  `json:"cprice" form:"cprice"`
	Stock   int64   `json:"stock" form:"stock"`
	Volume  float64 `json:"volume" form:"volume"`
	Weight  float64 `json:"weight" form:"weight"`
	Code    string  `json:"code" form:"code"`
	Skus    []Skus  `json:"skus" form:"skus"`
}

type GoodsUpdateSkus struct {
	GetById
	SkuType   int64       `json:"sku_type" form:"sku_type"`
	SkuValue  SkuValue    `json:"sku_value" form:"sku_value"`
	GoodsSkus []GoodsSkus `json:"goods_skus" form:"goods_skus"`
}
