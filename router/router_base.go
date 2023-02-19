package router

import (
	"gin_mall/global"
	"gin_mall/middleware"
	v1Router "gin_mall/router/v1"
	"github.com/gin-gonic/gin"
	"net/http"
	//_ "gin_mall/docs" // 千万不要忘了导入把你上一步生成的docs
)

func Setup() *gin.Engine {
	if global.GVA_CONFIG.System.Env == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	r.Use(middleware.CorsMiddleware()).Use(middleware.GinLogger()).Use(middleware.GinRecovery(true)).Use(middleware.Sentinel())
	v1 := r.Group("/admin/v1")
	{

		//无token 访问
		PublicGroup := v1.Group("")
		{
			v1Router.InitPublic(PublicGroup) // 后台登录路由

		}
		//token 访问
		PrivateGroup := v1.Group("")
		PrivateGroup.Use(middleware.JWTAuthMiddleware()).Use(middleware.RbacHandler()).Use(middleware.OperationRecord())
		{
			v1Router.InitManager(PrivateGroup)
			v1Router.InitImageClass(PrivateGroup)
			v1Router.InitNotice(PrivateGroup)
			v1Router.InitImage(PrivateGroup)
			v1Router.InitMenu(PrivateGroup)
			v1Router.InitRole(PrivateGroup)
			v1Router.InitSkus(PrivateGroup)
			v1Router.InitCoupon(PrivateGroup)
			v1Router.InitGoods(PrivateGroup)
			v1Router.InitGoodsSkusCard(PrivateGroup)
			v1Router.InitCategory(PrivateGroup)
			v1Router.InitUser(PrivateGroup)
			v1Router.InitUserLevel(PrivateGroup)
			v1Router.InitOrder(PrivateGroup)
			v1Router.InitApi(PrivateGroup)
			v1Router.InitGoodsSkusCardValue(PrivateGroup)
			v1Router.InitAppCategoryItem(PrivateGroup)
			v1Router.InitDashboard(PrivateGroup)
			v1Router.InitGoodsComment(PrivateGroup)
			v1Router.InitOperationRecord(PrivateGroup)
			v1Router.InitAgent(PrivateGroup)
			v1Router.InitSystemRouter(PrivateGroup)

		}
	}
	r.GET("/", func(c *gin.Context) {

		c.String(http.StatusOK, "OK")

	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg":    "not found",
			"status": "404",
		})
	})
	return r

}
