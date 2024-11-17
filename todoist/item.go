package todolist

type Item string

func NewItem(s string) Item {
	return Item(s)
}
