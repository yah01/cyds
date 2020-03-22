package cyds

type Trie struct {
	Nodes []Node
}

func (t *Trie) at(i int) *Node {
	return &(t.Nodes[i])
}

func (t *Trie) Insert(str string, val interface{}) {
	if t.Nodes == nil {
		t.Nodes = []Node{NewNode()}
	}

	var (
		s   = []rune(str)
		cur = 0
	)

	for i := 0; i < len(s); i++ {
		c := s[i]

		if nxt, ok := t.at(cur).ch[c]; ok {
			cur = nxt
		} else {
			t.Nodes = append(t.Nodes, NewNode())
			t.at(cur).ch[c] = len(t.Nodes) - 1
			cur = len(t.Nodes) - 1
		}
	}

	t.at(cur).isWord = true
	t.at(cur).Values = append(t.at(cur).Values, val)
}

func (t *Trie) Match(str string) (*Node, bool) {
	var (
		s   = []rune(str)
		cur = 0
	)

	for i := 0; i < len(s); i++ {
		if nxt, ok := t.at(cur).ch[s[i]]; ok {
			cur = nxt
		} else {
			return nil, false
		}
	}

	return t.at(cur), true
}
