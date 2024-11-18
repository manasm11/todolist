package todolist

import (
	"reflect"
	"strings"
)

type List struct {
	todos []*Todo
}

func (l *List) Add(todo *Todo) {
	l.todos = append(l.todos, todo)
}

func NewList() *List {
	return &List{}
}

func (l *List) Contains(todo *Todo) bool {
	for _, item := range l.todos {
		if reflect.DeepEqual(item, todo) {
			return true
		}
	}
	return false
}

func (l *List) SearchText(str string) *List {
	l2 := NewList()
	for _, item := range l.todos {
		if strings.Contains(item.title, str) {
			l2.Add(item)
		}
	}
	return l2
}

func (l *List) Equal(other *List) bool {
	return reflect.DeepEqual(*l, *other)
}

func (l *List) IsEmpty() bool {
	return len(l.todos) == 0
}

func (l *List) Len() int {
	return len(l.todos)
}

func (l *List) Remove(item *Todo) {
	for i, item_ := range l.todos {
		if reflect.DeepEqual(item_, item) {
			l.todos = append(l.todos[:i], l.todos[i+1:]...)
			return
		}
	}
}
