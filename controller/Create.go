package controller

import (
	"common/clients"
	"common/model"
	"common/request"
	"common/response"

	"github.com/gin-gonic/gin"
)

// ActionCreate 创建
type ActionCreate struct {
	Model model.IModel
}

// Create 创建 - post
func (ths *ActionCreate) Create(c *gin.Context) {
	// 过滤字段
	requestData := c.Request.PostForm    // 获取请求数据
	tableFields := ths.Model.Fields()    // 获取表字段
	realData := map[string]interface{}{} // 过滤后的数据
	for _, field := range tableFields {
		if val, exists := requestData[field]; exists {
			realData[field] = val
		}
	}

	platform := request.GetPlatform(c)              // 获取平台
	db, err := clients.GetMySQLByPlatform(platform) // 获取数据库连接
	if err != nil {
		response.Err(c, "获取DB失败")
		return
	}

	if _, err := ths.Model.Create(db, realData); err != nil {
		response.Err(c, "创建失败")
		return
	}

	response.Ok(c)
}
