package types

import (
	"fmt"
	"strings"
)

// Cond 查询条件
type Cond struct {
	Ands []string
	Ors  []string
}

// And 条件: and
func (ths Cond) And(key string, value interface{}) *Cond {
	ths.Ands = append(ths.Ands, fmt.Sprintf("%s = %s", key, value))
	return &ths
}

// Eq 条件: =
func (ths Cond) Eq(key string, value interface{}) *Cond {
	ths.Ands = append(ths.Ands, fmt.Sprintf("%s = %s", key, value))
	return &ths
}

// Or 条件: or
func (ths Cond) Or(key string, value interface{}) *Cond {
	ths.Ors = append(ths.Ors, fmt.Sprintf("%s = %s", key, value))
	return &ths
}

// Gt 条件: >
func (ths Cond) Gt(key string, value interface{}) *Cond {
	ths.Ands = append(ths.Ands, fmt.Sprintf("%s > %s", key, value))
	return &ths
}

// Lt 条件: <
func (ths Cond) Lt(key string, value interface{}) *Cond {
	ths.Ands = append(ths.Ands, fmt.Sprintf("%s < %s", key, value))
	return &ths
}

// Ge 条件: >=
func (ths Cond) Ge(key string, value interface{}) *Cond {
	ths.Ands = append(ths.Ands, fmt.Sprintf("%s >= %s", key, value))
	return &ths
}

// Le 条件: <=
func (ths Cond) Le(key string, value interface{}) *Cond {
	ths.Ands = append(ths.Ands, fmt.Sprintf("%s <= %s", key, value))
	return &ths
}

// Like 条件: like
func (ths Cond) Like(key string, value interface{}) *Cond {
	ths.Ands = append(ths.Ands, fmt.Sprintf("%s LIKE '%%%s%%'", key, value))
	return &ths
}

// Where 条件: where
func (ths Cond) Where(wh string, args ...interface{}) *Cond {
	if len(args) > 0 {
		wh = fmt.Sprintf(wh, args...)
		return &ths
	}

	ths.Ands = append(ths.Ands, wh)
	return &ths
}

// Build 构建查询条件
func (ths Cond) Build() string {
	cond := ""
	if len(ths.Ands) > 0 {
		cond += " AND " + strings.Join(ths.Ands, " AND ")
	}
	if len(ths.Ors) > 0 {
		cond += " OR " + strings.Join(ths.Ors, " OR ")
	}
	return cond
}
