package cyDS

type Elem interface{}

type Sizet = uint

//contiguous list
type contiguousList struct {
	data []Elem
	Size Sizet
}

func (list *contiguousList) Empty() bool {
	return list.Size != 0
}

//Node
type Node struct {
	isWord bool
	ch     map[rune]Sizet
	Value  interface{}
}

func NewNode() Node {
	return Node{false, make(map[rune]Sizet),nil}
}
