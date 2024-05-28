package get

import (
	"fmt"
	"github.com/gin-ctl/zero/package/helper"
	"github.com/spf13/cast"
	lib "github.com/spf13/viper"
)

var viper *lib.Viper

func NewViper(filename, path string) {
	viper = lib.New()
	viper.SetConfigName(filename)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.WatchConfig()
}

func internalGet(path string, defaultValue ...interface{}) interface{} {
	// config 或者环境变量不存在的情况
	if !viper.IsSet(path) || helper.Empty(viper.Get(path)) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return viper.Get(path)
}

func Get(path string, defaultValue ...interface{}) string {
	return String(path, defaultValue...)
}

// String 获取 String 类型的配置信息
func String(path string, defaultValue ...interface{}) string {
	return cast.ToString(internalGet(path, defaultValue...))
}

// Int 获取 Int 类型的配置信息
func Int(path string, defaultValue ...interface{}) int {
	return cast.ToInt(internalGet(path, defaultValue))
}

// Float64 获取 float64 类型的配置信息
func Float64(path string, defaultValue ...interface{}) float64 {
	return cast.ToFloat64(internalGet(path, defaultValue...))
}

// Int64 获取 Int64 类型的配置信息
func Int64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(internalGet(path, defaultValue...))
}

// Uint 获取 Uint 类型的配置信息
func Uint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(internalGet(path, defaultValue...))
}

// Bool 获取 Bool 类型的配置信息
func Bool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(internalGet(path, defaultValue...))
}

// StringMapString 获取结构数据
func StringMapString(path string) map[string]string {
	return viper.GetStringMapString(path)
}

// StringSlice 获取结构数据
func StringSlice(path string) []string {
	return viper.GetStringSlice(path)
}
