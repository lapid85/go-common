package controller

import (
	"common/clients"
	"common/log"
	"common/model"
	"common/request"
	"common/response"
	"common/trans"

	"github.com/gin-gonic/gin"
)

// ActionList 数据列表
type ActionList struct {
	Model       model.IModel                    // model - 必须
	QueryCond   map[string]interface{}          // 查询条件
	ProcessRow  func(*gin.Context, interface{}) // 默认处理函数
	OrderBy     func(*gin.Context) string       // 获取排序
	AfterAction func(*gin.Context, *map[string]interface{})
}

// List 记录列表 - get
func (ths *ActionList) List(c *gin.Context) {
	limit, offset := request.GetOffset(c)
	platform := request.GetPlatform(c)

	db, dbErr := clients.GetMySQLByPlatform(platform)
	if dbErr != nil {
		response.Err(c, trans.Tr(c, "errGetDbConn"))
		return
	}

	// 自定查询条件
	cond := request.GetQueryCond(c, ths.QueryCond)
	// 关于 order by 的判断
	var rows interface{}
	var total int64
	var err error
	whereStr := cond.Build()
	log.Info("whereStr: %s", whereStr)
	if ths.OrderBy != nil {
		rows, total, err = ths.Model.GetAll(db, whereStr, []int{limit, offset}, ths.OrderBy(c))
	} else {
		rows, total, err = ths.Model.GetAll(db, whereStr, []int{limit, offset})
	}
	if err != nil {
		log.Error(err.Error())
		response.Err(c, trans.Tr(c, "errGetListData"))
		return
	}

	// 针对数组进行处理
	if ths.ProcessRow != nil {
		ths.ProcessRow(c, rows)
	}

	// 处理发送之前数据
	data := map[string]interface{}{
		"title": trans.Tr(c, "errGetListData"), // 测试多语言自动翻译
		"rows":  rows,
		"total": total,
	}
	if ths.AfterAction != nil {
		ths.AfterAction(c, &data)
	}

	response.Data(c, data)
}
