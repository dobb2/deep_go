package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type CircularQueue struct {
	values []int
	size   int
	front  int
	rear   int
}

func NewCircularQueue(size int) CircularQueue {
	return CircularQueue{
		values: make([]int, size),
		front:  -1,
		rear:   -1,
		size:   size,
	}
}

func (q *CircularQueue) Push(value int) bool {
	if q.Full() {
		return false
	}
	if q.front == -1 {
		q.front++
	}
	q.rear = (q.rear + 1) % q.size
	q.values[q.rear] = value

	return true
}

func (q *CircularQueue) Pop() bool {
	if q.Empty() {
		return false
	}
	if q.front == q.rear {
		q.front, q.rear = -1, -1
		return true
	}

	q.front = (q.front + 1) % q.size

	return true
}

func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}
	return q.values[q.front]
}

func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}
	return q.values[q.rear]
}

func (q *CircularQueue) Empty() bool {
	if q.front == -1 {
		return true
	}
	return false
}

func (q *CircularQueue) Full() bool {
	if q.front == 0 && q.rear == q.size+1 {
		return true
	}
	if q.front == (q.rear+1)%q.size {
		return true
	}
	return false
}

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Back())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queue.values))

	assert.False(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 3, queue.Back())

	assert.True(t, queue.Pop())
	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queue.values))

	assert.Equal(t, 2, queue.Front())
	assert.Equal(t, 4, queue.Back())

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

}

func TestCircularSecondQueue(t *testing.T) {
	const secondQueueSize = 5
	secondQueue := NewCircularQueue(secondQueueSize)

	assert.True(t, secondQueue.Push(5))
	assert.True(t, secondQueue.Push(4))
	assert.True(t, secondQueue.Push(3))

	assert.True(t, secondQueue.Pop())

	assert.True(t, reflect.DeepEqual([]int{5, 4, 3, 0, 0}, secondQueue.values))

	assert.True(t, secondQueue.Push(5))
	assert.True(t, secondQueue.Push(7))
	assert.True(t, secondQueue.Push(8))

	assert.False(t, secondQueue.Push(0))

	assert.True(t, reflect.DeepEqual([]int{8, 4, 3, 5, 7}, secondQueue.values))

	assert.True(t, secondQueue.Full())

	assert.Equal(t, 4, secondQueue.Front())
	assert.Equal(t, 8, secondQueue.Back())

	assert.True(t, secondQueue.Pop())
	assert.True(t, secondQueue.Pop())
	assert.True(t, secondQueue.Pop())
	assert.True(t, secondQueue.Pop())

	assert.Equal(t, 8, secondQueue.Front())
	assert.Equal(t, 8, secondQueue.Back())

	assert.True(t, secondQueue.Pop())

	assert.True(t, secondQueue.Empty())

	assert.True(t, secondQueue.Push(1))
	assert.True(t, secondQueue.Push(2))
	assert.True(t, secondQueue.Push(3))
	assert.True(t, secondQueue.Push(4))
	assert.True(t, secondQueue.Push(5))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3, 4, 5}, secondQueue.values))
	assert.Equal(t, 1, secondQueue.Front())
	assert.Equal(t, 5, secondQueue.Back())
}
