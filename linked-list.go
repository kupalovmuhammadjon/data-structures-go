package datastructuresgo

import "fmt"

type linkedList struct {
	Head *Node
	Tail *Node
	Size int
}

type LinkedList interface {
	/*
		Append adds a new value to the end of the list.
		Time Complexity: O(1)
	*/
	Append(value any)
	/*
		Prepend adds a new value to the start of the list.
		Time Complexity: O(1)
	*/
	Prepend(value any)
	/*
		Find searches for a node with the given value and returns its pointer.
		Time Complexity: O(n)
	*/
	Find(value any) *Node
	/*
		Delete removes the first node with the given value.
		Time Complexity: O(n)
	*/
	Delete(value any)
	/*
		Length returns the number of elements in the list.
		Time Complexity: O(1)
	*/
	Length() int
	/*
		Print displays the values in the linked list in order.
		Time Complexity: O(n)
	*/
	Print()
}

// NewLinkedList initializes and returns a new empty linked list.
func NewLinkedList() LinkedList {
	return &linkedList{}
}

func (l *linkedList) Append(value any) {

	var (
		newNode = Node{
			Val: value,
		}
	)

	if l.Head == nil {
		l.Head = &newNode
		l.Tail = &newNode
	} else {
		l.Tail.Next = &newNode
		l.Tail = &newNode
	}

	l.Size++
}

func (l *linkedList) Prepend(value any) {
	var (
		newNode = Node{
			Val:  value,
			Next: l.Head,
		}
	)

	l.Head = &newNode
	if l.Tail == nil {
		l.Tail = &newNode
	}

	l.Size++
}

func (l *linkedList) Find(value any) *Node {
	if l.Head == nil {
		return nil
	}

	if l.Head.Val == value {
		return l.Head
	} else if l.Tail.Val == value {
		return l.Tail
	}

	current := l.Head.Next

	for current != nil {
		if current.Val == value {
			return current
		}

		current = current.Next
	}

	return nil
}

func (l *linkedList) Delete(value any) {
	if l.Head == nil {
		return
	}

	if l.Head.Val == value {
		l.Head = l.Head.Next
		if l.Head == nil {
			l.Tail = nil
		}

		l.Size--
		return
	}

	current := l.Head.Next
	prev := l.Head

	for current != nil {

		if current.Val == value {
			prev.Next = current.Next
			l.Size--
			return
		}

		prev = current
		current = current.Next
	}
}

func (l *linkedList) Length() int {
	return l.Size
}

func (l *linkedList) Print() {
	current := l.Head
	for current != nil {
		fmt.Printf("%v -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}
