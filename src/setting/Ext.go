package setting

import (
	"consts/consts"
	"fmt"
	"os"
	"strings"

	"gopkg.in/ini.v1"
)

// Ext 扩展配置信息
var Ext *ini.File

// LoadExtConfigs 加载扩展配置信息
func LoadExtConfigs(args ...string) {
	configFile := "setting_ext.ini" // 默认文件名称
	if len(args) >= 1 {
		configFile = args[0]
	}
	cfg, err := ini.Load(configFile) //加载配置文件
	if err != nil {
		fmt.Printf("读取配置文件出错: %v\n", err)
		os.Exit(1)

	}
	Ext = cfg //设置全局Ini信息

	appName := GetExt("app_name")
	if appName != "" {
		consts.AppName = appName //应用程序名称
	}

	runMode := GetExt("sys.run_mode")
	if runMode != "" {
		consts.RunMode = runMode //运行模式
	}

}

// 依据key获得ini的设置项
func getExtIniKey(key string) *ini.Key {
	sArr := strings.Split(key, ".")
	arrLen := len(sArr)
	if arrLen == 1 {
		return Ext.Section("").Key(key)
	} else if arrLen == 2 {
		return Ext.Section(sArr[0]).Key(sArr[1])
	}
	return nil
}

// GetExt 获取扩展配置项信息, 注意: 此方法必须在 LoadExtConfigs() 方法之后执行, 否则将无法获取配置信息
func GetExt(key string, args ...interface{}) string {
	defaultValue := ""
	if len(args) >= 1 {
		defaultValue = args[0].(string)
	}
	if Ext == nil {
		return defaultValue
	}

	val := getExtIniKey(key).String()
	if val == "" {
		return defaultValue
	}
	return val
}

// GetExtInt 得到Int的值
func GetExtInt(key string) int {
	if Ext == nil {
		return 0
	}
	v, err := getIniKey(key).Int()
	if err != nil {
		return 0
	}
	return v
}

// GetExtBool 得到true/false的值
func GetExtBool(key string) bool {
	if Ext == nil {
		return false
	}
	v, err := getExtIniKey(key).Bool()
	if err != nil {
		return false
	}
	return v
}

// GetExtFloat64 得到float64的值
func GetExtFloat64(key string) float64 {
	if Ext == nil {
		return 0
	}
	v, err := getExtIniKey(key).Float64()
	if err != nil {
		return 0
	}
	return v
}
