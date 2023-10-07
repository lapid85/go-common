package clients

import (
	"consts/consts"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PgSQLServers 服务器连接
var PgSQLServers = map[string]*gorm.DB{}

// PgSQL 获取 postgres 连接
// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
func PgSQL(connStr string) (*gorm.DB, error) {
	conf := gorm.Config{}
	return gorm.Open(postgres.Open(connStr), &conf)
}

// PgSQLSystem 获取默认的 postgres 连接
func PgSQLSystem() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	return PgSQL(dsn)
}

// GetPgSQLBySite 依据平台获取DB
func GetPgSQLBySite(siteCode string) (*gorm.DB, error) {
	if siteCode == "" {
		panic("未指定平台名称")
	}
	if val, exists := PgSQLServers[siteCode]; exists {
		return val, nil
	}

	if val, exists := consts.SitePgSQLStrings[siteCode]; !exists {
		panic("未找到平台(" + siteCode + ")的数据库连接信息")
	} else {
		db, err := PgSQL(val)
		if err != nil {
			return nil, err
		}
		PgSQLServers[siteCode] = db
		return db, nil
	}
}
