package queue

import (
	"sync"
)

// ItemQueue the queue of Items
type ItemQueue struct {
	items []interface{}
	lock  sync.RWMutex
}

// New creates a new ItemQueue
func New() *ItemQueue {
	return &ItemQueue{}
}

// Enqueue adds an Item to the end of the queue
func (s *ItemQueue) Enqueue(t interface{}) {
	s.lock.Lock()
	s.items = append(s.items, t)
	s.lock.Unlock()
}

// Dequeue removes an Item from the start of the queue
func (s *ItemQueue) Dequeue() *interface{} {
	s.lock.Lock()
	item := s.items[0]
	s.items = s.items[1:len(s.items)]
	s.lock.Unlock()
	return &item
}

// Front returns the item next in the queue, without removing it
func (s *ItemQueue) Front() *interface{} {
	s.lock.RLock()
	item := s.items[0]
	s.lock.RUnlock()
	return &item
}

// IsEmpty returns true if the queue is empty
func (s *ItemQueue) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of Items in the queue
func (s *ItemQueue) Size() int {
	return len(s.items)
}
