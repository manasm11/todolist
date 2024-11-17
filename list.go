package todolist

import (
	"reflect"
	"strings"
)

type List []Item

func (l *List) Add(i Item) {
	*l = append(*l, Item(i))
}

func NewList() *List {
	return &List{}
}

func (l *List) Contains(i Item) bool {
	for _, item := range *l {
		if item == i {
			return true
		}
	}
	return false
}

func (l *List) Search(i Item) []Item {
	for _, item := range *l {
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
	return len(*l) == 0
}
