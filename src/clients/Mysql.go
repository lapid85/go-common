package clients

import (
	"consts/consts"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQLServers 服务器连接
var MySQLServers = map[string]*gorm.DB{}

// MySQL 获取 mysql 连接
// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
func MySQL(connStr string) (*gorm.DB, error) {
	conf := gorm.Config{}
	return gorm.Open(mysql.Open(connStr), &conf)
}

// MySQLSystem 获取默认的 mysql 连接
func MySQLSystem() (*gorm.DB, error) {
	return MySQL(consts.MySQLCode)
}

// GetMySQLBySite 依据平台获取DB
func GetMySQLBySite(siteCode string) *gorm.DB {
	if siteCode == "" {
		panic("未指定平台名称")
	}
	if val, exists := MySQLServers[siteCode]; exists {
		return val
	}

	if val, exists := consts.SiteMysqlStrings[siteCode]; !exists {
		panic("未找到平台(" + siteCode + ")的数据库连接信息")
	} else {
		db, err := MySQL(val)
		if err != nil {
			return nil
		}
		MySQLServers[siteCode] = db
		return db
	}
}
