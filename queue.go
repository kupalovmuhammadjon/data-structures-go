package datastructuresgo

import "sync"

// Queue interface defines the methods that any queue implementation should have
type Queue interface {
	// Push adds a new element to the queue
	Push(val interface{})
	// Pop removes and returns the front element of the queue
	Pop() interface{}
	// Peek returns the front element without removing it
	Peek() interface{}
	// IsEmpty checks if the queue is empty
	IsEmpty() bool
	// Size returns the number of elements in the queue
	Size() int
}

type linkedQueue struct {
	mu    sync.Mutex
	front *Node
	rear  *Node
	len   int
}

// NewQueue creates and returns a new Queue
func NewQueue() Queue {
	return &linkedQueue{}
}

func (q *linkedQueue) Push(val interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()

	newNode := &Node{Val: val}

	if q.len == 0 {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.Next = newNode
		q.rear = newNode
	}

	q.len++
}

func (q *linkedQueue) Pop() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.len == 0 {
		return nil
	}

	val := q.front.Val
	q.front = q.front.Next
	q.len--

	if q.len == 0 {
		q.rear = nil
	}

	return val
}

func (q *linkedQueue) IsEmpty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.len == 0
}

func (q *linkedQueue) Peek() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.len == 0 {
		return nil
	}
	return q.front.Val
}

func (q *linkedQueue) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.len
}
