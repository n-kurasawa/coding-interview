package main

import (
	"sort"
)

func main() {
}

// 1.1 重複のない文字列
func noDuplication(str string) bool {
	m := map[int32]bool{}
	for _, v := range str {
		if m[v] {
			return false
		}
		m[v] = true
	}
	return true
}

func noDuplication2(str string) bool {
	length := len(str)
	for i := 0; i < length; i++ {
		for j := i+1; j < length; j++ {
			if str[i] == str[j] {
				return false
			}
		}
	}
	return true
}

func noDuplication3(str string) bool {
	arr := []int32(str)
	sort.Slice(arr, func(i int, j int) bool { return arr[i] < arr[j]})

	for i := 0; i + 1 < len(arr); i++ {
		if arr[i] == arr[i+1] {
			return false
		}
	}
	return true
}