package request

import (
	"common/utils"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

// Cond 查询条件
type Cond struct {
	Ands []string
	Ors  []string
}

// And 条件: and
func (ths *Cond) And(key string, value interface{}) {
	ths.Ands = append(ths.Ands, fmt.Sprintf("%s = %s", key, value))
}

// Or 条件: or
func (ths *Cond) Or(key string, value interface{}) {
	ths.Ors = append(ths.Ors, fmt.Sprintf("%s = %s", key, value))
}

// Gt 条件: >
func (ths *Cond) Gt(key string, value interface{}) {
	ths.Ands = append(ths.Ands, fmt.Sprintf("%s > %s", key, value))
}

// Lt 条件: <
func (ths *Cond) Lt(key string, value interface{}) {
	ths.Ands = append(ths.Ands, fmt.Sprintf("%s < %s", key, value))
}

// Ge 条件: >=
func (ths *Cond) Ge(key string, value interface{}) {
	ths.Ands = append(ths.Ands, fmt.Sprintf("%s >= %s", key, value))
}

// Le 条件: <=
func (ths *Cond) Le(key string, value interface{}) {
	ths.Ands = append(ths.Ands, fmt.Sprintf("%s <= %s", key, value))
}

// Like 条件: like
func (ths *Cond) Like(key string, value interface{}) {
	ths.Ands = append(ths.Ands, fmt.Sprintf("%s LIKE '%%%s%%'", key, value))
}

// Build 构建查询条件
func (ths *Cond) Build() string {
	cond := ""
	if len(ths.Ands) > 0 {
		cond += " AND " + strings.Join(ths.Ands, " AND ")
	}
	if len(ths.Ors) > 0 {
		cond += " OR " + strings.Join(ths.Ors, " OR ")
	}
	return cond
}

// GetQueryCond 得到查询条件
func GetQueryCond(c *gin.Context, data map[string]interface{}) *Cond {
	cond := &Cond{}
	for k, s := range data {
		val, exists := c.GetQuery(k)
		if !exists || val == "" {
			continue
		}
		switch sign := s.(type) {
		case string:
			switch sign {
			case "=":
				cond.And(k, val)
			case ">":
				cond.Gt(k, val)
			case ">=":
				cond.Ge(k, val)
			case "<":
				cond.Lt(k, val)
			case "<=":
				cond.Le(k, val)
			case "%":
				cond.Like(k, val)
			case "between", "[]": // 2个字段
				setQueryCondBetween(c, k, func(val interface{}) {
					cond.Ge(k, val)
				}, func(val interface{}) {
					cond.Le(k, val)
				})
			case "between|timestamp": // 2个字段
				setQueryCondBetween(c, k, func(val interface{}) {
					cond.Ge(k, utils.GetUnixMicroByDateTime(val.(string)))
				}, func(val interface{}) {
					cond.Le(k, utils.GetUnixMicroByDateTime(val.(string)))
				})
			case "[]|timestamp": // 2个字段
				setQueryCondBetween(c, k, func(val interface{}) {
					cond.Ge(k, utils.GetUnixMicroByDateTime(val.(string)))
				}, func(val interface{}) {
					cond.Le(k, utils.GetUnixMicroByDateTime(val.(string)))
				})
			case "[_]|timestamp": // 1个字段拆分 => 比如 created, 传过来的是时间区间 2020-01-02 - 2020-02-02 => 转换为时间戳对比
				var startAt int64
				var endAt int64
				value, exists := c.GetQuery(k)
				if exists {
					areas := strings.Split(value, " - ")
					startAt = utils.GetUnixMicroByDateTime(areas[0])
					endAt = utils.GetUnixMicroByDateTime(areas[1])
					cond.Ge(k, startAt)
					cond.Le(k, endAt)
				}
			case "()":
				setQueryCondBetween(c, k, func(val interface{}) {
					cond.Gt(k, val)
				}, func(val interface{}) {
					cond.Lt(k, val)
				})
			case "[)":
				setQueryCondBetween(c, k, func(val interface{}) {
					cond.Ge(k, val)
				}, func(val interface{}) {
					cond.Lt(k, val)
				})
			case "(]":
				setQueryCondBetween(c, k, func(val interface{}) {
					cond.Gt(k, val)
				}, func(val interface{}) {
					cond.Le(k, val)
				})
			default:
				continue
			}
		case func(*gin.Context) Cond: // 多余的条件
			result := s.(func(*gin.Context) Cond)(c)
			cond.Ands = append(result.Ands)
		default:
			continue
		}
	}
	return cond
}

// setQueryCondBetween 设置以区间为条件的查询字段
func setQueryCondBetween(c *gin.Context, field string, callStart func(interface{}), callEnd func(interface{})) {
	start := field + "_start"
	end := field + "_end"
	if startVal, exists := c.GetQuery(start); exists {
		callStart(startVal)
	}
	if endVal, exists := c.GetQuery(end); exists {
		callEnd(endVal)
	}
}
