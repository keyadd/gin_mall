package v1

import (
	"gin_mall/v1/api"
	"github.com/gin-gonic/gin"
)

func InitCoupon(Router *gin.RouterGroup) {
	CouponRouter := Router.Group("/coupon") //?
	{
		CouponRouter.POST("/list", api.CouponList)                  //后台用户列表
		CouponRouter.POST("/update_status", api.CouponUpdateStatus) //修改状态
		CouponRouter.POST("/create", api.CouponCreate)              //创建用户
		CouponRouter.POST("/edit", api.CouponEdit)                  //新增用户
		CouponRouter.POST("/delete_all", api.CouponDelete)          //删除用户
	}

}
