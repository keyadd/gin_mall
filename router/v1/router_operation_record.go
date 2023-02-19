package v1

import (
	"gin_mall/v1/api"
	"github.com/gin-gonic/gin"
)

func InitOperationRecord(Router *gin.RouterGroup) {
	OperationRecordRouter := Router.Group("/operation_record") //?
	{
		OperationRecordRouter.POST("/list", api.OperationRecordList)     //公告列表
		OperationRecordRouter.POST("/delete", api.OperationRecordDelete) //删除公告
	}

}
