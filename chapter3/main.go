package main

import (
	"fmt"
	"math"

	"github.com/golang-collections/collections/stack"
)

func main() {
	s := newStackWithMin()
	s.push(3)
	s.push(4)
	s.push(5)
	fmt.Println(s.min())
	s.push(1)
	fmt.Println(s.min())
	s.pop()
	fmt.Println(s.min())
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

// 3.2 最小値を返すスタック
type stackWithMin struct {
	*stack.Stack
	minStack *stack.Stack
}

func newStackWithMin() *stackWithMin {
	min := stack.New()
	min.Push(math.MaxInt32)
	return &stackWithMin{
		stack.New(),
		min,
	}
}

func (s *stackWithMin) push(value int) {
	if value <= s.minStack.Peek().(int) {
		s.minStack.Push(value)
	}
	s.Stack.Push(value)
}

func (s *stackWithMin) pop() int {
	value := s.Stack.Pop().(int)
	if s.minStack.Peek().(int) == value {
		s.minStack.Pop()
	}
	return value
}

func (s *stackWithMin) min() int {
	return s.minStack.Peek().(int)
}

// 積み上がっている皿
type SetOfStack struct {
	stacks *stack.Stack
}

func newSetOfStack() *SetOfStack {
	return &SetOfStack{
		stacks: stack.New(),
	}
}

func (s *SetOfStack) push(value int) {
	if s.stacks.Peek().(*stack.Stack).Len() > 100 {
		s.stacks.Push(stack.New())
	}
	s.stacks.Peek().(*stack.Stack).Push(value)
}

func (s *SetOfStack) pop() (int, error) {
	if s.stacks.Peek().(*stack.Stack).Len() == 0 {
		s := s.stacks.Pop()
		if s == nil {
			return 0, fmt.Errorf("empty")
		}
	}
	return s.stacks.Peek().(*stack.Stack).Pop().(int), nil
}
