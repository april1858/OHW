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
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	lenList := 0
	current := l.front
	for current != nil {
		lenList++
		current = current.Next
	}
	return lenList
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
	return newItem
}

func (l *list) Remove(i *ListItem) {
	if i.Next == nil {
		i.Prev.Next = nil
		l.back = i.Prev
	} else {
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}
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
