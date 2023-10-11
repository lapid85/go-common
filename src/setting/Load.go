package setting

import (
	"consts/consts"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// LoadPlatformConfigs 得到平台-站点-配置相关信
func LoadPlatformConfigs(db *gorm.DB) {

	// 平台信息
	var platforms []map[string]interface{}
	db.Table("platforms").Where("status = ?", 2).Find(&platforms)
	if len(platforms) == 0 {
		panic("无法获取平台列表信息")
	}

	// 站点信息
	for _, pv := range platforms {
		var sites []map[string]interface{}
		if result := db.Table("sites").Where("platform_id = ? AND status = ?", pv["id"], 2).Find(&sites); result.Error != nil {
			panic(fmt.Sprintf("无法获取平台(id: %v)下属站点信息: %v", pv["id"], result.Error))
		} else if len(sites) == 0 {
			panic(fmt.Sprintf("无法获取平台(id: %v)下属站点信息", pv["id"]))
		}

		for _, sv := range sites {
			var configs []map[string]interface{}
			if result := db.Table("site_configs").Where("platform_id = ? AND site_id = ? AND status = ?", pv["id"], sv["id"], 2).Find(&configs); result.Error != nil {
				panic(fmt.Sprintf("无法获取平台(id: %v)/站点(id: %v)配置信息: %v", pv["id"], sv["id"], result.Error))
			} else if len(configs) == 0 {
				panic(fmt.Sprintf("无法获取平台(id: %v)/站点(id: %v)配置信息", pv["id"], sv["id"]))
			}

			siteCode := strings.ToUpper(sv["code"].(string)) // 站点代码
			cArr := map[string]string{}

			consts.SiteCodes[siteCode] = siteCode // 站点代码

			for _, cv := range configs {
				name := cv["name"].(string)
				value := cv["value"].(string)
				cArr[name] = value
				if name == "platform" {
					consts.SitePlatforms[siteCode] = value // 平台识别号
				} else if name == "static_url" {
					consts.SiteStaticURLs[siteCode] = value // 静态文件地址
				} else if name == "upload_url" {
					consts.SiteUploadURLs[siteCode] = value // 上传路径
				} else if name == "conn_strings" {
					consts.SiteMysqlStrings[siteCode] = value // mysql 连接信息
				} else if name == "pgsql_strings" {
					consts.SitePgSQLStrings[siteCode] = value // pgsql 连接信息
				} else if name == "redis_strings" {
					consts.SiteRedisStrings[siteCode] = value // redis 连接信息
				} else if name == "kafka_strings" {
					consts.SiteKafkaStrings[siteCode] = value // kafka 连接信息
				} else if name == "site_name" {
					consts.SiteNames[siteCode] = value // 站点名称
				}
			}

			consts.SiteConfigs[siteCode] = cArr
		}
	}
}
