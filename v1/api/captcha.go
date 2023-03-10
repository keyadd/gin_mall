package api

import (
	"gin_mall/global"
	"gin_mall/utils"
	"gin_mall/v1/model/response"

	//"gin-vue-admin/global"
	//"gin-vue-admin/model/response"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

var store = base64Captcha.DefaultMemStore

// @Tags Base
// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证码获取成功"}"
// @Router /base/captcha [post]
func Captcha(c *gin.Context) {
	//字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.GVA_CONFIG.Captcha.ImgHeight, global.GVA_CONFIG.Captcha.ImgWidth, global.GVA_CONFIG.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		global.GVA_LOG.Error("验证码获取失败!", zap.Any("err", err))
		return
		//response.FailWithMessage("验证码获取失败", c)
	} else {
		captchaResponse := response.CaptchaResponse{
			CaptchaId: id,
			PicPath:   b64s,
		}
		utils.ResponseSuccess(c, captchaResponse)

	}
}
