package datastructuresgo

import (
	"fmt"
)

type doublyLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

type DoublyLinkedList interface {
	/*
		Append adds a new value to the end of the list.
		Time Complexity: O(1)
	*/
	Append(value any) *Node

	/*
		Prepend adds a new value to the start of the list.
		Time Complexity: O(1)
	*/
	Prepend(value any) *Node

	/*
		InsertAfter inserts a value after the given node.
		Time Complexity: O(1)
	*/
	InsertAfter(value any, mark *Node) *Node

	/*
		InsertBefore inserts a value before the given node.
		Time Complexity: O(1)
	*/
	InsertBefore(value any, mark *Node) *Node

	/*
		Remove removes the given node and returns its value.
		Time Complexity: O(1)
	*/
	Remove(node *Node) any

	/*
		Delete removes the first node with the given value.
		Time Complexity: O(n)
	*/
	Delete(value any)

	/*
		Find searches for a node with the given value and returns its pointer.
		Time Complexity: O(n)
	*/
	Find(value any) *Node

	/*
		MoveToFront moves the given node to the front (head) of the list.
		Time Complexity: O(1)
	*/
	MoveToFront(node *Node)

	/*
		MoveToBack moves the given node to the back (tail) of the list.
		Time Complexity: O(1)
	*/
	MoveToBack(node *Node)

	/*
		MoveAfter moves the node 'node' right after 'mark'.
		Time Complexity: O(1)
	*/
	MoveAfter(node *Node, mark *Node)

	/*
		MoveBefore moves the node 'node' right before 'mark'.
		Time Complexity: O(1)
	*/
	MoveBefore(node *Node, mark *Node)

	/*
		PushFrontList prepends another list to the front.
		Time Complexity: O(n)
	*/
	PushFrontList(other *doublyLinkedList)

	/*
		PushBackList appends another list to the back.
		Time Complexity: O(n)
	*/
	PushBackList(other *doublyLinkedList)

	/*
		Front returns the first node (head).
		Time Complexity: O(1)
	*/
	Front() *Node

	/*
		Back returns the last node (tail).
		Time Complexity: O(1)
	*/
	Back() *Node

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

func NewDoublyLinkedList() DoublyLinkedList {
	return &doublyLinkedList{}
}

func (dl *doublyLinkedList) Append(value any) *Node {
	newNode := &Node{
		Val: value,
	}

	if dl.Head == nil {
		dl.Head = newNode
		dl.Tail = newNode
	} else {
		dl.Tail.Next = newNode
		newNode.Prev = dl.Tail
		dl.Tail = newNode
	}

	dl.Size++
	return newNode
}

func (dl *doublyLinkedList) Prepend(value any) *Node {
	var (
		newNode = &Node{
			Val:  value,
			Next: dl.Head,
		}
	)

	if dl.Head != nil {
		newNode.Next = dl.Head
		dl.Head.Prev = newNode
	} else {
		dl.Tail = newNode
	}

	dl.Head = newNode
	dl.Size++

	return newNode
}

func (dl *doublyLinkedList) Find(value any) *Node {
	if dl.Head == nil {
		return nil
	}

	if dl.Head.Val == value {
		return dl.Head
	} else if dl.Tail.Val == value {
		return dl.Tail
	}

	current := dl.Head.Next

	for current != nil {
		if current.Val == value {
			return current
		}

		current = current.Next
	}

	return nil
}

func (dl *doublyLinkedList) Delete(value any) {
	if dl.Head == nil {
		return
	}

	if dl.Head.Val == value {
		dl.Head = dl.Head.Next
		if dl.Head == nil {
			dl.Tail = nil
		} else {
			dl.Head.Prev = nil
		}

		dl.Size--
		return
	}

	current := dl.Head.Next

	for current != nil {

		if current.Val == value {

			if current.Next != nil {
				current.Next.Prev = current.Prev
			}

			current.Prev.Next = current.Next
			dl.Size--
			return
		}

		current = current.Next
	}
}

/*
InsertAfter inserts a value after the given node.
Time Complexity: O(1)
*/
func (dl *doublyLinkedList) InsertAfter(value any, mark *Node) *Node {
	newNode := &Node{
		Val:  value,
		Prev: mark,
		Next: mark.Next,
	}

	if mark.Next != nil {
		mark.Next.Prev = newNode
	} else {
		dl.Tail = newNode
	}
	mark.Next = newNode

	dl.Size++
	return newNode
}

/*
InsertBefore inserts a value before the given node.
Time Complexity: O(1)
*/
func (dl *doublyLinkedList) InsertBefore(value any, mark *Node) *Node {
	var (
		newNode = &Node{
			Val:  value,
			Prev: mark.Prev,
			Next: mark,
		}
	)

	mark.Prev = newNode
	if newNode.Next != nil {
		newNode.Next.Prev = newNode
	}

	dl.Size++
	return newNode
}

/*
Remove removes the given node and returns its value.
Time Complexity: O(1)
*/
func (dl *doublyLinkedList) Remove(node *Node) any {
	if node == nil {
		return nil
	}

	if node == dl.Head {
		dl.Head = dl.Head.Next
		if dl.Head == nil {
			dl.Tail = nil
		}
	} else {

		if node.Prev != nil {
			node.Prev.Next = node.Next
		}
		if node.Next != nil {
			node.Next.Prev = node.Prev
		}

		if node == dl.Tail {
			dl.Tail = node.Prev
		}
	}

	node.Next = nil
	node.Prev = nil
	dl.Size--

	return node.Val
}

/*
MoveToFront moves the given node to the front (head) of the list.
Time Complexity: O(1)
*/
func (dl *doublyLinkedList) MoveToFront(node *Node) {
	if node == nil || node == dl.Head {
		return
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next

	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		dl.Tail = node.Prev
	}

	node.Prev = nil
	node.Next = dl.Head
	if dl.Head != nil {
		dl.Head.Prev = node
	}
	dl.Head = node

	if dl.Tail == nil {
		dl.Tail = node
	}
}

/*
MoveToBack moves the given node to the back (tail) of the list.
Time Complexity: O(1)
*/
func (dl *doublyLinkedList) MoveToBack(node *Node) {
	if node == nil || node == dl.Tail || dl.Head == nil {
		return
	}

	dl.Remove(node)

	dl.Tail.Next = node
	node.Prev = dl.Tail
	dl.Tail = node

	if dl.Head == nil {
		dl.Head = node
	}
}

/*
MoveAfter moves the node 'node' right after 'mark'.
Time Complexity: O(1)
*/
func (dl *doublyLinkedList) MoveAfter(node *Node, mark *Node) {
	if node == nil || mark == nil || node == mark {
		return
	}

	dl.Remove(node)

	node.Prev = mark
	node.Next = mark.Next

	if mark.Next != nil {
		mark.Next.Prev = node
	} else {
		dl.Tail = node
	}
	mark.Next = node
}

/*
MoveBefore moves the node 'node' right before 'mark'.
Time Complexity: O(1)
*/
func (dl *doublyLinkedList) MoveBefore(node *Node, mark *Node) {
	if node == nil || mark == nil || node == mark {
		return
	}

	dl.Remove(node)

	if mark.Prev != nil {
		mark.Prev.Next = node
	}
	node.Next = mark.Next
	node.Prev = mark.Prev

	node.Next = mark
	mark.Prev = node
}

/*
PushFrontList prepends another list to the front.
Time Complexity: O(n)
*/
func (dl *doublyLinkedList) PushFrontList(other *doublyLinkedList) {

}

/*
PushBackList appends another list to the back.
Time Complexity: O(n)
*/
func (dl *doublyLinkedList) PushBackList(other *doublyLinkedList) {

}

func (dl *doublyLinkedList) Length() int {
	return dl.Size
}

func (dl *doublyLinkedList) Print() {
	current := dl.Head
	for current != nil {
		fmt.Printf("%v -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

func (dl *doublyLinkedList) Back() *Node {
	return dl.Tail
}

func (dl *doublyLinkedList) Front() *Node {
	return dl.Head
}
