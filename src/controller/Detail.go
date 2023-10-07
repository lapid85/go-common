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

type ActionDetail struct {
	Model       model.IModel                    // model - 必须
	DetailAfter func(*gin.Context, interface{}) // 默认处理函数
}

// Detail 详情 - get
func (ths *ActionDetail) Detail(c *gin.Context) {
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

	row, err := ths.Model.Get(db, map[string]interface{}{
		"id": req.ID,
	})
	if err != nil {
		response.Err(c, trans.Tr(c, "errGetDetailData"))
		return
	}

	// 针对获取结果进行处理
	if ths.DetailAfter != nil {
		ths.DetailAfter(c, &row)
	}

	response.Data(c, row)
}
