package Utils

import (
	"time"
)

// TimeStr 返回字符串类型 2006-01-02 15:04:05 格式时间
func TimeStr() string {
	Time := time.Now().Format("2006-01-02 15:04:05")
	return Time
}

// Time 返回字符串类型 20060102150405 格式时间
func Time() string {
	Time := time.Now().Format("20060102150405")
	return Time
}

// Year 获取当前年份字符串格式
func Year() string {
	NowYear := time.Now().Format("2006")
	return NowYear
}
func Month() string {
	NowMonth := time.Now().Format("01")
	return NowMonth
}

// Date 获取当前日期字符串格式
func Date() string {
	NowDate := time.Now().Format("2006-01-02")
	return NowDate
}

// FirstOfMonth 获取当月第一天字符串格式
func FirstOfMonth() string {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	layout := "2006-01-02"
	return firstOfMonth.Format(layout)
}

// LastOfMonth 获取当月最后一天字符串格式
func LastOfMonth() string {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	layout := "2006-01-02"
	return lastOfMonth.Format(layout)

}

// CaculateMonth 计算本月开始加减月
func CaculateMonth(Years int, Months int, days int) string {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(Years, Months, days)
	layout := "2006-01-02"
	return lastOfMonth.Format(layout)
}

// Str2Unix 字符串格式日期转时间戳
func Str2Unix(str string) (int64, error) {
	timeLayout := "2006-01-02 15:04:05"                        //转化所需模板
	loc, _ := time.LoadLocation("Local")                       //重要：获取时区
	theTime, err := time.ParseInLocation(timeLayout, str, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()
	return sr, err
}
