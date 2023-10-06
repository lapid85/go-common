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
	Model      model.IModel                                // model - 必须
	QueryCond  map[string]interface{}                      // 查询条件
	ProcessRow func(*gin.Context, interface{})             // 默认处理函数
	OrderBy    func(*gin.Context) string                   // 获取排序
	ListAfter  func(*gin.Context, *map[string]interface{}) // 默认处理函数 - 用于判断或者追加数据
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
	whereStr := cond.Build()
	rows, total, err := func() (interface{}, int64, error) {
		if ths.OrderBy != nil {
			return ths.Model.GetAll(db, whereStr, []int{limit, offset}, ths.OrderBy(c))
		}
		return ths.Model.GetAll(db, whereStr, []int{limit, offset})
	}()
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
		"rows":  rows,                          // 数据
		"total": total,                         // 总数
	}
	if ths.ListAfter != nil {
		ths.ListAfter(c, &data)
	}

	response.Data(c, data)
}
