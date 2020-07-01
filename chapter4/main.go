package main

import (
	"fmt"

	"github.com/golang-collections/collections/queue"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	node := createMinimalBST(arr)
	showTree(node, 0)
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
	value int
	left  *treeNode
	right *treeNode
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
