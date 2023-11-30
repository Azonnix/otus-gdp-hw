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
	// List         // Remove me after realization.
	lenght       int
	firstElement *ListItem
	backElement  *ListItem
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.lenght
}

func (l *list) Front() *ListItem {
	return l.firstElement
}

func (l *list) Back() *ListItem {
	return l.backElement
}

func (l *list) PushFront(v interface{}) *ListItem {
	if l.firstElement == nil {
		newListItem := &ListItem{Value: v, Next: nil, Prev: nil}
		l.firstElement = newListItem
		l.backElement = newListItem
	} else {
		newListItem := &ListItem{Value: v, Next: l.firstElement, Prev: nil}
		l.firstElement.Prev = newListItem
		l.firstElement = newListItem
	}

	l.lenght++
	return l.firstElement
}

func (l *list) PushBack(v interface{}) *ListItem {
	if l.backElement == nil {
		newListItem := &ListItem{Value: v, Next: nil, Prev: nil}
		l.firstElement = newListItem
		l.backElement = newListItem
	} else {
		newLIstItem := &ListItem{Value: v, Next: nil, Prev: l.backElement}
		l.backElement.Next = newLIstItem
		l.backElement = newLIstItem
	}

	l.lenght++
	return l.backElement
}

func (l *list) Remove(i *ListItem) {
	prevElement := i.Prev
	nextElement := i.Next

	if prevElement != nil {
		prevElement.Next = nextElement
	} else {
		l.firstElement = nextElement
	}

	if nextElement != nil {
		nextElement.Prev = prevElement
	} else {
		l.backElement = prevElement
	}

	l.lenght--
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}
