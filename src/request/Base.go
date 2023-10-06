package request

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// PageSize 默认分页大小
const PageSize = 15

// IsAjax 是否是ajax请求
func IsAjax(c *gin.Context) bool {
	return strings.ToLower(c.Request.Header.Get("X-Requested-With")) == "xmlhttprequest"
}

// GetPage 获取页码
func GetPage(c *gin.Context) int {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		return 1
	}
	if page == 0 {
		return 1
	}
	return page
}

// GetLimit 获取limit信息
func GetLimit(c *gin.Context) int {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "15"))
	if err != nil {
		return PageSize
	}
	if limit == 0 || limit > 10000 {
		return PageSize
	}
	return limit
}

// GetOffset 获取 (limit, offset)
func GetOffset(c *gin.Context) (int, int) {
	page := GetPage(c)
	size := GetLimit(c)
	return size, (page - 1) * size
}

// GetPlatform 获取平台信息
func GetPlatform(c *gin.Context) string {
	return c.DefaultQuery("platform", "default")
}
