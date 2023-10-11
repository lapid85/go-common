package types

import (
	"errors"
	"gorm.io/gorm"
)

// IsNotFound 判断是否为未找到
func IsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
