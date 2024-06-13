package logic

import (
	"github.com/gin-ctl/zero/package/http"
	"github.com/gin-gonic/gin"
)

type Basic interface {
	ParseAndCheckParams(c *gin.Context) (err error)
}

type BasicRequest[T any] interface {
	*T
	Basic
}

// ParseAndCheckParams Parse and check params.
func ParseAndCheckParams[T any, P BasicRequest[T]](c *gin.Context) (params http.RequestType[T], err error) {
	var v T
	err = P(&v).ParseAndCheckParams(c)
	if err != nil {
		return http.RequestType[T]{}, err
	}
	return http.NewRequestType(v), nil
}