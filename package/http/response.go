package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Code uint

// Success 请求成功
func Success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    StatusOK,
		"message": StatusText(StatusOK),
	})
}

// Alert404Route 路由不存在
func Alert404Route(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    StatusNotFound,
			"message": StatusText(StatusNotFound),
		})
	})
}

// Alert400WithoutMessage 客户端请求错误 不传具体信息
func Alert400WithoutMessage(c *gin.Context, code Code) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    code,
		"message": StatusText(code),
	})
}

// Alert400 客户端请求错误
func Alert400(c *gin.Context, code Code, message string) {
	if message == "" {
		message = StatusText(code)
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    code,
		"message": message,
	})
}

// Alert500WithoutMessage 系统错误 不传具体信息
func Alert500WithoutMessage(c *gin.Context, code Code) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":    code,
		"message": StatusText(code),
	})
}

// Alert500 系统错误
func Alert500(c *gin.Context, code Code, message string) {
	if message == "" {
		message = StatusText(code)
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":    code,
		"message": message,
	})
}

const (
	StatusOK Code = 200
)

const (
	StatusMultipleChoices Code = iota + 300000
)

const (
	StatusBadRequest Code = iota + 400000
	MissUserAgent
	StatusNotFound
)

const (
	StatusInternalServerError Code = iota + 500000
)

var statusText = map[Code]string{
	StatusOK: "请求成功",

	StatusMultipleChoices: "服务器提供可以多个选择的资源",

	StatusBadRequest: "客户端请求错误",
	MissUserAgent:    "请求头部缺少 User-Agent",
	StatusNotFound:   "路由不存在",

	StatusInternalServerError: "系统错误",
}

func StatusText(code Code) string {
	if message, ok := statusText[code]; ok {
		return message
	} else {
		return "未知错误"
	}
}
