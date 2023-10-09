package setting

import (
	"strings"

	"gopkg.in/ini.v1"
)

// Ini 配置信息
var Ini *ini.File

// 依据key获得ini的设置项
func getIniKey(key string) *ini.Key {
	sArr := strings.Split(key, ".")
	arrLen := len(sArr)
	if arrLen == 1 {
		return Ini.Section("").Key(key)
	} else if arrLen == 2 {
		return Ini.Section(sArr[0]).Key(sArr[1])
	}
	return nil
}

// Get 获取配置项信息, 注意: 此方法必须在 LoadConfigs() 方法之后执行, 否则将无法获取配置信息
// key: 配置项名称
// 顶级配置项获取: config.Get("app_name")
// 二级配置项获取: config.Get("database.name")
func Get(key string, args ...interface{}) string {
	defaultValue := ""
	if len(args) >= 1 {
		defaultValue = args[0].(string)
	}
	if Ini == nil {
		return defaultValue
	}

	val := getIniKey(key).String()
	if val == "" {
		return defaultValue
	}
	return val
}

// GetInt 得到Int的值
// 请依据自己需要自动对int进行转换
func GetInt(key string) int {
	if Ini == nil {
		return 0
	}
	v, err := getIniKey(key).Int()
	if err != nil {
		return 0
	}
	return v
}

// GetBool 得到true/false的值
func GetBool(key string) bool {
	if Ini == nil {
		return false
	}
	v, err := getIniKey(key).Bool()
	if err != nil {
		return false
	}
	return v
}

// GetFloat 得到float的值
func GetFloat(key string) float64 {
	if Ini == nil {
		return 0
	}
	v, err := getIniKey(key).Float64()
	if err != nil {
		return 0
	}
	return v
}
