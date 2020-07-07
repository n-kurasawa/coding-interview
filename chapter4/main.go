package main

import (
	"container/list"
	"errors"
	"fmt"
	"math"

	"github.com/golang-collections/collections/stack"

	"github.com/golang-collections/collections/queue"
)

func main() {
	arr := []string{"a", "b", "c", "d", "e", "f"}
	d := [][]string{{"d", "a"}, {"b", "f"}, {"d", "b"}, {"a", "f"}, {"c", "d"}}
	fmt.Println(findBuildOrder(arr, d))
	fmt.Println(findBuildOrder2(arr, d))
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

// 4.7 実行順序
type projectGraph struct {
	nodes []*project
	m     map[string]*project
}

func newProjectGraph(projects []string, dependencies [][]string) *projectGraph {
	g := &projectGraph{m: map[string]*project{}}
	for _, v := range projects {
		g.createNode(v)
	}
	for _, v := range dependencies {
		first := v[0]
		second := v[1]
		g.addEdge(second, first)
	}
	return g
}

func (g *projectGraph) createNode(name string) *project {
	if g.m[name] == nil {
		p := newProject(name)
		g.m[name] = p
		g.nodes = append(g.nodes, p)
	}
	return g.m[name]
}

func (g *projectGraph) addEdge(parent, child string) {
	p := g.m[parent]
	c := g.m[child]
	p.addChild(c)
}

type project struct {
	children     []*project
	m            map[string]*project
	name         string
	dependencies int
	state        state
}

func newProject(name string) *project {
	return &project{
		name:  name,
		m:     map[string]*project{},
		state: unvisited,
	}
}

func (p *project) String() string {
	arr := make([]string, len(p.children))
	for i, v := range p.children {
		arr[i] = v.name
	}
	return fmt.Sprintf("{name: %v, children: %v, dependencies: %v}", p.name, arr, p.dependencies)
}

func (p *project) addChild(child *project) {
	if p.m[child.name] == nil {
		p.children = append(p.children, child)
		p.m[child.name] = child
		child.dependencies++
	}
}

func findBuildOrder(projects []string, dependencies [][]string) []string {
	g := newProjectGraph(projects, dependencies)
	ps := orderProjects(g.nodes)
	var arr []string
	for _, v := range ps {
		arr = append(arr, v.name)
	}
	return arr
}

func orderProjects(projects []*project) []*project {
	order := make([]*project, len(projects))
	endOfList := addNonDependent(order, projects, 0)

	var toBeProcessed int
	for toBeProcessed < len(order) {
		current := order[toBeProcessed]

		if current == nil {
			return nil
		}

		for _, v := range current.children {
			v.dependencies--
		}

		endOfList = addNonDependent(order, current.children, endOfList)
		toBeProcessed++
	}
	return order
}

func addNonDependent(order, projects []*project, offset int) int {
	for _, v := range projects {
		if v.dependencies == 0 {
			order[offset] = v
			offset++
		}
	}
	return offset
}

func findBuildOrder2(projects []string, dependencies [][]string) []string {
	g := newProjectGraph(projects, dependencies)
	s := orderProjects2(g.nodes)
	if s == nil {
		return nil
	}

	arr := make([]string, s.Len())
	pro := s.Pop()
	for i := 0; pro != nil; i++ {
		arr[i] = pro.(*project).name
		pro = s.Pop()
	}
	return arr
}

func orderProjects2(projects []*project) *stack.Stack {
	stk := stack.New()
	for _, v := range projects {
		if v.state == unvisited {
			if !doDFS(v, stk) {
				return nil
			}
		}
	}
	return stk
}

func doDFS(project *project, stk *stack.Stack) bool {
	if project.state == visiting {
		return false
	}

	if project.state == unvisited {
		project.state = visiting
		for _, v := range project.children {
			if !doDFS(v, stk) {
				return false
			}
		}
		project.state = visited
		stk.Push(project)
	}
	return true
}

// 4.8 最初の共通祖先
func commonAncestor(p, q *treeNode) *treeNode {
	delta := depth(p) - depth(q)
	var first, second *treeNode
	if delta > 0 {
		first = q
		second = p
	} else {
		first = p
		second = q
	}
	second = goUpBy(second, int(math.Abs(float64(delta))))

	for first != second && first != nil && second != nil {
		first = first.parent
		second = second.parent
	}
	if first == nil || second == nil {
		return nil
	} else {
		return first
	}
}

func depth(node *treeNode) int {
	var dep int
	for node != nil {
		node = node.parent
		dep++
	}
	return dep
}

func goUpBy(node *treeNode, delta int) *treeNode {
	for delta > 0 && node != nil {
		node = node.parent
		delta--
	}
	return node
}

func commonAncestor2(root, p, q *treeNode) *treeNode {
	if !covers(root, p) || !covers(root, q) {
		return nil
	} else if covers(p, q) {
		return p
	} else if covers(q, p) {
		return q
	}

	sibling := getSibling(p)
	parent := p.parent
	for !covers(sibling, q) {
		sibling = getSibling(parent)
		parent = parent.parent
	}
	return parent
}

func covers(root, node *treeNode) bool {
	if root == nil {
		return false
	}
	if root == node {
		return true
	}
	return covers(root.left, node) || covers(root.right, node)
}

func getSibling(node *treeNode) *treeNode {
	if node == nil || node.parent == nil {
		return nil
	}
	parent := node.parent
	if parent.left == node {
		return parent.right
	} else {
		return parent.left
	}
}

func commonAncestor3(root, p, q *treeNode) *treeNode {
	if !covers(root, p) || !covers(root, q) {
		return nil
	}
	return ancestorHelper(root, p, q)
}

func ancestorHelper(root, p, q *treeNode) *treeNode {
	if root == nil || root == p || root == q {
		return root
	}

	pIsLeft := covers(root.left, p)
	qIsLeft := covers(root.left, q)
	if pIsLeft != qIsLeft {
		return root
	}
	var childSide *treeNode
	if pIsLeft {
		childSide = root.left
	} else {
		childSide = root.right
	}
	return ancestorHelper(childSide, p, q)
}

var ancestorError = errors.New("not ancestor")

func commonAncestor4(root, p, q *treeNode) (*treeNode, error) {
	if root == nil {
		return nil, ancestorError
	}

	if root == p && root == q {
		return root, nil
	}

	x, xErr := commonAncestor4(root.left, p, q)
	if x != nil && x != p && x != q {
		return x, xErr
	}

	y, yErr := commonAncestor4(root.right, p, q)
	if y != nil && y != p && y != q {
		return y, yErr
	}

	if x != nil && y != nil {
		return root, nil
	} else if root == p || root == y {
		if x != nil || y != nil {
			return root, nil
		} else {
			return root, ancestorError
		}
	} else {
		if x == nil {
			return y, ancestorError
		} else {
			return x, ancestorError
		}
	}
}
