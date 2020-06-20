package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isPermutationOfPalindrome("tacocatt"))
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

// 1.2順列チェック
func isPermutation(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	arr1 := []int32(str1)
	sort.Slice(arr1, func(i int, j int) bool { return arr1[i] < arr1[j]})

	arr2 := []int32(str2)
	sort.Slice(arr2, func(i int, j int) bool { return arr2[i] < arr2[j]})

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func isPermutation2(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	m := map[int32]int{}
	for _, v := range str1 {
		m[v]++
	}
	for _, v := range str2 {
		m[v]--
		if m[v] < 0 {
			return false
		}
	}
	return true
}

// 1.3 URLify
func urlify(str []rune, length int) {
	var blank int
	for i := 0; i < length; i++ {
		if str[i] == ' ' {
			blank++
		}
	}
	index := length + blank * 2
	for i := length - 1; i > -1; i-- {
		if str[i] == ' ' {
			str[index-1] = '0'
			str[index-2] = '2'
			str[index-3] = '%'
			index -= 3
		} else {
			str[index-1] = str[i]
			index--
		}
	}
}

// 1.4 回文の順列
func isPermutationOfPalindrome(str string) bool {
	m := map[rune]bool{}
	for _, str := range str {
		m[str] = !m[str]
	}
	var count int
	for _, v := range m {
		if v {
			count++
			if count > 1 {
				return false
			}
		}
	}
	return true
}