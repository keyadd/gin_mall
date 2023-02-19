package response

type GoodsComment struct {
	Id         int64  `json:"id" `
	ShopId     int64  `json:"shop_id"`
	OrderId    int64  `json:"order_id"`
	GoodsId    int64  `json:"goods_id"`
	Num        int64  `json:"num"`
	Price      string `json:"price"`
	Rating     int64  `json:"rating"`
	Review     string `json:"review"`
	ReviewTime int64  `json:"review_time"`
	CreateTime int64  `json:"create_time"`
	SkusType   int64  `json:"skus_type"`
	GoodsNum   int64  `json:"goods_num"`
	UserId     int64  `json:"user_id"`
	Extra      string `json:"extra"`
	Status     int64  `json:"status"`
	Username   string `json:"username" form:"username" binding:"required"`
	NickName   string `json:"nickname"`
	Avatar     string `json:"avatar" form:"avatar"`
	Title      string `json:"title"`
	Cover      string `json:"cover"`
}

type OrderItem struct {
	Id      int64  `json:"id" `
	OrderId int64  `json:"order_id"`
	ShopId  int64  `json:"shop_id"`
	Num     int64  `json:"num"`
	Price   string `json:"price"`
	GoodsId int64  `json:"goods_id"`

	SkusType  int64      `json:"skus_type"`
	GoodsItem GoodsOrder `json:"goods_item" gorm:"-"`
	GoodsSkus []SkusInfo `json:"goods_skus" gorm:"-"`
}

type GoodsCommentList struct {
	Id               int64     `json:"id" `
	Status           int64     `json:"status"`
	CreateTime       int64     `json:"create_time"`
	UpdateTime       int64     `json:"update_time"`
	GoodsCommentname string    `json:"GoodsCommentname"`
	Avatar           string    `json:"avatar"`
	RoleId           int64     `json:"role_id"`
	Super            int64     `json:"super"`
	Role             *InfoRole `json:"role"`
}

type GoodsCommentResList struct {
	GoodsCommentList []*GoodsCommentList `json:"list"`
	TotalCount       int64               `json:"totalCount"`
	InfoRole         []*InfoRole         `json:"roles"`
}

type GoodsCommentInfo struct {
	Id               int64     `json:"id" `
	GoodsCommentname string    `json:"GoodsCommentname" `
	Avatar           string    `json:"avatar"`
	Super            int64     `json:"super"`
	Role             *InfoRole `json:"role"`
	Menus            []*Menu   `json:"menus"`
	RuleNames        []string  `json:"ruleNames"`
}
type CreateGoodsComment struct {
	Id                  int64  `json:"id" `
	GoodsCommentname    string `json:"GoodsCommentname" form:"GoodsCommentname" binding:"required"`
	Password            string `json:"password" form:"password" binding:"required"`
	Status              int64  `json:"status" form:"status" binding:"required"`
	Avatar              string `json:"avatar" form:"avatar"`
	NickName            string `json:"nick_name"`
	Email               string `json:"email"`
	Phone               int64  `json:"phone"`
	GoodsCommentLevelId int64  `json:"GoodsComment_level_id"`
	CreateTime          int64  `json:"create_time"`
	UpdateTime          int64  `json:"update_time"`
}
