package queue

import (
	"container/list"
	"fmt"
)

type Queue[T any] struct {
	l *list.List
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{l: list.New()}
}

func (q *Queue[T]) Enqueue(value T) {
	q.l.PushBack(value)
}

func (q *Queue[T]) Dequeue() (T, error) {
	e := q.l.Front()
	if e == nil {
		var zero T
		return zero, fmt.Errorf("Queue is empty")
	}
	q.l.Remove(e)
	return e.Value.(T), nil
}

func (q *Queue[T]) IsEmpty() bool {
	return q.l.Len() == 0
}
