package validator

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strings"
)

var (
	validate *validator.Validate
)

// init 创建一个验证器实例 初始化翻译器
func init() {
	validate = validator.New()
}

// ValidateStruct 验证结构体
func ValidateStruct(c *gin.Context, s interface{}) (err error) {
	err = validate.Struct(s)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return errors.New(fmt.Sprintf("参数验证失败:%s", err))
		}
		for _, e := range err.(validator.ValidationErrors) {
			return errors.New(fmt.Sprintf("参数 %s 验证失败，错误原因为：%s %s ",
				strings.ToLower(e.StructNamespace()), e.Tag(), e.Param()))
		}
	}
	return
}

// ValidateStructWithOutCtx 验证结构体
func ValidateStructWithOutCtx(s interface{}) (err error) {
	err = validate.Struct(s)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return
		}
		for _, e := range err.(validator.ValidationErrors) {
			return errors.New(
				fmt.Sprintf("参数 %s 验证失败，错误原因为：%s %s ",
					strings.ToLower(e.StructNamespace()), e.Tag(), e.Param()))
		}
	}
	return
}
