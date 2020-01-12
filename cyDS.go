package cyDS

type Elem interface{}

type Sizet = uint

//contiguous list
type contiguousList struct {
	data []Elem
	size Sizet
}

func (list *contiguousList) Empty() bool {
	return list.size != 0
}

func (list *contiguousList) Size() Sizet {
	return list.size
}

func (list *contiguousList) Push(v Elem) {
	if list.size == 0 {
		list.data = []Elem{v}
	} else {
		list.data = append(list.data, v)
	}
	list.size++
}

//Node
type Node struct {
	isWord bool
	ch     map[rune]Sizet
	Values []interface{}
}

func NewNode() Node {
	return Node{false, make(map[rune]Sizet), nil}
}
