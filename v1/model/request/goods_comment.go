package request

type GoodsCommentRequest struct {
	GoodsCommentname    string `json:"GoodsCommentname" form:"GoodsCommentname" binding:"required"`
	Password            string `json:"password" form:"password" binding:"required"`
	Status              int64  `json:"status" form:"status" binding:"required"`
	Avatar              string `json:"avatar" form:"avatar"`
	NickName            string `json:"nick_name"`
	Email               string `json:"email"`
	Phone               int64  `json:"phone"`
	GoodsCommentLevelId int64  `json:"GoodsComment_level_id"`
}

type GoodsCommentEditRequest struct {
	GetById
	GoodsCommentRequest
}
