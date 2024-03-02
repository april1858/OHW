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

type CacheItem struct {
	Key   Key
	Value interface{}
}

type lruCache struct {
	mx       sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mx.Lock()
	defer c.mx.Unlock()
	item := &CacheItem{Key: key, Value: value}
	if i, ok := c.items[key]; ok {
		i.Value = *item
		c.queue.MoveToFront(i)
		c.items[key] = c.queue.Front()
		return true
	}
	if c.queue.Len() >= c.capacity {
		back := c.queue.Back()
		delete(c.items, back.Value.(CacheItem).Key)
		c.queue.Remove(back)
	}
	c.items[key] = c.queue.PushFront(*item)
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mx.Lock()
	defer c.mx.Unlock()
	if i, ok := c.items[key]; ok {
		c.queue.MoveToFront(i)
		return i.Value.(CacheItem).Value, true
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
