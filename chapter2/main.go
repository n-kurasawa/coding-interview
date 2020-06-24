package main

import "fmt"

func main() {
	node1 := &node{next: nil, value: 1}
	node2 := &node{next: node1, value: 2}
	node3 := &node{next: node2, value: 3}
	node4 := &node{next: node3, value: 4}
	node5 := &node{next: node4, value: 5}
	node6 := &node{next: node5, value: 6}
	node7 := &node{next: node6, value: 7}
	showAll(node7)
	deleteNode(node4)
	showAll(node7)
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
