package main

import (
	"fmt"
	"gin_mall/global"
	"gin_mall/initialize"
	"gin_mall/router"
	"gin_mall/validator"
	"go.uber.org/zap"
	"time"
)

// @title web_app
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /

// @host 127.0.0.1:8888

func main() {
	//加载配置文件

	global.GVA_VP = initialize.Viper()

	//初始化日志
	global.GVA_LOG = initialize.Zap()

	//初始化Mysql连接
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	db, _ := global.GVA_DB.DB()
	defer db.Close()
	//初始化Redis连接
	global.GVA_REDIS = initialize.Redis()
	defer global.GVA_REDIS.Close()
	//雪花算法生产唯一ID
	initialize.Snowflake()
	//validator参数验证
	global.TRANS = validator.InitTrans("zh")

	// 初使化 minio client对象
	global.MINIO = initialize.InitMinIO()

	//utils.CreateMinoBuket(global.GVA_CONFIG.Minio.BucketName, global.GVA_CONFIG.Minio.Location)
	//注册路由
	r := router.Setup()

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initialize.InitServer(address, r)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))
	global.GVA_LOG.Error(s.ListenAndServe().Error())

}
