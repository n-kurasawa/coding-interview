package main

import (
	"fmt"
	"strconv"
)

func main() {
	n, _ := strconv.ParseInt("11111111", 2, 0)
	m, _ := strconv.ParseInt("10011", 2, 0)
	fmt.Printf("%08b\n", n)
	fmt.Printf("%08b\n", m)
	fmt.Printf("%08b\n", updateBits(int(n), int(m), 2, 6))
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
