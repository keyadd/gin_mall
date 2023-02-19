package v1

import (
	"gin_mall/v1/api"
	"github.com/gin-gonic/gin"
)

func InitGoods(Router *gin.RouterGroup) {
	GoodsRouter := Router.Group("/goods") //？
	{
		GoodsRouter.POST("/create", api.GoodsCreate)              //增加商品
		GoodsRouter.POST("/edit", api.GoodsEdit)                  //修改商品
		GoodsRouter.POST("/change_status", api.GoodsChangeStatus) //商品上架下架
		GoodsRouter.POST("/list", api.GoodsList)                  //商品列表？
		GoodsRouter.POST("/delete_all", api.GoodsDeleteAll)       //删除商品到回收站
		GoodsRouter.POST("/destroy", api.GoodsDestroyAll)         //彻底删除商品
		GoodsRouter.POST("/check", api.GoodsCheck)                //审核商品
		GoodsRouter.POST("/read", api.GoodsRead)                  //查看商品详细
		GoodsRouter.POST("/banners_set", api.GoodsBannerSet)      //设置商品轮播图
		GoodsRouter.POST("/update_skus", api.GoodsUpdateSkus)     //设置商品规格
	}

}
