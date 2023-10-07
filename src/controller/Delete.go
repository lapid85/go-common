package controller

import (
	"common/clients"
	"common/model"
	"common/request"
	"common/response"
	"common/trans"
	"common/types"

	"github.com/gin-gonic/gin"
)

// Delete 删除
type ActionDelete struct {
	Model        model.IModel                    // model - 必须
	DeleteBefore func(*gin.Context) bool         // 默认处理函数
	DeleteAfter  func(*gin.Context, interface{}) // 默认处理函数
}

// Delete 删除 - delete
func (ths *ActionDelete) Delete(c *gin.Context) {
	req := types.QueryID{}
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Err(c, trans.Tr(c, "errGetQueryData"))
		return
	}

	siteCode := request.GetSiteCode(c)
	db, dbErr := clients.GetMySQLBySite(siteCode)
	if dbErr != nil {
		response.Err(c, trans.Tr(c, "errGetDbConn"))
		return
	}

	// 删除之前的处理
	if ths.DeleteBefore != nil {
		if !ths.DeleteBefore(c) {
			response.Err(c, trans.Tr(c, "errDeleteData"))
			return
		}
	}

	// 删除数据
	row, err := ths.Model.Delete(db, map[string]interface{}{
		"id": req.ID,
	})
	if err != nil {
		response.Err(c, trans.Tr(c, "errDeleteData"))
		return
	}

	// 针对获取结果进行处理
	if ths.DeleteAfter != nil {
		ths.DeleteAfter(c, row)
	}

	response.Ok(c)
}
