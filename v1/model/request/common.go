package request

// Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page" binding:"required"`
	PageSize int `json:"page_size" form:"page_size" binding:"required"`
}

// Find by id structure
type GetById struct {
	Id int64 `json:"id" form:"id"`
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids" binding:"required"`
}

type UpdateStatusRequest struct {
	GetById
	Status string `json:"status" form:"status" binding:"required"`
}

type StatusIdsRequest struct {
	IdsReq
	Status string `json:"status" form:"status" binding:"required"`
}

// Get role by id structure
type GetAuthorityId struct {
	AuthorityId string
}

type Empty struct{}
