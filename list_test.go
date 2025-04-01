package todolist

import (
	"testing"
)

func TestListAdd(t *testing.T) {
	t.Run("add one todo", func(t *testing.T) {
		l := NewList()
		l.Add(NewTodo("Get milk"))

		if !l.Contains(NewTodo("Get milk")) {
			t.Errorf("List should contain 'Get milk'")
		}

		if l.Contains(NewTodo("Get eggs")) {
			t.Errorf("List should not contain 'Get eggs', but it does")
		}
	})

	t.Run("add two todos", func(t *testing.T) {
		l := createListWithItems("Get apples", "Get grapes")

		if l.Len() != 2 {
			t.Errorf("List should have 2 todos, but it has %d", l.Len())
		}

		if !l.Contains(NewTodo("Get apples")) {
			t.Errorf("List should contain 'Get apples'")
		}

		if !l.Contains(NewTodo("Get grapes")) {
			t.Errorf("List should contain 'Get grapes'")
		}
	})
}

func TestListSearchText(t *testing.T) {
	t.Run("search in one todos list", func(t *testing.T) {
		l := NewList()
		l.Add(NewTodo("Get spinach"))
		l2 := NewList()
		l2.Add(NewTodo("Get spinach"))

		got := l.SearchText("spinach")
		want := l2

		if !got.Equal(want) {
			t.Errorf("got %v, want %v", got, want)
		}

		got = l.SearchText("milk")
		if !got.IsEmpty() {
			t.Errorf("l.Search('milk') should return nil, got %v", got)
		}
	})

	t.Run("search in three todos list", func(t *testing.T) {
		l := createListWithItems("Get spinach", "Get spinner", "Get milk")
		l2 := createListWithItems("Get spinach", "Get spinner")

		got := l.SearchText("spin")
		want := l2

		if !got.Equal(want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestRemoveItem(t *testing.T) {
	t.Run("remove todo from list", func(t *testing.T) {
		l := createListWithItems("Get milk", "Get eggs")
		l.Remove(NewTodo("Get milk"))
		if l.Contains(NewTodo("Get milk")) {
			t.Errorf("List should not contain 'Get milk'")
		}

		if !l.Contains(NewTodo("Get eggs")) {
			t.Errorf("List should contain 'Get eggs'")
		}
	})
}

func createListWithItems(strs ...string) *List {
	l := NewList()
	for _, str := range strs {
		l.Add(NewTodo(str))
	}
	return l
}
