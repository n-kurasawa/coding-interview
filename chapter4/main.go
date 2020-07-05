package main

import (
	"container/list"
	"errors"
	"fmt"
	"math"

	"github.com/golang-collections/collections/queue"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	node := createMinimalBST(arr)
	showTree(node, 0)
	fmt.Println(checkBST(node))
	fmt.Println(checkBST2(node, math.MinInt32))
	fmt.Println(checkBST3(node, math.MinInt32, math.MaxInt32))
}

func showTree(node *treeNode, depth int) {
	if node == nil {
		return
	}
	showTree(node.left, depth+1)
	fmt.Printf("value: %v, depth: %v \n", node.value, depth)
	showTree(node.right, depth+1)
}

type graph struct {
	nodes []*node
}

type node struct {
	adjacent []*node
	state    state
}

type state int

const (
	unvisited state = iota
	visited
	visiting
)

type treeNode struct {
	value  int
	left   *treeNode
	right  *treeNode
	parent *treeNode
}

// 4.1 ノード間の経路
func search(g *graph, start, end *node) bool {
	if start == end {
		return true
	}
	q := queue.New()
	for _, v := range g.nodes {
		v.state = unvisited
	}
	start.state = visiting
	q.Enqueue(start)

	var u *node
	for q.Len() != 0 {
		u = q.Dequeue().(*node)
		if u != nil {
			for _, v := range u.adjacent {
				if v.state == unvisited {
					if v == end {
						return true
					} else {
						v.state = visiting
						q.Enqueue(v)
					}
				}
			}
			u.state = visited
		}
	}
	return false
}

// 4.2 最小の木
func createMinimalBST(arr []int) *treeNode {
	return createMinimalBSTHelper(arr, 0, len(arr)-1)
}

func createMinimalBSTHelper(arr []int, start, end int) *treeNode {
	if end < start {
		return nil
	}
	mid := (start + end) / 2
	n := &treeNode{value: arr[mid]}
	n.left = createMinimalBSTHelper(arr, start, mid-1)
	n.right = createMinimalBSTHelper(arr, mid+1, end)
	return n
}

// 4.3 深さのリスト
func createLevelLinkedList(root *treeNode) []*list.List {
	lists := make([]*list.List, 0)
	return createLevelLinkedListHelper(root, lists, 0)
}

func createLevelLinkedListHelper(root *treeNode, lists []*list.List, level int) []*list.List {
	if root == nil {
		return lists
	}
	var l *list.List
	if len(lists) == level {
		l = list.New()
		lists = append(lists, l)
	} else {
		l = lists[level]
	}
	l.PushBack(root)
	lists = createLevelLinkedListHelper(root.left, lists, level+1)
	return createLevelLinkedListHelper(root.right, lists, level+1)
}

// 4.4
func isBalanced(root *treeNode) bool {
	if root == nil {
		return true
	}
	diff := getHeight(root.left) - getHeight(root.right)
	if math.Abs(diff) > 1 {
		return false
	} else {
		return isBalanced(root.left) && isBalanced(root.right)
	}
}

func getHeight(root *treeNode) float64 {
	if root == nil {
		return -1
	}
	return math.Max(getHeight(root.left), getHeight(root.right)) + 1
}

func isBalanced2(root *treeNode) bool {
	if _, err := checkHeight(root); err != nil {
		return false
	}
	return true
}

var notBalanced = errors.New("not balanced")

func checkHeight(root *treeNode) (float64, error) {
	if root == nil {
		return -1, nil
	}
	leftHeight, err := checkHeight(root.left)
	if err != nil {
		return 0, err
	}
	rightHeight, err := checkHeight(root.right)
	if err != nil {
		return 0, err
	}

	diff := leftHeight - rightHeight
	if math.Abs(diff) > 1 {
		return 0, notBalanced
	} else {
		return math.Max(leftHeight, rightHeight) + 1, nil
	}
}

// BSTチェック 4.5
func checkBST(root *treeNode) bool {
	arr := make([]int, 0)
	arr = copyBST(root, arr)
	for i := 1; i < len(arr); i++ {
		if arr[i] <= arr[i-1] {
			return false
		}
	}
	return true
}

func copyBST(root *treeNode, arr []int) []int {
	if root == nil {
		return arr
	}
	arr = copyBST(root.left, arr)
	arr = append(arr, root.value)
	arr = copyBST(root.right, arr)
	return arr
}

func checkBST2(root *treeNode, last int) (bool, int) {
	if root == nil {
		return true, last
	}
	result, last := checkBST2(root.left, last)
	if !result {
		return result, last
	}
	if root.value < last {
		return false, last
	}
	return checkBST2(root.right, root.value)
}

func checkBST3(root *treeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if !checkBST3(root.left, min, root.value) {
		return false
	}
	if !checkBST3(root.right, root.value, max) {
		return false
	}
	return true
}

// 4.6 次のノード
func inorderSucc(node *treeNode) *treeNode {
	if node == nil {
		return nil
	}
	if node.right != nil {
		return leftMostChild(node.right)
	} else {
		q := node
		x := q.parent
		for x != nil && x.left != q {
			q = x
			x = q.parent
		}
		return x
	}
}

func leftMostChild(node *treeNode) *treeNode {
	if node == nil {
		return nil
	}
	for node != nil {
		node = node.left
	}
	return node
}
