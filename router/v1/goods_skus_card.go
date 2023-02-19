package v1

import (
	"gin_mall/v1/api"
	"github.com/gin-gonic/gin"
)

func InitGoodsSkusCard(Router *gin.RouterGroup) {
	GoodsSkusCardRouter := Router.Group("/goods_skus_card")
	{

		GoodsSkusCardRouter.POST("/update_status", api.GoodsSkusCardUpdateOrder) //修改排序
		GoodsSkusCardRouter.POST("/create", api.GoodsSkusCardCreate)             //创建规格
		GoodsSkusCardRouter.POST("/edit", api.GoodsSkusCardEdit)                 //编辑规格
		GoodsSkusCardRouter.POST("/delete", api.GoodsSkusCardDelete)             //删除规格
	}

}
