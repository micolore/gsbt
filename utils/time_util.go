package utils

import (
"fmt"
"strconv"
"time"
)

const DefaultTimeStr string = "2006-01-02 15:04:05"

// 获取指定日期的时间（前或者后）
func GetAnyDayDate(day int) time.Time {
	agoDay := day * 24
	currentTime := time.Now()
	agoDayStr := fmt.Sprintf("-%dh", agoDay)
	m, _ := time.ParseDuration(agoDayStr)
	result := currentTime.Add(m)
	return result
}

// 获取指定月份之前的日期（不是很准）
func GetAnyMonthDate(month int) time.Time {
	agoDay := month * 24 * 31
	currentTime := time.Now()
	agoDayStr := fmt.Sprintf("-%dh", agoDay)
	m, _ := time.ParseDuration(agoDayStr)
	result := currentTime.Add(m)
	return result
}

// GetMonthStartAndEndByFlag 获取月份的第一天和最后一天
func GetMonthStartAndEndByFlag(flag bool, year string, month string) time.Time {
	if len(month) == 1 {
		month = "0" + month
	}
	yearInt, _ := strconv.Atoi(year)
	startTime := fmt.Sprintf("%s-%s-01 00:00:00", year, month)
	theTime, _ := time.ParseInLocation(DefaultTimeStr, startTime, getLocalTime())
	newMonth := theTime.Month()
	var t time.Time
	if flag {
		t = time.Date(yearInt, newMonth, 1, 0, 0, 0, 0, time.Local)
	} else {
		t = time.Date(yearInt, newMonth+1, 0, 0, 0, 0, 0, time.Local)
	}
	return t
}

// getLocalTime 获取当前的Location
func getLocalTime() *time.Location {
	loc, _ := time.LoadLocation("Local")
	return loc
}

// 获取一个月里面所有的天数
func GetMonthDays(year string, month string) []time.Time {
	var list = make([]time.Time, 0)
	if len(month) == 1 {
		month = "0" + month
	}
	yearInt, _ := strconv.Atoi(year)
	startTime := fmt.Sprintf("%s-%s-01 00:00:00", year, month)
	theTime, _ := time.ParseInLocation(DefaultTimeStr, startTime, getLocalTime())
	newMonth := theTime.Month()

	start := time.Date(yearInt, newMonth, 1, 0, 0, 0, 0, time.Local)
	end := time.Date(yearInt, newMonth+1, 0, 0, 0, 0, 0, time.Local)

	days := end.Day() - start.Day()
	for i := 0; i <= days; i++ {
		day := time.Date(yearInt, newMonth, i+1, 0, 0, 0, 0, time.Local)
		list = append(list, day)
	}
	return list

}

// TimeToString 日期转字符串
func TimeToString(time time.Time) string {

	return time.Format(DefaultTimeStr)
}

