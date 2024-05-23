package queue_from_scratch

import (
	"errors"
)

type ListNode[T any] struct {
	Val  T
	Next *ListNode[T]
}

type Queue[T any] struct {
	Left  *ListNode[T]
	Right *ListNode[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		Left:  nil,
		Right: nil,
	}
}

func (q *Queue[T]) Enqueue(value T) {
	node := &ListNode[T]{
		Val:  value,
		Next: nil,
	}
	if q.Right == nil { // Queue is empty
		q.Left = node
		q.Right = node
	} else { // Queue is not empty
		q.Right.Next = node
		q.Right = q.Right.Next
	}
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.Left == nil {
		var zeroValue T
		return zeroValue, errors.New("queue is empty")
	}
	val := q.Left.Val
	q.Left = q.Left.Next
	if q.Left == nil {
		q.Right = nil
	}
	return val, nil
}
