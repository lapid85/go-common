package model

import "gorm.io/gorm"

type IModel interface {
	GetAll(*gorm.DB, ...interface{}) (interface{}, int64, error)
	Delete(*gorm.DB, map[string]interface{}) (interface{}, error)
}
