package main

import "fmt"

/* v is the value of node; next is the pointer to next node */
type node struct {
	v int
	next *node
}

/* first node, called head. It points from first node to last node */
var head *node = nil

func (l *node) pushFront(val int) *node {
	/* if there's no nodes, head points to l (first node) */
	if head == nil {
		l.v = val
		l.next = nil
		head = l
		return l
	} else {
		/* create a new node equals to head */
		nnode := new(node)
		nnode = head
		/* create a second node with new value and `next -> nnode`
		*  is this way, nnode2 is before nnode
		*/
		nnode2 := &node {
			v: val,
			next: nnode,
		}
		/* now head is equals nnode2 */
		head = nnode2
		return head
	}
}

func (l *node) pushBack(val int) *node {
	/* if there's no nodes, head points to l (first node) */
	if head == nil {
		l.v = val
		l.next = nil
		head = l
		return l
	} else {
		/* read all list to last node */
		for l.next != nil {
			l = l.next
		}
		/* allocate a new portion of memory */
		l.next = new(node)
		l.next.v = val
		l.next.next = nil
		return l
	}
}

func (l *node) popFront() *node {
	if head == nil {
		return head
	}
	/* create a new node equals to first node pointed by head */
	cpnode := new(node)
	cpnode = head.next
	
	/* now head is equals cpnode (second node) */
	head = cpnode

	return head
}

func (l *node) popBack() *node {
	if head == nil {
		return head
	}
	/* create a new node equals to head */
	cpnode := new(node)
	cpnode = head
	
	/* read list to the penultimate node */
	for cpnode.next.next != nil {
		cpnode = cpnode.next
	}
	/* the penultimate node points to null. In this way the last node is deleted */
	cpnode.next = nil
	return head
}

func main() {
	listA := new(node)
	listA.pushBack(25).pushBack(24).pushBack(32) /* listA: 25 24 32 */
	listA.pushBack(56) /* listA: 25 24 32 56 */
	listA.pushFront(36) /* listA: 36 25 24 32 56 */
	listA.popFront() /* listA: 25 24 32 56 */
	listA.popBack() /* listA: 25 24 32 */
	
	/* read the list until head is not nil */
	for head != nil {
		fmt.Printf("%d ",head.v)
		head = head.next /*head points to next node */
	}
}
