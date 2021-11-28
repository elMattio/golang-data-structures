package linkedlists

import (
	"fmt"
)

type node struct {
	value int
	next  *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printListValues() {
	current := l.head
	iterator := l.length
	for iterator > 0 {
		fmt.Println(current.value)
		current = current.next
		iterator--
	}
}

func (l *linkedList) deleteWithValue(valueToDelete int) {
	if l.length == 0 {
		return
	}
	var previous *node = nil
	current := l.head
	found := false
	for !found {
		if current.value == valueToDelete {
			if previous == nil {
				l.head = current.next
			} else {
				previous.next = current.next
			}
			found = true
			l.length--
		} else {
			if current.next == nil {
				found = true
			} else {
				previous = current
				current = current.next
			}
		}
	}
}

func InitLinkedListExample() {
	myList := linkedList{}
	node1 := &node{value: 23}
	node2 := &node{value: 73}
	node3 := &node{value: 3}
	node4 := &node{value: 2}
	node5 := &node{value: 17}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.deleteWithValue(23)
	myList.printListValues()
}
