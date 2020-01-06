package cyDS

type Trie struct {
	Nodes []Node
	Size  Sizet
}

func (t *Trie) at(i Sizet) *Node {
	return &(t.Nodes[i])
}

func (t *Trie) Insert(str string) {
	s := []rune(str)
	if t.Size == 0 {
		t.Nodes = []Node{NewNode()}
		t.Size = 1
	}

	var cur Sizet = 0
	for i := 0; i < len(s); i++ {
		var c rune = s[i]

		nxt, ok := t.at(cur).ch[c]
		if !ok {
			t.Size++
			t.Nodes = append(t.Nodes, NewNode())
			t.at(cur).ch[c] = t.Size - 1
			cur = t.Size - 1
		} else {
			cur = nxt
		}
	}

	t.at(cur).isWord = true
}

func (t *Trie) Match(str string) bool {
	s := []rune(str)
	var cur Sizet = 0

	for i := 0; i < len(s); i++ {
		nxt, ok := t.at(cur).ch[s[i]]

		if ok {
			cur = nxt
		} else {
			return false
		}
	}

	return t.at(cur).isWord
}
