package controller

import (
	"common/clients"
	"common/model"
	"common/request"
	"common/response"
	"common/trans"

	"github.com/gin-gonic/gin"
)

// ActionCreate 创建
type ActionCreate struct {
	Model        model.IModel                                     // model - 必须
	CreateBefore func(*gin.Context, *map[string]interface{}) bool // 默认处理函数
	CreateAfter  func(*gin.Context, interface{})                  // 默认处理函数
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

	// 添加数据之前的处理
	if ths.CreateBefore != nil {
		if !ths.CreateBefore(c, &realData) {
			response.Err(c, trans.Tr(c, "errCreateData"))
			return
		}
	}

	siteCode := request.GetSiteCode(c)          // 获取平台
	db, err := clients.GetMySQLBySite(siteCode) // 获取数据库连接
	if err != nil {
		response.Err(c, trans.Tr(c, "errGetDbConn"))
		return
	}

	if row, err := ths.Model.Create(db, realData); err != nil {
		response.Err(c, trans.Tr(c, "errCreateData"))
		return
	} else if ths.CreateAfter != nil {
		ths.CreateAfter(c, row)
	}

	response.Ok(c)
}
