package validator

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strings"
	"zero/package/http"
)

var (
	validate *validator.Validate
)

// init 创建一个验证器实例 初始化翻译器
func init() {
	validate = validator.New()
}

// ValidateStruct 验证结构体
func ValidateStruct(c *gin.Context, s interface{}) {
	err := validate.Struct(s)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			http.Alert400(c, http.StatusBadRequest, fmt.Sprintf("参数验证失败:%s", err))
			return
		}
		for _, e := range err.(validator.ValidationErrors) {
			http.Alert400(c, http.StatusBadRequest,
				fmt.Sprintf("参数 %s 验证失败，错误原因为：%s %s ", strings.ToLower(e.StructNamespace()), e.Tag(), e.Param()))
			return
		}
	}
	return
}

// ValidateStructWithOutContext 验证结构体
func ValidateStructWithOutContext(s interface{}) (success bool, err error) {
	err = validate.Struct(s)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return false, err
		}
		for _, e := range err.(validator.ValidationErrors) {
			return false, errors.New(
				fmt.Sprintf("参数 %s 验证失败，错误原因为：%s %s ",
					strings.ToLower(e.StructNamespace()), e.Tag(), e.Param()))
		}
	}
	return true, nil
}
