package v1

import (
	"gin_mall/v1/api"
	"github.com/gin-gonic/gin"
)

func InitOrder(Router *gin.RouterGroup) {
	OrderRouter := Router.Group("/order") //？
	{

		OrderRouter.POST("/list", api.OrderList)                  //后台分类列表
		OrderRouter.POST("/send", api.OrderSend)                  //订单发货
		OrderRouter.POST("/handle_refund", api.OrderHandleRefund) //是否退货
		OrderRouter.POST("/delete_all", api.OrderDelete)          //订单删除
	}
}
