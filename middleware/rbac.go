package middleware

import (
	"gin_mall/global"
	"gin_mall/utils"
	"gin_mall/v1/service"
	"github.com/gin-gonic/gin"
)

func RbacHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取请求的URI
		obj := c.Request.URL.RequestURI()
		path := obj[9:len(obj)]
		// 获取用户id
		UserIdName := global.GVA_CONFIG.JWT.UserIdName
		userId, _ := c.Get(UserIdName)
		// 获取请求方法
		method := c.Request.Method
		ruleId := service.GetRule(path, method, userId.(int64))

		if !ruleId {
			utils.ResponseError(c, global.CodeRuleNotExist)
			c.Abort()
			return
		}
		c.Next()
	}
}
