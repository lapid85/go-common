package clients

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PgSQL 获取 postgres 连接
func PgSQL(connStr string) (*gorm.DB, error) {
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	conf := gorm.Config{}
	return gorm.Open(postgres.Open(connStr), &conf)

}
