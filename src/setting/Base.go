package setting

import (
	"common/clients"
	"consts/consts"
	"fmt"
	"github.com/k0kubun/pp/v3"
	"os"
	"strings"

	"gopkg.in/ini.v1"
)

// LoadConfigs 加载配置信息
func LoadConfigs(args ...string) {
	configFile := "setting.ini" // 默认文件名称
	if len(args) >= 1 {
		configFile = args[0]
	}
	cfg, err := ini.Load(configFile) //加载配置文件
	if err != nil {
		fmt.Printf("读取配置文件出错: %v\n", err)
		os.Exit(1)

	}
	Ini = cfg //设置全局Ini信息

	// ---------------------------------------------------------------------
	// 注意: 需要加什么额外的ini配置信息, 可以追加到这里面
	// 此段区域为除数据库/redis/mango之外的ini配置信息
	// 开始 ->
	// ---------------------------------------------------------------------
	consts.UploadPath = Get("upload.save_path")   //上传保存文件位置
	consts.UploadURLPath = Get("upload.url_path") //上传文件url前缀
	consts.IpDbPath = Get("ip.db_path")           //IP数据库保存路径
	consts.LogPath = Get("log.path")

	//-----------加载内网虚拟域名的配置------------
	consts.InternalMemberServUrl = Get("internal.internal_member_service")
	consts.InternalGameServUrl = Get("internal.internal_game_service")
	consts.InternalOssServUrl = Get("internal.internal_oss_service")
	consts.InternalAdminServUrl = Get("internal.internal_admin_service")

	//-----------加载kafka配置信息------------
	// consts.KafkaBrokerList = strings.Split(Get("kafka.broker_list_str"), ",")
	// fmt.Printf("KafkaBrokerList: %s\n", consts.KafkaBrokerList)
	// consts.KafkaTopicList = strings.Split(Get("kafka.topic_list_str"), ",")
	// consts.KafkaVersion = Get("kafka.version")

	// <- 结束
	// ---------------------------------------------------------------------

	// APP名称
	appName := Get("app_name")
	if appName != "" {
		consts.AppName = appName //应用程序名称
	}
	// 设置运行模式
	runMode := Get("sys.run_mode")
	if runMode != "" {
		consts.RunMode = runMode //运行模式
	}
	consts.PlatformIntegrated = Get("platform_integrated") //综合平台名称 - 总的平台名称
	consts.SiteDefault = Get("platform_default")           //设定默认平台 - 默认平台名称
	// 内网ip列表 - 内网白名单
	internalIpListStr := Get("platform.internal_ip_list")
	if internalIpListStr != "" {
		if strings.Index(internalIpListStr, ",") > 0 {
			consts.InternalIpList = strings.Split(",", internalIpListStr)
		} else {
			consts.InternalIpList = append(consts.InternalIpList, internalIpListStr)
		}
	}

	// 加载总的平台信息
	// ----------------------->> 开始 <<----------------------------------------------
	// integratedPlatform := Get("platform.name") // 平台名称
	dbHost := Get("platform.host")         // 主机地址
	dbUser := Get("platform.user")         // 用户名称
	dbPassword := Get("platform.password") // 默认密码
	dbName := Get("platform.database")     // 数据库名
	dbPort := Get("platform.port")         // 默认端口
	connString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8", dbUser, dbPassword, dbHost, dbPort, dbName)
	rediString := Get("platform.redis_host") // 默认平台的redis配置信息
	// ----------------------->> 结束 <<----------------------------------------------

	consts.SiteMysqlStrings["system"] = connString // 系统平台的mysql连接信息
	consts.SiteRedisStrings["system"] = rediString // 系统平台的redis连接信息

	// 加载平台配置信息
	platDB, err := clients.MySQL(connString)
	fmt.Println(connString)
	if err != nil {
		panic("无法连接到平台数据库: " + err.Error())
	}

	// 初始化
	consts.SitePlatforms = map[string]string{}          // 平台识别号
	consts.SiteNames = map[string]string{}              // 站点名称
	consts.SiteCodes = map[string]string{}              // 站点代码
	consts.SiteConfigs = map[string]map[string]string{} // 站点配置信息
	consts.SiteStaticURLs = map[string]string{}         // 静态文件地址
	consts.SiteUploadURLs = map[string]string{}         // 上传路径
	consts.SiteMysqlStrings = map[string]string{}       // mysql 连接信息
	consts.SitePgSQLStrings = map[string]string{}       // pgsql 连接信息
	consts.SiteRedisStrings = map[string]string{}       // redis 连接信息
	consts.SiteKafkaStrings = map[string]string{}       // kafka 连接信息

	LoadPlatformConfigs(platDB)
}

// PrintConfigs 打印配置信息
func PrintConfigs() {
	loadedData := map[string]interface{}{
		"AppName":               consts.AppName,               // 应用程序名称
		"RunMode":               consts.RunMode,               // 运行模式
		"UploadPath":            consts.UploadPath,            // 上传保存文件位置
		"UploadURL":             consts.UploadURLPath,         // 上传文件url前缀
		"LogPath":               consts.LogPath,               // 日志文件保存位置
		"IpDbPath":              consts.IpDbPath,              // IP数据库保存路径
		"InternalMemberServUrl": consts.InternalMemberServUrl, // 内网虚拟域名 - 会员服务
		"InternalGameServUrl":   consts.InternalGameServUrl,   // 内网虚拟域名 - 游戏服务
		"InternalOssServUrl":    consts.InternalOssServUrl,    // 内网虚拟域名 - oss服务
		"InternalAdminServUrl":  consts.InternalAdminServUrl,  // 内网虚拟域名 - 后台服务
		"PlatformIntegrated":    consts.PlatformIntegrated,    // 综合平台名称 - 总的平台名称
		"SiteDefault":           consts.SiteDefault,           // 设定默认平台 - 默认平台名称
		"InternalIpList":        consts.InternalIpList,        // 内网ip列表 - 内网白名单
		"SitePlatforms":         consts.SitePlatforms,         // 平台识别号
		"SiteNames":             consts.SiteNames,             // 站点名称
		"SiteCodes":             consts.SiteCodes,             // 站点代码
		"SiteConfigs":           consts.SiteConfigs,           // 站点配置信息
		"SiteStaticURLs":        consts.SiteStaticURLs,        // 静态文件地址
		"SiteUploadURLs":        consts.SiteUploadURLs,        // 上传路径
		"SiteMysqlStrings":      consts.SiteMysqlStrings,      // mysql 连接信息
		"SitePgSQLStrings":      consts.SitePgSQLStrings,      // pgsql 连接信息
		"SiteRedisStrings":      consts.SiteRedisStrings,      // redis 连接信息
		"SiteKafkaStrings":      consts.SiteKafkaStrings,      // kafka 连接信息
	}

	_, _ = pp.Println(loadedData)
}
