package hw04lrucache

import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mx       sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
	elements map[interface{}]Key
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
		elements: make(map[interface{}]Key, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mx.Lock()
	defer c.mx.Unlock()
	if i, ok := c.items[key]; ok {
		i.Value = value
		c.queue.MoveToFront(i)
		c.items[key] = c.queue.Front()
		return true
	}
	if c.queue.Len() >= c.capacity {
		back := c.queue.Back()
		delete(c.items, c.elements[back.Value])
		delete(c.elements, back.Value)
		c.queue.Remove(back)
	}
	c.elements[value] = key
	c.items[key] = c.queue.PushFront(value)
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mx.Lock()
	defer c.mx.Unlock()
	if i, ok := c.items[key]; ok {
		c.queue.MoveToFront(i)
		return i.Value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.mx.Lock()
	defer c.mx.Unlock()
	for i := range c.items {
		delete(c.items, i)
	}
	c.queue = NewList()
}
