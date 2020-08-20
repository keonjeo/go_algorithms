// Binary search tree
// https://en.wikipedia.org/wiki/Binary_search_tree

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
	return &node{val, nil, nil}
}

func inOrder(n *node) {
	if n != nil {
		inOrder(n.left)
		fmt.Print(n.val, " ")
		inOrder(n.right)
	}
}

func insert(root *node, val int) *node {
	if root == nil {
		return newNode(val)
	}
	if val < root.val {
		root.left = insert(root.left, val)
	} else {
		root.right = insert(root.right, val)
	}
	return root
}

func inOrderSuccessor(root *node) *node {
	cur := root
	for cur.left != nil {
		cur = cur.left
	}
	return cur
}

func bstDelete(root *node, val int) *node {
	if root == nil {
		return nil
	}
	if val < root.val {
		root.left = bstDelete(root.left, val)
	} else if val > root.val {
		root.right = bstDelete(root.right, val)
	} else {
		// this is the node to delete

		// node with one child
		if root.left == nil {
			root = root.right
		} else if root.right == nil {
			root = root.left
		} else {
			// node with two children
			d := inOrderSuccessor(root)
			root.val = d.val
			root.right = bstDelete(root.right, d.val)
		}
	}
	return root
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
	t := &btree{nil}
	inOrder(t.root)
	t.root = insert(t.root, 10)
	t.root = insert(t.root, 20)
	t.root = insert(t.root, 15)
	t.root = insert(t.root, 30)
	fmt.Print(t.depth(), "\n")
	inOrder(t.root)
	fmt.Print("\n")
	t.root = bstDelete(t.root, 10)
	inOrder(t.root)
	fmt.Print("\n")
	t.root = bstDelete(t.root, 30)
	inOrder(t.root)
	fmt.Print("\n")
	t.root = bstDelete(t.root, 15)
	inOrder(t.root)
	fmt.Print("\n")
	t.root = bstDelete(t.root, 20)
	inOrder(t.root)
	fmt.Print("\n")
	fmt.Print(t.depth(), "\n")
}
