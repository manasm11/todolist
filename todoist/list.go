package todolist

import (
	"reflect"
	"strings"
)

type List struct {
	items []Item
}

func (l *List) Add(i Item) {
	l.items = append(l.items, Item(i))
}

func NewList() *List {
	return &List{}
}

func (l *List) Contains(i Item) bool {
	for _, item := range l.items {
		if item == i {
			return true
		}
	}
	return false
}

func (l *List) SearchText(str string) *List {
	l2 := NewList()
	for _, item := range l.items {
		if strings.Contains(string(item), str) {
			l2.Add(item)
		}
	}
	return l2
}

func (l *List) Equal(other *List) bool {
	return reflect.DeepEqual(*l, *other)
}

func (l *List) IsEmpty() bool {
	return len(l.items) == 0
}

func (l *List) Len() int {
	return len(l.items)
}

func (l *List) Remove(item Item) {
	for i, item_ := range l.items {
		if item_ == item {
			l.items = append(l.items[:i], l.items[i+1:]...)
			return
		}
	}
}
