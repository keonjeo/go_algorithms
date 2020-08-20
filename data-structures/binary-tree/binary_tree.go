// basic binary tree and related operations

package main

import "fmt"

type node struct {
	val   int
	left  *node
	right *node
}

type btree struct {
	root *node
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func newNode(val int) *node {
	n := &node{val, nil, nil}
	return n
}

func inOrder(n *node) {
	if n == nil {
		return
	}
	inOrder(n.left)
	fmt.Print(n.val, " ")
	inOrder(n.right)
}

func preOrder(n *node) {
	if n == nil {
		return
	}
	fmt.Print(n.val, " ")
	preOrder(n.left)
	preOrder(n.right)
}

func postOrder(n *node) {
	if n == nil {
		return
	}
	postOrder(n.left)
	postOrder(n.right)
	fmt.Print(n.val, " ")
}

func levelOrder(root *node) {
	var q []*node // queue
	var n *node   // temporary node

	q = append(q, root)

	for len(q) != 0 {
		n, q = q[0], q[1:]
		fmt.Print(n.val, " ")
		if n.left != nil {
			q = append(q, n.left)
		}
		if n.right != nil {
			q = append(q, n.right)
		}
	}
}

// helper function for t.depth
func _calculate_depth(n *node, depth int) int {
	if n == nil {
		return depth
	}
	return max(_calculate_depth(n.left, depth+1), _calculate_depth(n.right, depth+1))
}

func (t *btree) depth() int {
	return _calculate_depth(t.root, 0)
}


func main() {
	t := btree{nil}
	t.root = newNode(0)
	t.root.left = newNode(1)
	t.root.right = newNode(2)
	t.root.left.left = newNode(3)
	t.root.left.right = newNode(4)
	t.root.right.left = newNode(5)
	t.root.right.right = newNode(6)
	t.root.right.right.right = newNode(10)

	inOrder(t.root)
	fmt.Print("\n")
	preOrder(t.root)
	fmt.Print("\n")
	postOrder(t.root)
	fmt.Print("\n")
	levelOrder(t.root)
	fmt.Print("\n")
	fmt.Print(t.depth(), "\n")
}

