package main

import "fmt"

func main() {
	arr := make([]int, 0, 10)
	fmt.Println(arr[10])
}

// 3.1 3つのスタック
type FixedMultiStack struct {
	numberOfStacks int
	capacity       int
	values         []int
	sizes          []int
}

func NewFixedMultiStack() *FixedMultiStack {
	return &FixedMultiStack{
		numberOfStacks: 3,
		capacity:       100,
		values:         make([]int, 300, 300),
		sizes:          make([]int, 3, 3),
	}
}

func (s *FixedMultiStack) push(stackNum, value int) error {
	if s.isFull(stackNum) {
		return fmt.Errorf("stack is full. stackNum: %v", stackNum)
	}
	s.sizes[stackNum]++
	s.values[s.indexOfTop(stackNum)] = value
	return nil
}

func (s *FixedMultiStack) pop(stackNum int) (int, error) {
	if s.isEmpty(stackNum) {
		return 0, fmt.Errorf("stack is empty. stackNum: %v", stackNum)
	}
	index := s.indexOfTop(stackNum)
	value := s.values[index]
	s.values[index] = 0
	s.sizes[stackNum]--
	return value, nil
}

func (s *FixedMultiStack) peek(stackNum int) (int, error) {
	if s.isEmpty(stackNum) {
		return 0, fmt.Errorf("stack is empty. stackNum: %v", stackNum)
	}
	return s.values[s.indexOfTop(stackNum)], nil
}

func (s *FixedMultiStack) isFull(stackNum int) bool {
	return s.sizes[stackNum] == s.capacity
}

func (s *FixedMultiStack) isEmpty(stackNum int) bool {
	return s.sizes[stackNum] == 0
}

func (s *FixedMultiStack) indexOfTop(stackNum int) int {
	offset := stackNum * s.capacity
	size := s.sizes[stackNum]
	return offset + size - 1
}
