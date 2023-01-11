package v1

import (
	"gin_mall/v1/api"
	"github.com/gin-gonic/gin"
)

func InitBase(Router *gin.RouterGroup) {
	ManagerRouter := Router.Group("")
	{
		ManagerRouter.POST("/login", api.Login)
	}

}
