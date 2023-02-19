package response

type GoodsSkus struct {
	Id      int64   `json:"id"`
	GoodsId int64   `json:"goods_id"`
	Image   string  `json:"image"`
	Pprice  float64 `json:"pprice"`
	Oprice  float64 `json:"oprice"`
	Cprice  float64 `json:"cprice"`
	Stock   int64   `json:"stock"`
	Volume  float64 `json:"volume"`
	Weight  float64 `json:"weight"`
	Code    string  `json:"code"`
	Skus    string  `json:"skus"`
}

type GoodsSkusJson struct {
	Id   int64  `json:"id"`
	Skus string `json:"skus"`
}

type GoodsSkusRead struct {
	Id      int64       `json:"id"`
	GoodsId int64       `json:"goods_id"`
	Image   string      `json:"image"`
	Pprice  float64     `json:"pprice"`
	Oprice  float64     `json:"oprice"`
	Cprice  float64     `json:"cprice"`
	Stock   int64       `json:"stock"`
	Volume  float64     `json:"volume"`
	Weight  float64     `json:"weight"`
	Code    string      `json:"code"`
	Skus    interface{} `json:"skus"gorm:"-"`
}

type UpdateGoodsSkus struct {
	Id      int64   `json:"id"`
	GoodsId int64   `json:"goods_id"`
	Image   string  `json:"image"`
	Pprice  float64 `json:"pprice"`
	Oprice  float64 `json:"oprice"`
	Cprice  float64 `json:"cprice"`
	Stock   int64   `json:"stock"`
	Volume  float64 `json:"volume"`
	Weight  float64 `json:"weight"`
	Code    string  `json:"code"`
	Skus    string  `json:"skus"`
}

type GoodsSkusArr struct {
	Skus string `json:"skus"`
}
