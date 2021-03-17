package basic

import "fmt"

func TesSlice() {
	var numbers = []int{1, 3, 5, 6, 7, 8, 9, 10, 11, 12, 13, 134, 111, 234, 14445}
	l := len(numbers)
	page := l / 3
	startInsertIndex := 0
	endInsertSize := 0
	insertSize := 3
	for i := 0; i <= page; i++ {
		endInsertSize = endInsertSize + insertSize
		if endInsertSize > l {
			endInsertSize = l
		}
		insertData := numbers[startInsertIndex:endInsertSize]
		if insertData == nil || len(insertData) <= 0 {
			break
		}
		fmt.Printf("l %v\n", insertData)
		startInsertIndex = startInsertIndex + insertSize
	}
}
