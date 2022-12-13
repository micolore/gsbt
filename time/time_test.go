package time

import (
	"fmt"
	"github.com/golang-module/carbon"
	"testing"
)

func TestCarbonTime(t *testing.T) {
	nowTime := GetNowTime()
	nowDate := GetNowDate()
	afterDay := AddDay("2022-11-18 18:40:55", 60)

	fmt.Printf("AddDay:%s\n", afterDay)
	fmt.Printf("nowTime:%s\n", nowTime)
	fmt.Printf("nowDate:%s\n", nowDate)

	nowStr := carbon.Now().ToString() // 2020-08-05 13:14:15 +0800 CST
	fmt.Println(nowStr)
	//utcTime := carbon.SetTimezone(carbon.UTC).Now().ToDateTimeString() // 2022-06-28 09:25:38
	utcT := carbon.SetTimezone(carbon.UTC).Now()
	fmt.Println(utcT.ToDateTimeString())

	secondTime := utcT.SetTimezone(carbon.Shanghai).ToDateTimeString()
	fmt.Println(secondTime)

}
