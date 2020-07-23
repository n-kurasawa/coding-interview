package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(printBinary(0.625))
}

func getBit(num int, i uint) bool {
	return (num & (1 << i)) != 0
}

func setBit(num int, i uint) int {
	return num | (1 << i)
}

func clearBit(num int, i uint) int {
	mask := ^(1 << i)
	return num & mask
}

func clearBitsMSBThroughI(num int, i uint) int {
	mask := (1 << i) - 1
	return num & mask
}

func clearBitsIThrough0(num int, i uint) int {
	mask := -1 << (i + 1)
	return num & mask
}

func updateBit(num int, i uint, bitIs1 bool) int {
	var value int
	if bitIs1 {
		value = 1
	}
	mask := ^(1 << i)
	return (num & mask) | value<<i
}

// 5.1 挿入
func updateBits(n, m int, i, j uint) int {
	allOnes := ^0
	left := allOnes << (j + 1)
	right := (1 << i) - 1

	mask := left | right

	nCleared := n & mask
	mShifted := m << i

	return nCleared | mShifted
}

// 実数の2進数表記
func printBinary(num float64) string {
	if num >= 1 || num <= 0 {
		return "ERROR"
	}

	binary := make([]string, 0)
	binary = append(binary, ".")

	for num > 0 {
		if len(binary) >= 32 {
			return "ERROR"
		}
		r := num * 2
		if r >= 1 {
			binary = append(binary, "1")
			num = r - 1
		} else {
			binary = append(binary, "0")
			num = r
		}
	}
	return strings.Join(binary, "")
}
