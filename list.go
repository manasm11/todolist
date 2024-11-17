package todolist

type List []Item

func (l *List) Add(i Item) {
	*l = append(*l, Item(i))
}

func NewList() *List {
	return &List{}
}

func (l List) Contains(i Item) bool {
	for _, item := range l {
		if item == i {
			return true
		}
	}
	return false
}
