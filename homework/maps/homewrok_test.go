package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type OrderedMap struct {
	node *ElementTree
	size int
}

type ElementTree struct {
	key   int
	value int
	left  *ElementTree
	right *ElementTree
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{
		size: 0,
	}
}

func (m *OrderedMap) Insert(key, value int) {
	m.node = m.insert(m.node, key, value)
	m.size++
}

func (m *OrderedMap) insert(node *ElementTree, key, value int) *ElementTree {
	if node == nil {
		return &ElementTree{
			key:   key,
			value: value,
		}
	} else if key < node.key {
		node.left = m.insert(node.left, key, value)
	} else if key > node.key {
		node.right = m.insert(node.right, key, value)
	}

	node.value = value
	return node
}

func (m *OrderedMap) min(node *ElementTree) (int, int) {
	if node.left == nil {
		return node.key, node.value
	}
	return m.min(node.left)
}

func (m *OrderedMap) delete(node *ElementTree, key int) *ElementTree {
	if node == nil {
		return node
	}
	if key < node.key {
		node.left = m.delete(node.left, key)
	} else if key > node.key {
		node.right = m.delete(node.right, key)
	} else if node.left != nil && node.right != nil {
		node.key, node.value = m.min(node.right)
		node.right = m.delete(node.right, node.key)
	} else {
		if node.left != nil {
			node = node.left
		} else if node.right != nil {
			node = node.right
		} else {
			node = nil
		}
	}
	return node
}

func (m *OrderedMap) Erase(key int) {
	m.node = m.delete(m.node, key)
	m.size--
}

func (m *OrderedMap) Contains(key int) bool {
	if m.search(m.node, key) != nil {
		return true
	}
	return false
}

func (m *OrderedMap) search(node *ElementTree, key int) *ElementTree {
	if node == nil {
		return node
	}
	if node.key == key {
		return node
	}
	if key < node.key {
		return m.search(node.left, key)
	} else {
		return m.search(node.right, key)
	}
}

func (m *OrderedMap) Size() int {
	return m.size
}

func (m *OrderedMap) inorderForEach(node *ElementTree, action func(int, int)) {
	if node != nil {
		m.inorderForEach(node.left, action)
		action(node.key, node.value)
		m.inorderForEach(node.right, action)
	}
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	m.inorderForEach(m.node, action)
}

func TestCircularQueue(t *testing.T) {
	data := NewOrderedMap()
	assert.Zero(t, data.Size())

	data.Insert(10, 10)
	data.Insert(5, 5)
	data.Insert(15, 15)
	data.Insert(2, 2)
	data.Insert(4, 4)
	data.Insert(12, 12)
	data.Insert(14, 14)

	assert.Equal(t, 7, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(3))
	assert.False(t, data.Contains(13))

	var keys []int
	expectedKeys := []int{2, 4, 5, 10, 12, 14, 15}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))

	data.Erase(15)
	data.Erase(14)
	data.Erase(2)

	assert.Equal(t, 4, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(2))
	assert.False(t, data.Contains(14))

	keys = nil
	expectedKeys = []int{4, 5, 10, 12}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))
}
