package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func main() {
	node1 := &node{next: nil, value: 7}
	node2 := &node{next: node1, value: 1}
	node3 := &node{next: node2, value: 6}
	node4 := &node{next: node3, value: 1}
	node5 := &node{next: node4, value: 7}

	node6 := &node{next: node3, value: 11}
	node7 := &node{next: node6, value: 17}
	node8 := &node{next: node7, value: 18}
	node9 := &node{next: node8, value: 19}
	showAll(node5)
	showAll(node9)
	node := commonNode(node5, node9)
	showAll(node)
}

func showAll(node *node) {
	p := node
	for p != nil {
		fmt.Printf("%v, ", p.value)
		p = p.next
	}
	fmt.Println("")
}

type node struct {
	next  *node
	value int
}

// 2.1 重複要素の削除
func deleteDuplication(node *node, m map[int]bool) *node {
	if node == nil {
		return nil
	}
	if m[node.value] {
		return node.next
	}
	m[node.value] = true
	node.next = deleteDuplication(node.next, m)
	return node
}

func deleteDuplication2(head *node) {
	n := head
	m := map[int]bool{}
	var previous *node
	for n != nil {
		if m[n.value] {
			previous.next = n.next
		} else {
			m[n.value] = true
			previous = n
		}
		n = n.next
	}
}

// 2.2 後ろからK番目を返す
func getFromBack(head *node, index int) (*node, int) {
	if head == nil {
		return nil, 0
	}

	node, i := getFromBack(head.next, index)
	i++
	if i == index {
		return head, i
	} else {
		return node, i
	}
}

// 2.3 間の要素を削除
func deleteNode(n *node) {
	if n == nil || n.next == nil {
		return
	}

	n.value = n.next.value
	n.next = n.next.next
}

// 2.4 リストの分割
func divideList(head *node, divide int) *node {
	var beforeHead *node
	var before *node
	var afterHead *node
	var after *node
	n := head
	for n != nil {
		if n.value < divide {
			if before == nil {
				beforeHead = n
				before = n
			} else {
				before.next = n
				before = n
			}
		} else {
			if after == nil {
				afterHead = n
				after = n
			} else {
				after.next = n
				after = n
			}
		}
		n = n.next
	}
	after.next = nil
	before.next = afterHead
	return beforeHead
}

func divideList2(n *node, divide int) *node {
	head := n
	tail := n

	p := n
	for p != nil {
		next := p.next
		if p.value < divide {
			p.next = head
			head = p
		} else {
			tail.next = p
			tail = p
		}
		p = next
	}
	tail.next = nil
	return head
}

// 2.5 リストで著された2数の和
func sum(head1, head2 *node, carry int) *node {
	if head1 == nil && head2 == nil && carry == 0 {
		return nil
	}
	num := carry
	var next1 *node
	if head1 != nil {
		num += head1.value
		next1 = head1.next
	}
	var next2 *node
	if head2 != nil {
		num += head2.value
		next2 = head2.next
	}
	var nextCarry int
	if num > 9 {
		nextCarry = 1
	}
	var result *node
	if next1 != nil || next2 != nil {
		result = sum(next1, next2, nextCarry)
	}

	return &node{
		value: num % 10,
		next:  result,
	}
}

// 2.5 発展
func sum2(head1, head2 *node) *node {
	len1 := length(head1)
	len2 := length(head2)

	if len1 < len2 {
		head1 = padList(head1, len2-len1)
	} else {
		head2 = padList(head2, len1-len2)
	}

	result, carry := sumHelper(head1, head2)
	if carry > 0 {
		return &node{
			value: 1,
			next:  result,
		}
	}
	return result
}

func sumHelper(head1, head2 *node) (*node, int) {
	if head1 == nil && head2 == nil {
		return nil, 0
	}

	result, carry := sumHelper(head1.next, head2.next)

	sum := carry + head1.value + head2.value
	var nextCarry int
	if sum > 9 {
		nextCarry = 1
	}
	return &node{
		value: sum % 10,
		next:  result,
	}, nextCarry

}

func length(head *node) int {
	var count int
	for head != nil {
		count++
		head = head.next
	}
	return count
}

func padList(head *node, num int) *node {
	for i := 0; i < num; i++ {
		n := &node{
			value: 0,
			next:  head,
		}
		head = n
	}
	return head
}

// 2.6 回文
func isPalindrome(head *node) bool {
	reverseNode := reverse(head)
	return isEqual(head, reverseNode)
}

func reverse(n *node) *node {
	var head *node
	for n != nil {
		newNode := &node{
			value: n.value,
			next:  n.next,
		}
		newNode.next = head
		head = newNode
		n = n.next
	}
	return head
}

func isEqual(head1, head2 *node) bool {
	for head1 != nil && head2 != nil {
		if head1.value != head2.value {
			return false
		}
		head1 = head1.next
		head2 = head2.next
	}
	return head1 == nil && head2 == nil
}

func isPalindrome2(head *node) bool {
	fast := head
	slow := head
	s := stack.New()

	for fast != nil && fast.next != nil {
		s.Push(slow)
		slow = slow.next
		fast = fast.next.next
	}

	if fast != nil {
		slow = slow.next
	}

	for slow != nil {
		top := s.Pop().(*node)
		if slow.value != top.value {
			return false
		}
		slow = slow.next
	}
	return true
}

func isPalindrome3(head *node) bool {
	length := length(head)
	_, result := isPalindromeHelper(head, length)
	return result
}

func isPalindromeHelper(head *node, length int) (*node, bool) {
	if head == nil || length <= 0 {
		return head, true
	} else if length == 1 {
		return head.next, true
	}

	n, result := isPalindromeHelper(head.next, length-2)

	if n == nil || !result || n.value != head.value {
		return nil, false
	}

	return n.next, true
}

// 2.7 共通するノード
func commonNode(head1, head2 *node) *node {
	n1 := head1
	var len1 int
	for n1 != nil {
		len1++
		n1 = n1.next
	}

	n2 := head2
	var len2 int
	for n2 != nil {
		len2++
		n2 = n2.next
	}

	if n1 != n2 {
		return nil
	}

	if len1 < len2 {
		head2 = removeNode(head2, len2-len1)
	} else if len2 < len1 {
		head1 = removeNode(head1, len1-len2)
	}

	for i := 0; head1 != nil && head2 != nil; i++ {
		if head1 == head2 {
			return head1
		}
		head1 = head1.next
		head2 = head2.next
	}
	return nil
}

func removeNode(head *node, num int) *node {
	for i := 0; i < num; i++ {
		head = head.next
	}
	return head
}
