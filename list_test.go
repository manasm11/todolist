package todolist

import (
	"testing"
)

func TestListAdd(t *testing.T) {
	t.Run("add one item", func(t *testing.T) {
		l := NewList()
		l.Add(NewItem("Get milk"))

		if !l.Contains("Get milk") {
			t.Errorf("List should contain 'Get milk'")
		}

		if l.Contains("Get eggs") {
			t.Errorf("List should not contain 'Get eggs', but it does")
		}
	})

	t.Run("add two items", func(t *testing.T) {
		l := NewList()
		l.Add("Get apples")
		l.Add("Get grapes")

		if l.Len() != 2 {
			t.Errorf("List should have 2 items, but it has %d", l.Len())
		}

		if !l.Contains("Get apples") {
			t.Errorf("List should contain 'Get apples'")
		}

		if !l.Contains("Get grapes") {
			t.Errorf("List should contain 'Get grapes'")
		}
	})
}

func TestListSearchText(t *testing.T) {
	t.Run("search in one items list", func(t *testing.T) {
		l := NewList()
		l.Add("Get spinach")
		l2 := NewList()
		l2.Add("Get spinach")

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

	t.Run("search in three items list", func(t *testing.T) {
		l := NewList()
		l.Add("Get spinach")
		l.Add("Get spinner")
		l.Add("Get milk")
		l2 := NewList()
		l2.Add("Get spinach")
		l2.Add("Get spinner")

		got := l.SearchText("spin")
		want := l2

		if !got.Equal(want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestRemoveItem(t *testing.T) {
	t.Run("remove item from list", func(t *testing.T) {
		l := NewList()
		l.Add("Get milk")
		l.Add("Get eggs")
		l.Remove("Get milk")
		if l.Contains("Get milk") {
			t.Errorf("List should not contain 'Get milk'")
		}

		if !l.Contains("Get eggs") {
			t.Errorf("List should contain 'Get eggs'")
		}
	})
}
