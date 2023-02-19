package middleware

import (
	"bytes"
	"gin_mall/global"
	"gin_mall/v1/model/response"
	"gin_mall/v1/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"time"
)

func OperationRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body []byte
		var userId int64
		if c.Request.Method != http.MethodGet {
			var err error
			body, err = ioutil.ReadAll(c.Request.Body)
			if err != nil {
				global.GVA_LOG.Error("read body from request error:", zap.Any("err", err))
			} else {
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			}
		}
		UserIdName := global.GVA_CONFIG.JWT.UserIdName
		claims, err := c.Get(UserIdName)
		if err != false {
			userId = claims.(int64)
		} else {
			userId = 0
		}
		obj := c.Request.URL.RequestURI()
		path := obj[9:len(obj)]

		method := c.Request.Method

		name := service.GetRuleNameOperationRecord(path, method)

		record := response.OperationRecord{
			Ip:         c.ClientIP(),
			Method:     method,
			Path:       path,
			RuleName:   name,
			Agent:      c.Request.UserAgent(),
			CreateTime: time.Now().Unix(),
			Body:       string(body),
			ManagerID:  userId,
		}
		// 存在某些未知错误 TODO
		//values := c.Request.Header.Values("content-type")
		//if len(values) >0 && strings.Contains(values[0], "boundary") {
		//	record.Body = "file"
		//}
		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		now := time.Now()

		c.Next()

		latency := time.Now().Sub(now)
		record.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		record.Status = c.Writer.Status()
		record.Latency = latency
		record.Resp = writer.body.String()

		if err := service.CreateOperationRecord(record); err != nil {
			global.GVA_LOG.Error("create operation record error:", zap.Any("err", err))
		}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
