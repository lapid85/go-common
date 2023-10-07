package setting

import (
	"consts/consts"
	"fmt"

	"gorm.io/gorm"
)

// LoadPlatformConfigs 得到平台-站点-配置相关信
func LoadPlatformConfigs(db *gorm.DB) {

	// 初始化
	consts.SitePlatforms = map[string]string{}         // 平台识别号
	consts.SiteNames = map[string]string{}             // 站点名称
	consts.SiteCodes = map[string]string{}             // 站点代码
	consts.SiteConfgs = map[string]map[string]string{} // 站点配置信息
	consts.SiteStaticURLs = map[string]string{}        // 静态文件地址
	consts.SiteUploadURLs = map[string]string{}        // 上传路径
	consts.SiteMysqlStrings = map[string]string{}      // mysql 连接信息
	consts.SitePgSQLStrings = map[string]string{}      // pgsql 连接信息
	consts.SiteRedisStrings = map[string]string{}      // redis 连接信息
	consts.SiteKafkaStrings = map[string]string{}      // kafka 连接信息

	// 平台信息
	var platforms []map[string]string
	db.Table("platforms").Where("status = ?", 2).Find(&platforms)
	if len(platforms) == 0 {
		panic("无法获取平台列表信息")
	}

	// 站点信息
	for _, pv := range platforms {
		var sites []map[string]string
		db.Table("sites").Where("platform_id = ? AND status = ?", pv["id"], 2).Find(&sites)
		if len(sites) == 0 {
			panic(fmt.Sprintf("无法获取平台(id: %v)下属站点信息", pv["id"]))
		}

		for _, sv := range sites {
			var configs []map[string]string
			db.Table("site_configs").Where("platform_id = ? AND site_id = ? AND status = ?", pv["id"], sv["id"], 2).Find(&configs)
			if len(configs) == 0 {
				panic(fmt.Sprintf("无法获取平台(id: %v)/站点(id: %v)配置信息", pv["id"], sv["id"]))
			}

			siteCode := sv["code"] // 站点代码
			cArr := map[string]string{}

			for _, cv := range configs {
				name := cv["name"]
				value := cv["value"]
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

			consts.SiteConfgs[siteCode] = cArr
		}
	}
}
