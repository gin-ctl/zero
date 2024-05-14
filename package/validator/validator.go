package validator

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

var (
	validate *validator.Validate
)

type Param struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Rule  string `json:"rule"`
}

// init 创建一个验证器实例 初始化翻译器
func init() {
	validate = validator.New()
}

// ValidateStruct 验证结构体
func ValidateStruct(s interface{}) bool {
	err := validate.Struct(s)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Sprintf("参数验证失败:%s", err)
			//response.ValidationError(c, fmt.Sprintf("参数验证失败:%s", err))
			return false
		}
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Sprintf("参数 %s 验证失败，错误原因为：%s %s ", strings.ToLower(e.StructNamespace()), e.Tag(), e.Param())
			//response.ValidationError(
			//	c,
			//	fmt.Sprintf("参数 %s 验证失败，错误原因为：%s %s ", strings.ToLower(e.StructNamespace()), e.Tag(), e.Param()),
			//)
			return false
		}
	}
	return true
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

// ValidateVariable 验证变量
func ValidateVariable(params []Param) bool {
	for _, param := range params {
		err := validate.Var(param.Value, param.Rule)
		if err != nil {
			fmt.Sprintf("参数 %s 验证失败，错误原因为：%s", param.Field, err)
			//response.ValidationError(c, fmt.Sprintf("参数 %s 验证失败，错误原因为：%s", param.Field, err))
			return false
		}
	}
	return true
}
