package main

import "fmt"

func main() {
	// thống kê số lần "từ" xuất hiện nhiều nhất
	words := map[string]int{
		"Gonna": 2,
		"You":   3,
		"Give":  3,
		"Never": 1,
		"Up":    4,
	}
	topCount := 0
	topWord := ""
	for key, val := range words {
		if val > topCount {
			topCount = val
			topWord = key
		}
	}
	fmt.Println("key : ", topWord)
	fmt.Println("count : ", topCount)
}
