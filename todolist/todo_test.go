package todolist

import "testing"

func TestTodoMarkDone(t *testing.T) {
	todo := NewTodo("test")

	if todo.IsDone() {
		t.Errorf("todo should not be marked as done when initialized")
	}

	todo.MarkDone()
	if !todo.IsDone() {
		t.Errorf("todo should be marked as done")
	}
}
