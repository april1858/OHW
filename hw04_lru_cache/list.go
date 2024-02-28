package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	front *ListItem
	back  *ListItem
	len   int
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{Value: v, Next: nil, Prev: nil}
	if l.front == nil {
		l.front = newItem
		l.back = newItem
	} else {
		current := l.front
		newItem.Next = current
		current.Prev = newItem
		l.front = newItem
	}
	l.len++
	return newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{Value: v, Next: nil, Prev: nil}
	if l.front == nil {
		l.front = newItem
		l.back = newItem
	} else {
		current := l.back
		newItem.Prev = current
		current.Next = newItem
		l.back = newItem
	}
	l.len++
	return newItem
}

func (l *list) Remove(i *ListItem) {
	if i.Next == nil {
		if i.Prev == nil {
			l.front = nil
			l.back = nil
			return
		}
		i.Prev.Next = nil
		l.back = i.Prev
	} else {
		if i.Prev == nil {
			i.Next.Prev = nil
			return
		}
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Prev == nil {
		return
	}
	if i.Next == nil {
		i.Prev.Next = nil
		l.back = i.Prev
	} else {
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}
	l.front = l.PushFront(i.Value)
}
