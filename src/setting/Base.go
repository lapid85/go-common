package setting

import (
	"common/clients"
	"consts/consts"
	"fmt"
	"os"
	"strings"

	"gopkg.in/ini.v1"
)

// LoadConfigs 加载配置信息
func LoadConfigs(configFile string) {
	// 默认文件名称
	if configFile == "" {
		configFile = "setting.ini"
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

	// 设置运行模式
	if consts.AppName == "" {
		consts.AppName = Get("app_name") //应用程序名称
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

	consts.SiteCodes = map[string]string{}      // 编码 -> 平台识别号初始化
	consts.SiteStaticURLs = map[string]string{} // 静态域名 -> 平台识别号
	consts.SiteUploadURLs = map[string]string{}

	// 加载总的平台信息
	// ----------------------->> 开始 <<----------------------------------------------
	// integratedPlatform := Get("platform.name") // 平台名称
	dbHost := Get("platform.host")         // 主机地址
	dbUser := Get("platform.user")         // 用户名称
	dbPassword := Get("platform.password") // 默认密码
	dbName := Get("platform.database")     // 数据库名
	dbPort := Get("platform.port")         // 默认端口
	connString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8", dbUser, dbPassword, dbHost, dbPort, dbName)
	// dataSources := []string{connString}
	rediString := Get("platform.redis_host") // 默认平台的redis配置信息
	// redis.LoadConfig(integratedPlatform, cacheConnString)
	// db.InitDbServers(integratedPlatform, "mysql", dataSources) //初始化默认的系统平台数据库
	// ----------------------->> 结束 <<----------------------------------------------

	consts.SiteMysqlStrings["system"] = connString // 系统平台的mysql连接信息
	consts.SiteRedisStrings["system"] = rediString // 系统平台的redis连接信息

	platDB, err := clients.MySQL(connString)
	if err != nil {
		panic("无法连接到平台数据库: " + err.Error())
	}
	LoadPlatformConfigs(platDB)
}
