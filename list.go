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

func (l *List) Search(i Item) []Item {
	for _, item := range l.items {
		if strings.Contains(string(item), string(i)) {
			return []Item{item}
		}
	}
	return []Item{}
}

func (l *List) Equal(other *List) bool {
	return reflect.DeepEqual(*l, *other)
}

func (l *List) IsEmpty() bool {
	return len(l.items) == 0
}
