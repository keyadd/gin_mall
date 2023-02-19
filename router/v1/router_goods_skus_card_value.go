package v1

import (
	"gin_mall/v1/api"
	"github.com/gin-gonic/gin"
)

func InitGoodsSkusCardValue(Router *gin.RouterGroup) {
	GoodsSkusCardValueRouter := Router.Group("/goods_skus_card_value") //？
	{
		GoodsSkusCardValueRouter.POST("/set", api.GoodsSkusCardValueSet)       //修改状态
		GoodsSkusCardValueRouter.POST("/create", api.GoodsSkusCardValueCreate) //创建商品规格选项值
		GoodsSkusCardValueRouter.POST("/edit", api.GoodsSkusCardValueEdit)     //新增商品规格选项值
		GoodsSkusCardValueRouter.POST("/delete", api.GoodsSkusCardValueDelete) //删除商品规格选项值
	}

}
