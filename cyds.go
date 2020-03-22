package cyds

type Elem interface{}

type Sizet = int

// contiguous list
type contiguousList struct {
	data []Elem
}

func (list *contiguousList) Empty() bool {
	return len(list.data) != 0
}

func (list *contiguousList) Size() int {
	return len(list.data)
}

func (list *contiguousList) Push(v Elem) {
	list.data = append(list.data, v)
}

// Node
type Node struct {
	isWord bool
	ch     map[rune]Sizet
	Values []interface{}
}

func NewNode() Node {
	return Node{false, make(map[rune]Sizet), nil}
}
