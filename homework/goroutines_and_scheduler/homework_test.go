package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Task struct {
	Identifier int
	Priority   int
}

type Scheduler struct {
	heap []Task
}

func (s *Scheduler) Parent(i int) int {
	return (i - 1) / 2
}

func (s *Scheduler) LeftChild(i int) int {
	return 2*i + 1
}

func (s *Scheduler) RightChild(i int) int {
	return 2*i + 2
}

func (s *Scheduler) SiftUp(i int) {
	for i > 0 && s.heap[s.Parent(i)].Priority < s.heap[i].Priority {
		s.heap[i], s.heap[s.Parent(i)] = s.heap[s.Parent(i)], s.heap[i]
		i = s.Parent(i)
	}
}

func (s *Scheduler) SiftDown(i int) {
	minIndex := i
	left := s.LeftChild(i)
	if left <= len(s.heap)-1 && s.heap[left].Priority > s.heap[minIndex].Priority {
		minIndex = left
	}
	right := s.RightChild(i)
	if right <= len(s.heap)-1 && s.heap[right].Priority > s.heap[minIndex].Priority {
		minIndex = right
	}
	if i != minIndex {
		s.heap[i], s.heap[minIndex] = s.heap[minIndex], s.heap[i]
		s.SiftDown(minIndex)
	}
}

func NewScheduler() Scheduler {
	return Scheduler{heap: make([]Task, 0, 0)}
}

func (s *Scheduler) AddTask(task Task) {
	s.heap = append(s.heap, task)

	s.SiftUp(len(s.heap) - 1)
}

func (s *Scheduler) ChangeTaskPriority(taskID int, newPriority int) {
	for i, task := range s.heap {
		if task.Identifier == taskID {
			oldPriority := task.Priority
			s.heap[i].Priority = newPriority
			if newPriority > oldPriority {
				s.SiftUp(i)
			} else {
				s.SiftDown(i)
			}
			return
		}
	}
}

func (s *Scheduler) GetTask() Task {
	task := s.heap[0]
	s.heap[0] = s.heap[len(s.heap)-1]
	s.heap = s.heap[:len(s.heap)-1]
	s.SiftDown(0)
	return task
}

func TestTrace(t *testing.T) {
	task1 := Task{Identifier: 1, Priority: 10}
	task2 := Task{Identifier: 2, Priority: 20}
	task3 := Task{Identifier: 3, Priority: 30}
	task4 := Task{Identifier: 4, Priority: 40}
	task5 := Task{Identifier: 5, Priority: 50}

	scheduler := NewScheduler()
	scheduler.AddTask(task1)
	scheduler.AddTask(task2)
	scheduler.AddTask(task3)
	scheduler.AddTask(task4)
	scheduler.AddTask(task5)

	task := scheduler.GetTask()
	assert.Equal(t, task5, task)

	task = scheduler.GetTask()
	assert.Equal(t, task4, task)

	scheduler.ChangeTaskPriority(1, 100)
	task1.Priority = 100

	task = scheduler.GetTask()
	assert.Equal(t, task1, task)

	task = scheduler.GetTask()
	assert.Equal(t, task3, task)
}
