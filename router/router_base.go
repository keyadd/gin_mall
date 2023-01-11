package router

import (
	"gin_mall/middleware"
	v1Router "gin_mall/router/v1"
	"github.com/gin-gonic/gin"
	"net/http"
	//_ "gin_mall/docs" // 千万不要忘了导入把你上一步生成的docs
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(middleware.CorsMiddleware()).Use(middleware.GinLogger()).Use(middleware.GinRecovery(true)).Use(middleware.Sentinel())
	v1 := r.Group("/admin/v1")
	{

		//无token 访问
		PublicGroup := v1.Group("")
		{
			v1Router.InitBase(PublicGroup)
			// 后台登录路由

		}
		//token 访问
		PrivateGroup := v1.Group("")
		PrivateGroup.Use(middleware.JWTAuthMiddleware())
		{
			v1Router.InitManager(PrivateGroup)
			v1Router.InitImageClass(PrivateGroup)

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
