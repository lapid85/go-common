package controller

import (
	"common/clients"
	"common/model"
	"common/request"
	"common/response"
	"common/trans"

	"github.com/gin-gonic/gin"
)

// ActionUpdate 更新
type ActionUpdate struct {
	Model model.IModel
}

// Update 更新 - put
func (ths *ActionUpdate) Update(c *gin.Context) {
	id := c.GetInt("id") // 获取ID
	if id == 0 {
		response.Err(c, trans.Tr(c, "errGetQueryData"))
		return
	}

	// 过滤字段
	requestData := c.Request.PostForm    // 获取请求数据
	tableFields := ths.Model.Fields()    // 获取表字段
	realData := map[string]interface{}{} // 过滤后的数据
	for _, field := range tableFields {
		if val, exists := requestData[field]; exists {
			realData[field] = val
		}
	}

	siteCode := request.GetSiteCode(c)          // 获取平台
	db, err := clients.GetMySQLBySite(siteCode) // 获取数据库连接
	if err != nil {
		response.Err(c, trans.Tr(c, "errGetDbConn"))
		return
	}

	cond := map[string]interface{}{
		"id": id,
	}
	if _, err := ths.Model.Update(db, realData, cond); err != nil {
		response.Err(c, trans.Tr(c, "errUpdateData"))
		return
	}

	response.Ok(c)
}
