package time

import (
	"fmt"
	"testing"
)

func TestCarbonTime(t *testing.T) {
	nowTime := GetNowTime()
	nowDate := GetNowDate()

	fmt.Printf("nowTime:%s\n", nowTime)
	fmt.Printf("nowDate:%s\n", nowDate)
}
