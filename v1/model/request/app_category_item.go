package request

type AppCategoryItemRequest struct {
	GetById
}

type AppCategoryItemCreateRequest struct {
	GoodsIds []int64 `json:"goods_ids"`
	AppCategoryItemRequest
}
