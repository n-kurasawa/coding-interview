package main

import (
	"fmt"

	"github.com/golang-collections/collections/queue"
)

func main() {
	node1 := &node{}
	node2 := &node{}
	node3 := &node{}
	node4 := &node{}
	node5 := &node{}
	node6 := &node{}

	node1.adjacent = []*node{node2, node4}
	node4.adjacent = []*node{node1, node6}
	node3.adjacent = []*node{node5}

	g := &graph{
		nodes: []*node{node1, node2, node3, node4, node5, node6},
	}
	fmt.Println(search(g, node1, node6))
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
