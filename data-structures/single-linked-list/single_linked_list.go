// demonstration of singly linked list in golang

package main

import "fmt"

type node struct {
	val  int
	next *node
}

type SingleLinkedList struct {
	head *node
}

// to avoid mistakes when using pointer vs struct for new node creation
func newNode(val int) *node {
	return &node{val, nil}
}

func (ll *SingleLinkedList) addAtBeg(val int) {
	n := newNode(val)
	n.next = ll.head
	ll.head = n
}

func (ll *SingleLinkedList) addAtEnd(val int) {
	n := newNode(val)

	if ll.head == nil {
		ll.head = n
		return
	}

	cur := ll.head
	for ; cur.next != nil; cur = cur.next {
	}
	cur.next = n
}

func (ll *SingleLinkedList) delAtBeg() int {
	if ll.head == nil {
		return -1
	}

	cur := ll.head
	ll.head = cur.next

	return cur.val
}

func (ll *SingleLinkedList) delAtEnd() int {
	if ll.head == nil {
		return -1
	}

	if ll.head.next == nil {
		return ll.delAtBeg()
	}

	cur := ll.head

	for ; cur.next.next != nil; cur = cur.next {
	}

	retval := cur.next.val
	cur.next = nil
	return retval

}

func (ll *SingleLinkedList) count() int {
	var ctr int = 0

	for cur := ll.head; cur != nil; cur = cur.next {
		ctr += 1
	}

	return ctr
}

func (ll *SingleLinkedList) reverse() {
	var prev, next *node
	cur := ll.head

	for cur != nil {
		next = cur.next
		cur.next = prev
		prev = cur
		cur = next
	}

	ll.head = prev
}

func (ll *SingleLinkedList) display() {
	for cur := ll.head; cur != nil; cur = cur.next {
		fmt.Print(cur.val, " ")
	}

	fmt.Print("\n")
}

func main() {
	ll := SingleLinkedList{}

	ll.addAtBeg(10)
	ll.addAtEnd(20)
	ll.display()
	ll.addAtBeg(30)
	ll.display()
	ll.reverse()
	ll.display()

	fmt.Print(ll.delAtBeg(), "\n")
	ll.display()

	fmt.Print(ll.delAtEnd(), "\n")
	ll.display()

}
