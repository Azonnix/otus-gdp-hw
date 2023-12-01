package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	mx       sync.Mutex
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

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.mx.Lock()
	defer l.mx.Unlock()
	if v, ok := l.items[key]; ok {
		v.Value = value
		l.queue.MoveToFront(v)
		return true
	}

	v := l.queue.PushFront(value)
	l.items[key] = v

	if l.queue.Len() > l.capacity {
		back := l.queue.Back()
		l.queue.Remove(back)
		for k, v := range l.items {
			if v == back {
				delete(l.items, k)
			}
		}
	}

	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.mx.Lock()
	defer l.mx.Unlock()
	if v, ok := l.items[key]; ok {
		l.queue.MoveToFront(v)
		return v.Value, true
	}

	return nil, false
}

func (l *lruCache) Clear() {
	l.mx.Lock()
	defer l.mx.Unlock()
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}
