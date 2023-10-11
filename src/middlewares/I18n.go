package middlewares

import (
	"common/log"
	"encoding/json"
	"os"
	"strings"

	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

// LangPathDefault 默认语言文件路径
const LangPathDefault = "./locales"

// LangPathCommon 公共语言文件路径
const LangPathCommon = "common"

// LangDefault 默认语言
var LangDefault = language.SimplifiedChinese

// LangLoader 语言加载器
type LangLoader struct{}

// LoadLangDir 加载语言文件 - 读取目录下所有文件
func (ths *LangLoader) LoadLangDir(dirPath string, langName string) (map[string]string, error) {
	pathArr := strings.Split(dirPath, "/")
	if len(pathArr) == 0 {
		log.Error("语言文件路径错误")
		return nil, nil
	}
	// 读取目录下所有文件
	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	// 语言文件内容
	result := map[string]string{}
	for _, file := range files {
		filePath := dirPath + "/" + file.Name()
		if file.IsDir() {
			jsonMap, err := ths.LoadLangDir(filePath, langName)
			if err != nil {
				log.Error(err.Error())
				return nil, err
			}
			for k, v := range jsonMap {
				result[k] = v
			}
			continue
		}

		// 如果不是 json 文件, 则跳过
		jsonLangFile := langName + ".json"
		if strings.LastIndex(file.Name(), jsonLangFile) != -1 {
			continue
		}

		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}

		var jsonMap map[string]string
		if err := json.Unmarshal(fileContent, &jsonMap); err != nil {
			log.Error(err.Error())
			return nil, err
		}
		for k, v := range jsonMap {
			result[k] = v
		}
	}

	return result, nil
}

// LoadMessage 加载语言文件 - 读取目录下所有文件 - 加载默认通用语言文件时, 会加载所有语言文件
func (ths *LangLoader) LoadMessage(path string) ([]byte, error) {
	pathArr := strings.Split(path, "/")
	if len(pathArr) == 0 {
		log.Error("语言文件路径错误: %s", path)
		return nil, nil
	}

	// 语言文件内容
	fileContent, err := os.ReadFile(path)
	if err != nil {
		log.Error("读取文件内容错误 %v %s", path, err.Error())
		return nil, err
	}

	resultMap := map[string]string{}
	if err := json.Unmarshal(fileContent, &resultMap); err != nil {
		log.Error("解析文件内容错误 %v %s", path, err.Error())
		return nil, err
	}

	// 所有相同语言文件内容
	langArr := strings.Split(pathArr[len(pathArr)-1], ".")
	if len(langArr) == 0 {
		log.Error("具体语言文件路径错误: %s", path)
		return nil, nil
	}
	langName := langArr[0]           // 当前加载语言文件名称
	allLangPath := "./" + pathArr[0] // 所有语言文件路径
	dirs, err := os.ReadDir(allLangPath)
	if err != nil {
		log.Error("读取所有语言文件目录错误: %s", err.Error())
		return nil, err
	}

	for _, dir := range dirs {
		dirPath := allLangPath + "/" + dir.Name()
		if stat, err := os.Stat(dirPath); err != nil {
			panic(err)
		} else if !stat.IsDir() {
			continue
		}
		if jsonMap, err := ths.LoadLangDir(dirPath, langName); err != nil {
			log.Error("加载所有目录内语言文件出错: %s", err.Error())
			return nil, err
		} else {
			for k, v := range jsonMap {
				if _, ok := resultMap[k]; !ok {
					resultMap[k] = v
				}
			}
		}
	}

	// _, _ = pp.Println(resultMap)
	return json.Marshal(resultMap)
}

// I18nWithLangHandler 国际化
func I18nWithLangHandler() gin.HandlerFunc {
	return ginI18n.Localize(
		ginI18n.WithGetLngHandle(
			func(context *gin.Context, defaultLng string) string {
				lng := context.Query("lng")
				if lng == "" {
					return defaultLng
				}
				return lng
			},
		),
	)
}

// I18n 国际化 - 参数: 语言文件路径, 语言列表, 默认语言
func I18n(args ...interface{}) gin.HandlerFunc {
	// 默认语言
	langPathDefault := LangPathDefault
	if len(args) > 0 {
		langPathDefault = args[0].(string)
	}

	// 语言列表
	languages := []language.Tag{language.English, language.SimplifiedChinese, language.TraditionalChinese, language.BrazilianPortuguese}
	if len(args) > 1 {
		languages = args[1].([]language.Tag)
	}

	// 默认语言
	langDefault := LangDefault
	if len(args) > 2 {
		langDefault = args[2].(language.Tag)
	}

	return ginI18n.Localize(ginI18n.WithBundle(&ginI18n.BundleCfg{
		RootPath:         langPathDefault,
		AcceptLanguage:   languages,
		DefaultLanguage:  langDefault,
		UnmarshalFunc:    json.Unmarshal,
		FormatBundleFile: "json",
		Loader:           &LangLoader{},
	}))
}
