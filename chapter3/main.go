package main

import (
	"fmt"
	"math"

	"github.com/golang-collections/collections/queue"

	"github.com/golang-collections/collections/stack"
)

func main() {
	s := stack.New()
	s.Push(3)
	s.Push(5)
	s.Push(1)
	s.Push(10)
	s.Push(4)
	showAll(*s)
	sortStack(s)
	showAll(*s)
}

func showAll(s stack.Stack) {
	for s.Len() != 0 {
		fmt.Printf("%v, ", s.Pop().(int))
	}
	fmt.Println("")
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

// 3.4 スタックでキュー
type myQueue struct {
	stackNewest *stack.Stack
	stackOldest *stack.Stack
}

func newMyQueue() *myQueue {
	return &myQueue{
		stackNewest: stack.New(),
		stackOldest: stack.New(),
	}
}

func (q *myQueue) add(value int) {
	q.stackNewest.Push(value)
}

func (q *myQueue) remove() int {
	q.shiftStack()
	return q.stackOldest.Pop().(int)
}

func (q *myQueue) shiftStack() {
	if q.stackOldest.Len() != 0 {
		return
	}
	for q.stackNewest.Len() != 0 {
		q.stackOldest.Push(q.stackNewest.Pop())
	}
}

// 3.5 スタックのソート
func sortStack(s *stack.Stack) {
	r := stack.New()
	for s.Len() != 0 {
		tmp := s.Pop()
		for r.Len() != 0 && r.Peek().(int) >= tmp.(int) {
			s.Push(r.Pop())
		}
		r.Push(tmp)
	}
	for r.Len() != 0 {
		s.Push(r.Pop())
	}
}

// 3.6 動物保護施設
type animalType int

const (
	dog animalType = iota
	cat
)

type animal struct {
	order        int
	typeOfAnimal animalType
}

type animalQueue struct {
	dogQueue *queue.Queue
	catQueue *queue.Queue
	order    int
}

func newAnimalQueue() *animalQueue {
	return &animalQueue{
		dogQueue: queue.New(),
		catQueue: queue.New(),
		order:    0,
	}
}

func (q *animalQueue) enqueue(a animal) {
	q.order++
	a.order = q.order
	if a.typeOfAnimal == cat {
		q.catQueue.Enqueue(a)
	} else {
		q.dogQueue.Enqueue(a)
	}
}

func (q *animalQueue) dequeueAny() animal {
	if q.dogQueue.Len() == 0 {
		return q.catQueue.Dequeue().(animal)
	} else if q.catQueue.Len() == 0 {
		return q.dogQueue.Dequeue().(animal)
	}
	if q.dogQueue.Peek().(animal).order < q.catQueue.Peek().(animal).order {
		return q.dogQueue.Dequeue().(animal)
	} else {
		return q.catQueue.Dequeue().(animal)
	}
}
