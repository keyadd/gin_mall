package response

type GoodsList struct {
	Id           int64         `json:"id"`
	Title        string        `json:"title"`
	CategoryId   int64         `json:"category_id"`
	Cover        string        `json:"cover"`
	Rating       int64         `json:"rating"`
	SaleCount    int64         `json:"sale_count"`
	ReviewCount  int64         `json:"review_count"`
	MinPrice     string        `json:"min_price"`
	MinOprice    string        `json:"min_oprice"`
	Desc         string        `json:"desc"`
	Unit         string        `json:"unit"`
	Stock        int64         `json:"stock"`
	MinStock     int64         `json:"min_stock"`
	IsCheck      int64         `json:"ischeck"`
	Status       int64         `json:"status"`
	StockDisplay int64         `json:"stock_display"`
	ExpressId    int64         `json:"express_id"`
	SkuType      int64         `json:"sku_type"`
	SkuValue     string        `json:"sku_value"`
	Content      string        `json:"content"`
	Discount     int64         `json:"discount"`
	CreateTime   string        `json:"create_time"`
	UpdateTime   string        `json:"update_time"`
	DeleteTime   int64         `json:"delete_time"`
	Order        int64         `json:"order"`
	Category     Category      `json:"category" gorm:"foreignKey:Id;references:CategoryId"`
	GoodsBanner  []GoodsBanner `json:"goods_banner"  gorm:"foreignKey:GoodsId;"`

	GoodsSkus     []GoodsSkus     `json:"goods_skus"  gorm:"foreignKey:GoodsId;references:Id"`
	SkusInfo      []SkusInfo      `json:"skus_info" gorm:"-"`
	GoodsSkusCard []GoodsSkusCard `json:"goods_skus_card" gorm:"foreignKey:GoodsId;references:Id"`
}

type GoodsCreate struct {
	Id            int64           `json:"id"`
	Title         string          `json:"title"`
	CategoryId    int64           `json:"category_id"`
	Cover         string          `json:"cover"`
	Rating        int64           `json:"rating"`
	SaleCount     int64           `json:"sale_count"`
	ReviewCount   int64           `json:"review_count"`
	MinPrice      string          `json:"min_price"`
	MinOprice     string          `json:"min_oprice"`
	Desc          string          `json:"desc"`
	Unit          string          `json:"unit"`
	Stock         int64           `json:"stock"`
	MinStock      int64           `json:"min_stock"`
	IsCheck       int64           `json:"ischeck"`
	Status        int64           `json:"status"`
	StockDisplay  int64           `json:"stock_display"`
	ExpressId     int64           `json:"express_id"`
	SkuType       int64           `json:"sku_type"`
	SkuValue      string          `json:"sku_value"`
	Content       string          `json:"content"`
	Discount      int64           `json:"discount"`
	CreateTime    int64           `json:"create_time"`
	UpdateTime    int64           `json:"update_time"`
	DeleteTime    int64           `json:"delete_time"`
	Order         int64           `json:"order"`
	Category      *Category       `json:"category" gorm:"-"`
	GoodsBanner   []GoodsBanner   `json:"goods_banner"  gorm:"-"`
	GoodsSkus     []SkusInfo      `json:"goods_skus"  gorm:"-"`
	GoodsSkusCard []GoodsSkusCard `json:"goods_skus_card" gorm:"-" `
}

type GoodsRead struct {
	Id            int64           `json:"id"`
	Title         string          `json:"title"`
	CategoryId    int64           `json:"category_id"`
	Content       string          `json:"content"`
	Cover         string          `json:"cover"`
	Desc          string          `json:"desc"`
	Unit          string          `json:"unit"`
	Stock         int64           `json:"stock"`
	MinStock      int64           `json:"min_stock"`
	Status        int64           `json:"status"`
	Order         int64           `json:"order"`
	StockDisplay  int64           `json:"stock_display"`
	MinPrice      string          `json:"min_price"`
	MinOprice     string          `json:"min_oprice"`
	CreateTime    int64           `json:"create_time"`
	UpdateTime    int64           `json:"update_time"`
	DeleteTime    int64           `json:"delete_time"`
	IsCheck       int64           `json:"ischeck"`
	Discount      float64         `json:"discount"`
	ExpressId     int64           `json:"express_id"`
	Rating        float64         `json:"rating"`
	ReviewCount   int64           `json:"review_count"`
	SaleCount     int64           `json:"sale_count"`
	SkuType       int64           `json:"sku_type"`
	SkuValue      string          `json:"sku_value"`
	GoodsBanner   []GoodsBanner   `json:"goods_banner" gorm:"-"`
	GoodsSkus     []GoodsSkusRead `json:"goods_skus" gorm:"-"`
	GoodsSkusCard []GoodsSkusCard `json:"goods_skus_card" gorm:"-"`
}

type GoodsOrder struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Cover   string `json:"cover"`
	SkuType int64  `json:"sku_type"`
}

type GoodsListRes struct {
	List     interface{} `json:"list"`
	Category interface{} `json:"category"`
	Total    int64       `json:"total"`
}
