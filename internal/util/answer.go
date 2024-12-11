package util

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// 通用响应结构
type Resp[T any] struct {
	Code    int    `json:"code"`    // 状态码
	Message string `json:"message"` // 信息
	Data    T      `json:"data"`    // 数据
}

// 成功响应
func SuccessResp(c *gin.Context, data ...interface{}) {
	resp := Resp[interface{}]{
		Code:    200,
		Message: "success",
		Data:    nil,
	}
	if len(data) > 0 {
		resp.Data = data[0]
	}
	c.JSON(200, resp)
}

// 错误响应
func ErrorResp(c *gin.Context, err error, code int, logError bool) {
	msg := "unknown error"
	if err != nil {
		msg = err.Error()
	}
	if logError {
		log.Errorf("Error %d: %s", code, msg)
	}
	c.JSON(200, Resp[interface{}]{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}

// 自定义消息的错误响应
func ErrorStrResp(c *gin.Context, msg string, code int, logError bool) {
	if logError {
		log.Errorf("Error %d: %s", code, msg)
	}
	c.JSON(200, Resp[interface{}]{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}
