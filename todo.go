package todolist

type Todo struct {
	title string
	done bool
}

func NewTodo(s string) *Todo {
	return &Todo{title: s}
}

func (t *Todo) IsDone() bool {
	return t.done
}

func (t *Todo) MarkDone() {
	t.done = true
}
