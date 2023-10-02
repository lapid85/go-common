package clients

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQL 获取 mysql 连接
func MySQL(connStr string) (*gorm.DB, error) {
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	conf := gorm.Config{}
	return gorm.Open(mysql.Open(connStr), &conf)
}

// MySQLDefault 获取默认的 mysql 连接
func MySQLDefault() (*gorm.DB, error) {
	dsn := "admin:qwe123QWE@tcp(127.0.0.1:3306)/integrated_platforms_v5?charset=utf8mb4"
	return MySQL(dsn)
}

// GetMySQLByPlatform 依据平台获取DB
func GetMySQLByPlatform(platform string) (*gorm.DB, error) {
	return MySQLDefault()
}
