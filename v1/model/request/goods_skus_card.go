package request

type GoodsSkusCardRequest struct {
	GoodsId int64  `json:"goods_id"`
	Name    string `json:"name"`
	Order   int64  `json:"order"`
	Type    int64  `json:"type"`
}

type EditGoodsSkusCardRequest struct {
	GetById
	GoodsSkusCardRequest
}
