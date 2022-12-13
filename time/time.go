package time

import (
	"github.com/golang-module/carbon"
)

func GetNowTime() string {
	return carbon.Now().ToDateTimeString()
}
func GetNowDate() string {
	return carbon.Now().ToDateString()
}
func AddDay(dateTime string, day int) string {
	return carbon.Parse(dateTime).AddDays(day).String()
}
