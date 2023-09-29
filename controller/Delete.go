package controller

import (
	"common/clients"
	"common/model"
	"common/response"
	"common/types"

	"github.com/gin-gonic/gin"
)

// Delete 删除
type ActionDelete struct {
	Model model.IModel // model - 必须
}

// Delete 删除 - delete
func (ths *ActionDelete) Delete(c *gin.Context) {
	req := types.QueryID{}
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Err(c, "参数错误")
		return
	}
	db, dbErr := clients.MySQLDefault()
	if dbErr != nil {
		response.Err(c, "获取数据库连接错误")
		return
	}

	_, err := ths.Model.Delete(db, map[string]interface{}{
		"id": req.ID,
	})

	if err != nil {
		response.Err(c, "删除数据错误")
		return
	}

	response.Ok(c)
}
