package todolist

import "testing"

func TestListAdd(t *testing.T) {
	l := List{}
	l.Add(NewItem("Get milk"))

	if !l.Contains("Get milk") {
		t.Errorf("List should contain 'Get milk'")
	}

	if l.Contains("Get eggs") {
		t.Errorf("List should not contain 'Get eggs', but it does")
	}
}

