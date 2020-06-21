package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(compress("aabccccccaaaa"))
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

// 1.5 一発変換
func oneChange(str1, str2 string) bool {
	length1 := len(str1)
	length2 := len(str2)
	if length1 == length2 {
		return oneReplace(str1, str2)
	} else if (length1 - length2) == 1 {
		return oneAdd(str2, str1)
	} else if (length2 - length1) == 1 {
		return oneAdd(str1, str2)
	}
	return false
}

func oneAdd(short, long string) bool {
	diff := 0
	for i := 0; i < len(short); i++ {
		if short[i] != long[i + diff] {
			if diff != 0 {
				return false
			}
			diff++
		}
	}
	return true
}

func oneReplace(str1, str2 string) bool {
	diff := 0
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			if diff > 0 {
				return false
			}
			diff++
		}
	}
	return true
}

// 1.6 文字列圧縮
func compress(str string) string {
	var tmp rune
	var tmpIndex int
	var result []byte
	for i, v := range str {
		if tmp != v {
			if tmp != 0 {
				result = append(result, string(tmp) + strconv.Itoa(i-tmpIndex)...)
			}
			tmp = v
			tmpIndex = i
		}
	}
	result = append(result, string(tmp) + strconv.Itoa(len(str) - tmpIndex)...)
	if len(str) < len(result) {
		return str
	}
	return string(result)
}