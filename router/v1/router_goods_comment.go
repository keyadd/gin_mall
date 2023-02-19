package v1

import (
	"gin_mall/v1/api"
	"github.com/gin-gonic/gin"
)

func InitGoodsComment(Router *gin.RouterGroup) {
	GoodsCommentRouter := Router.Group("/goods_comment")
	{

		GoodsCommentRouter.POST("/list", api.GoodsCommentList)                  //商品评价列表
		GoodsCommentRouter.POST("/update_status", api.GoodsCommentUpdateStatus) //修改商品评价状态
		GoodsCommentRouter.POST("/edit", api.GoodsCommentEdit)                  //回复商品评价
	}
}
