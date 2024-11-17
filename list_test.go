package todolist

import (
	"reflect"
	"testing"
)

func TestListAdd(t *testing.T) {
	l := NewList()
	l.Add(NewItem("Get milk"))

	if !l.Contains("Get milk") {
		t.Errorf("List should contain 'Get milk'")
	}

	if l.Contains("Get eggs") {
		t.Errorf("List should not contain 'Get eggs', but it does")
	}
}

func TestListSearch(t *testing.T) {
	l := NewList()
	l.Add("Get spinach")

	got := l.Search("spinach")
	want := []Item{"Get spinach"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

//	if !got.Equal(want) {
//		t.Errorf("got %v, want %v", got, want)
//	}

//	if !l.Search("milk").IsEmpty() {
//		t.Errorf("l.Search('milk') should return nil, got %v", l.Search("milk"))
//	}
}
