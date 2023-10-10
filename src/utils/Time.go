// Package utils @Author:  cobol
// @Comment: 时间相关工具 - 注意: 考虑到时区因素， 所有时间都是UTC时间
package utils

import (
	"strconv"
	"strings"
	"time"
)

// BaseTimeMicro 默认微秒累加  - 2099-01-01 00:00:00
const BaseTimeMicro = 4070880000

// Time 标准获取当前时间 - 考虑时区因素
func Time() time.Time {
	return time.Now().UTC()
}

// CurrentTime 标准获取当前时间 - 考虑时区因素
func CurrentTime() time.Time {
	return Time()
}

// Now 标准获取当前时间
func Now() time.Time {
	return Time()
}

// Timestamp 标准获取当前时间
func Timestamp() int64 {
	return Time().Unix()
}

// NowMicro 当前微秒
func NowMicro() int64 {
	return Time().UnixMicro()
}

// GetTimeByUnix 通过unix获取时间 - 秒
func GetTimeByUnix(ts int64) time.Time {
	return time.Unix(ts, 0).UTC()
}

// GetDateTimeByUnix 通过unix获取时间 - 秒
func GetDateTimeByUnix(ts int64) string {
	return GetTimeByUnix(ts).Format("2006-01-02 15:04:05")
}

// GetUnixByDateTime 通过时间获取时间
func GetUnixByDateTime(dateTime string) int64 {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", dateTime, time.UTC)
	return t.Unix()
}

// GetTimeByUnixMicro 通过unix获取时间 - 微秒
func GetTimeByUnixMicro(int64Time int64) time.Time {
	return time.Unix(int64Time*1000000, 0).UTC()
}

// GetDateTimeByUnixMicro 通过unix获取时间 - 微秒
func GetDateTimeByUnixMicro(ts int64) string {
	return GetTimeByUnixMicro(ts).Format("2006-01-02 15:04:05")
}

// GetUnixMicroByDateTime GetTimeByDateTime 通过时间获取时间
func GetUnixMicroByDateTime(dateTime string) int64 {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", dateTime, time.UTC)
	return t.UnixMicro()
}

// SecondToMicro 秒转为微秒
func SecondToMicro(second int64) int64 {
	if second > BaseTimeMicro {
		return second
	}
	return second * 1000000
}

// MicroToSecond 转换为秒数
func MicroToSecond(micro int64) int64 {
	if micro < BaseTimeMicro {
		return micro
	}
	return micro / 1000000
}

// Date 得到当前的年/月/日
func Date() (int, int, int) {
	now := Time()
	return DateOf(now)
}

// DateOf 得到当前的年/月/日
func DateOf(now time.Time) (int, int, int) {
	ymd := now.Format("2006-01-02")
	ymdArr := strings.Split(ymd, "-")
	year, _ := strconv.Atoi(ymdArr[0])
	month, _ := strconv.Atoi(ymdArr[1])
	day, _ := strconv.Atoi(ymdArr[2])
	return year, month, day
}

// GetDayStart 本日开始时间
func GetDayStart() time.Time {
	now := Time()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
}

// GetDayEnd 本日结束时间
func GetDayEnd() time.Time {
	now := Time()
	return time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.UTC)
}

// GetYesterdayStart 昨日开始时间
func GetYesterdayStart() time.Time {
	now := Time()
	return time.Date(now.Year(), now.Month(), now.Day()-1, 0, 0, 0, 0, time.UTC)
}

// GetYesterdayEnd 昨日结束时间
func GetYesterdayEnd() time.Time {
	now := Time()
	return time.Date(now.Year(), now.Month(), now.Day()-1, 23, 59, 59, 0, time.UTC)
}

// GetWeekStart 本周第一天
func GetWeekStart() time.Time {
	now := Time()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC).AddDate(0, 0, offset)
	return weekStart
}

// GetWeekEnd 本周最后一天
func GetWeekEnd() time.Time {
	now := Time()
	offset := int(time.Sunday - now.Weekday())
	if offset < 0 {
		offset = 7
	}
	weekEnd := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC).AddDate(0, 0, offset)
	return weekEnd
}

// GetLastWeekStart 上周第一天
func GetLastWeekStart() time.Time {
	now := Time()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC).AddDate(0, 0, offset-7)
	return weekStart
}

// GetLastWeekEnd 上周最后一天
func GetLastWeekEnd() time.Time {
	now := Time()
	offset := int(time.Sunday - now.Weekday())
	if offset < 0 {
		offset = 7
	}
	weekEnd := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC).AddDate(0, 0, offset-7)
	return weekEnd
}

// GetMonthStart 得到本月第一天
func GetMonthStart() time.Time {
	now := Time()
	currentYear, currentMonth, _ := now.Date()
	return time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, time.UTC)
}

// GetMonthEnd 得到本月最后一天
func GetMonthEnd() time.Time {
	now := Time()
	currentYear, currentMonth, _ := now.Date()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, time.UTC)
	return firstOfMonth.AddDate(0, 1, -1)
}

// GetLastMonthStart 得到上月第一天
func GetLastMonthStart() time.Time {
	now := Time()
	currentYear, currentMonth, _ := now.Date()
	return time.Date(currentYear, currentMonth-1, 1, 0, 0, 0, 0, time.UTC)
}

// GetLastMonthEnd 得到上月最后一天
func GetLastMonthEnd() time.Time {
	now := Time()
	currentYear, currentMonth, _ := now.Date()
	firstOfMonth := time.Date(currentYear, currentMonth-1, 1, 0, 0, 0, 0, time.UTC)
	return firstOfMonth.AddDate(0, 1, -1)
}
