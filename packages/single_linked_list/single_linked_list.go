package packages

import (
	"fmt"
)

// Node represents a node in the linked list
type Node struct {
	data int
	next *Node
}

// LinkedList represents a linked list
type LinkedList struct {
	head *Node
}

// InsertAtBeginning inserts a new node at the beginning of the linked list
func (l *LinkedList) InsertAtBeginning(data int) {
	node := &Node{
		data: data,
		next: l.head,
	}
	l.head = node
}

// InsertAtEnd inserts a new node at the end of the linked list
func (l *LinkedList) InsertAtEnd(data int) {
	node := &Node{
		data: data,
	}
	if l.head == nil {
		l.head = node
		return
	}
	last := l.head
	for last.next != nil {
		last = last.next
	}
	last.next = node
}

// Display displays the linked list
func (l *LinkedList) Display() {
	head := l.head
	for head != nil {
		fmt.Printf("%d -> ", head.data)
		head = head.next
	}
	fmt.Println("nil")
}

// Delete deletes a node from the linked list
func (l *LinkedList) Delete(data int) {
	if l.head == nil {
		return
	}
	if l.head.data == data {
		l.head = l.head.next
		return
	}
	prev := l.head
	for prev.next != nil {
		if prev.next.data == data {
			prev.next = prev.next.next
			return
		}
		prev = prev.next
	}
}

// LinkedListDemo demonstrates the linked list operations
func LinkedListDemo() {
	l := LinkedList{}
	l.InsertAtBeginning(1)
	l.InsertAtBeginning(2)
	l.InsertAtBeginning(3)
	l.InsertAtEnd(4)
	l.InsertAtEnd(5)
	l.Display()
	l.Delete(3)
	l.Display()
}
