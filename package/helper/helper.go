package helper

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

// Empty 类似于 PHP 的 empty() 函数
func Empty(val interface{}) bool {
	if val == nil {
		return true
	}
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return reflect.DeepEqual(val, reflect.Zero(v.Type()).Interface())
}

// MicrosecondsStr 将时间转换为毫秒字符串
func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}

// AssignStruct 将 src 的值赋值给 dst，src和dst必须是指针类型
func AssignStruct(dst, src interface{}) error {
	dstVal := reflect.ValueOf(dst)
	srcVal := reflect.ValueOf(src)

	if dstVal.Kind() != reflect.Ptr || dstVal.Elem().Kind() != reflect.Struct {
		return errors.New("dst must be a pointer to a struct")
	}

	if srcVal.Kind() != reflect.Ptr || srcVal.Elem().Kind() != reflect.Struct {
		return errors.New("src must be a pointer to a struct")
	}

	dstVal = dstVal.Elem()
	srcVal = srcVal.Elem()

	for i := 0; i < dstVal.NumField(); i++ {
		fieldName := dstVal.Type().Field(i).Name
		srcFieldVal := srcVal.FieldByName(fieldName)
		if srcFieldVal.IsValid() &&
			!reflect.DeepEqual(srcFieldVal.Interface(), reflect.Zero(srcFieldVal.Type()).Interface()) {
			dstVal.Field(i).Set(srcFieldVal)
		}
	}

	return nil
}
