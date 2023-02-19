gin_mall 用gin框架实现商城后台


线上地址

https://mallweb.acorn.eu.org/login

默认用户名密码 admin admin


线上api接口地址
https://mall.acorn.eu.org/login

postman 接口文件还在整理



配置文件
[config.yaml](config.yaml)

数据文件解压后 自行用工具导入

[gin_mall.sql.zip](gin_mall.sql.zip)
[gin_mall.sql](gin_mall.sql)

项目展示图片
![Screenshot 2023-02-19 at 17.31.22.png](demo_image%2FScreenshot%202023-02-19%20at%2017.31.22.png)
![Screenshot 2023-02-19 at 17.30.36.png](demo_image%2FScreenshot%202023-02-19%20at%2017.30.36.png)
![Screenshot 2023-02-19 at 17.29.29.png](demo_image%2FScreenshot%202023-02-19%20at%2017.29.29.png)




```
.
├── README.md
├── config
│   ├── captcha.go
│   ├── casbin.go
│   ├── config.go
│   ├── email.go
│   ├── excel.go
│   ├── jwt.go
│   ├── minio.go
│   ├── oss.go
│   ├── redis.go
│   ├── snowflake.go
│   ├── sqlx.go
│   ├── system.go
│   └── zap.go
├── config.yaml
├── daemon.log
├── daemon.txt
├── global
│   ├── code.go
│   ├── error_code.go
│   └── global.go
├── go.mod
├── go.sum
├── initialize
│   ├── gorm.go
│   ├── internal
│   │   └── logger.go
│   ├── jwt.go
│   ├── minio.go
│   ├── redis.go
│   ├── server_other.go
│   ├── snowflake.go
│   ├── viper.go
│   └── zap.go
├── latest_log -> log/2023-02-19.log
├── log
│   ├── 2023-02-12.log
│   ├── 2023-02-13.log
│   ├── 2023-02-14.log
│   ├── 2023-02-15.log
│   ├── 2023-02-16.log
│   ├── 2023-02-17.log
│   ├── 2023-02-18.log
│   └── 2023-02-19.log
├── main.go
├── mall-linux
├── mall_server
├── middleware
│   ├── cors.go
│   ├── error.go
│   ├── jwt.go
│   ├── logger.go
│   ├── operation.go
│   ├── rbac.go
│   └── sentinel.go
├── router
│   ├── router_base.go
│   └── v1
│       ├── goods_skus_card.go
│       ├── router_agent.go
│       ├── router_api.go
│       ├── router_app_category_item.go
│       ├── router_category.go
│       ├── router_coupon.go
│       ├── router_dashboard.go
│       ├── router_goods.go
│       ├── router_goods_comment.go
│       ├── router_goods_skus_card_value.go
│       ├── router_image.go
│       ├── router_image_class.go
│       ├── router_manager.go
│       ├── router_menu.go
│       ├── router_notice.go
│       ├── router_operation_record.go
│       ├── router_order.go
│       ├── router_public.go
│       ├── router_role.go
│       ├── router_skus.go
│       ├── router_system.go
│       ├── router_user.go
│       └── router_user_level.go
├── utils
│   ├── common.go
│   ├── constant.go
│   ├── directory.go
│   ├── md5.go
│   ├── minio.go
│   ├── request.go
│   ├── response.go
│   ├── rotatelogs_unix.go
│   └── server.go
├── v1
│   ├── api
│   │   ├── agent.go
│   │   ├── api.go
│   │   ├── app_category_item.go
│   │   ├── captcha.go
│   │   ├── category.go
│   │   ├── coupon.go
│   │   ├── dashboard.go
│   │   ├── goods.go
│   │   ├── goods_comment.go
│   │   ├── goods_skus_card.go
│   │   ├── goods_skus_card_value.go
│   │   ├── image.go
│   │   ├── image_class.go
│   │   ├── manager.go
│   │   ├── menu.go
│   │   ├── notice.go
│   │   ├── operation_record.go
│   │   ├── order.go
│   │   ├── role.go
│   │   ├── skus.go
│   │   ├── system.go
│   │   ├── user.go
│   │   └── user_level.go
│   ├── model
│   │   ├── app_category_item.go
│   │   ├── category.go
│   │   ├── coupon.go
│   │   ├── goods.go
│   │   ├── goods_banner.go
│   │   ├── goods_comment.go
│   │   ├── goods_skus.go
│   │   ├── goods_skus_card.go
│   │   ├── goods_skus_card_value.go
│   │   ├── image.go
│   │   ├── image_class.go
│   │   ├── manager.go
│   │   ├── notice.go
│   │   ├── operation_record.go
│   │   ├── order.go
│   │   ├── request
│   │   │   ├── agent.go
│   │   │   ├── api.go
│   │   │   ├── app_category_item.go
│   │   │   ├── category.go
│   │   │   ├── common.go
│   │   │   ├── coupon.go
│   │   │   ├── dashboard.go
│   │   │   ├── goods.go
│   │   │   ├── goods_banner.go
│   │   │   ├── goods_comment.go
│   │   │   ├── goods_skus.go
│   │   │   ├── goods_skus_card.go
│   │   │   ├── goods_skus_card_value.go
│   │   │   ├── image.go
│   │   │   ├── image_class.go
│   │   │   ├── manager.go
│   │   │   ├── menu.go
│   │   │   ├── notice.go
│   │   │   ├── operation_record.go
│   │   │   ├── order.go
│   │   │   ├── role.go
│   │   │   ├── skus.go
│   │   │   ├── user.go
│   │   │   └── user_level.go
│   │   ├── response
│   │   │   ├── agent.go
│   │   │   ├── api.go
│   │   │   ├── app_category_item.go
│   │   │   ├── captcha.go
│   │   │   ├── category.go
│   │   │   ├── common.go
│   │   │   ├── coupon.go
│   │   │   ├── dashboard.go
│   │   │   ├── goods.go
│   │   │   ├── goods_banner.go
│   │   │   ├── goods_comment.go
│   │   │   ├── goods_skus.go
│   │   │   ├── goods_skus_card.go
│   │   │   ├── goods_skus_card_value.go
│   │   │   ├── image.go
│   │   │   ├── image_class.go
│   │   │   ├── manager.go
│   │   │   ├── menu.go
│   │   │   ├── notice.go
│   │   │   ├── operation_record.go
│   │   │   ├── order.go
│   │   │   ├── role.go
│   │   │   ├── role_rule.go
│   │   │   ├── skus.go
│   │   │   ├── user.go
│   │   │   ├── user_info.go
│   │   │   └── user_level.go
│   │   ├── role.go
│   │   ├── role_rule.go
│   │   ├── rule.go
│   │   ├── skus.go
│   │   ├── user.go
│   │   ├── user_bill.go
│   │   ├── user_extract.go
│   │   ├── user_info.go
│   │   └── user_level.go
│   └── service
│       ├── agent.go
│       ├── api.go
│       ├── app_category_item.go
│       ├── category.go
│       ├── coupon.go
│       ├── dashboard.go
│       ├── goods.go
│       ├── goods_comment.go
│       ├── goods_skus_card.go
│       ├── goods_skus_card_value.go
│       ├── image.go
│       ├── image_class.go
│       ├── jwt.go
│       ├── manager.go
│       ├── menu.go
│       ├── notice.go
│       ├── operation_record.go
│       ├── order.go
│       ├── role.go
│       ├── skus.go
│       ├── system.go
│       ├── user.go
│       └── user_level.go
└── validator
    ├── login_validator.go
    └── validator.go

16 directories, 206 files


```