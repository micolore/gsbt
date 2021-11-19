package time

import (
	"fmt"
	"testing"
)

func TestCarbonTime(t *testing.T) {
	nowTime := GetNowTime()
	nowDate := GetNowDate()
	afterDay := AddDay("2022-11-18 18:40:55",60)

	fmt.Printf("AddDay:%s\n", afterDay)
	fmt.Printf("nowTime:%s\n", nowTime)
	fmt.Printf("nowDate:%s\n", nowDate)
}
