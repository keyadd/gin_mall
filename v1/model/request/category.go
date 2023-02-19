package request

type CategoryRequest struct {
	Name       string `json:"name"`
	Status     int64  `json:"status"`
	Order      int64  `json:"order"`
	CategoryId int64  `json:"category_id"`
}

type CategoryEditRequest struct {
	GetById
	CategoryRequest
}
