package controller

import (
	"common/clients"
	"common/log"
	"common/model"
	"common/request"
	"common/response"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

// ActionList 数据列表
type ActionList struct {
	Model      model.IModel                    // model - 必须
	QueryCond  map[string]interface{}          // 查询条件
	ProcessRow func(*gin.Context, interface{}) // 默认处理函数
	OrderBy    func(*gin.Context) string       // 获取排序
}

// List 记录列表 - get
func (ths *ActionList) List(c *gin.Context) {
	limit, offset := request.GetOffset(c)
	// platform := request.GetPlatform(c)

	db, dbErr := clients.MySQLDefault()
	if dbErr != nil {
		response.Err(c, "获取数据库连接错误")
		return
	}

	cond := request.GetQueryCond(c, ths.QueryCond)
	// 关于 order by 的判断
	var rows interface{}
	var total int64
	var err error
	whereStr := cond.Build()
	log.Info("whereStr:", whereStr)
	if ths.OrderBy != nil {
		rows, total, err = ths.Model.GetAll(db, whereStr, []int{limit, offset}, ths.OrderBy(c))
	} else {
		rows, total, err = ths.Model.GetAll(db, whereStr, []int{limit, offset})
	}

	if ths.ProcessRow != nil {
		ths.ProcessRow(c, rows)
	}
	if err != nil {
		slog.Error(err.Error())
		response.Err(c, "获取列表数据错误")
		return
	}

	response.Data(c, map[string]interface{}{
		"rows":  rows,
		"total": total,
	})
}
