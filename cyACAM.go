package cyACAM

type Node struct {
	isWord bool
	ch     map[rune]uint
}

func NewNode() Node {
	return Node{false, make(map[rune]uint)}
}

type ACAM struct {
	Nodes []Node
	Size  uint

	next []uint
	last []uint
}

func (ac *ACAM) at(i uint) *Node {
	return &(ac.Nodes[i])
}

func (ac *ACAM) Insert(str string) {
	s := []rune(str)
	if ac.Size == 0 {
		ac.Nodes = []Node{NewNode()}
		ac.Size=1
	}

	var cur uint = 0
	for i := 0; i < len(s); i++ {
		var c rune = s[i]

		nxt, ok := ac.at(cur).ch[c]
		if !ok {
			ac.Size++
			ac.Nodes = append(ac.Nodes, NewNode())
			ac.at(cur).ch[c] = ac.Size-1
			cur = ac.Size-1
		} else {
			cur = nxt
		}
	}

	ac.at(cur).isWord = true
}

func (ac *ACAM) Build() {
	ac.next = make([]uint, ac.Size)
	ac.last = make([]uint, ac.Size)

	var queue []uint

	for _, val := range ac.at(0).ch {
		queue = append(queue, val)
	}

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		for key, val := range ac.at(u).ch {
			queue = append(queue, val)
			w := ac.next[u]

			for _, ok := ac.at(w).ch[key]; w != 0 && !ok; {
				w = ac.next[w]
			}

			w, ok := ac.at(w).ch[key]
			if !ok {
				w = 0
			}
			ac.next[val] = w

			if ac.at(ac.next[val]).isWord {
				ac.last[val] = ac.next[val]
			} else {
				ac.last[val] = ac.last[ac.next[val]]
			}
		}
	}
}

func (ac *ACAM) Match(str string) uint {
	s := []rune(str)
	var cur uint = 0

	var match map[uint]bool = make(map[uint]bool)

	for i := 0; i < len(s); i++ {
		nxt, ok := ac.at(cur).ch[s[i]]
		for cur != 0 && !ok {
			cur = ac.next[cur]
			nxt, ok = ac.at(cur).ch[s[i]]
		}

		if ok {
			cur = nxt

			if ac.at(cur).isWord {
				tmp := cur
				for tmp != 0 {
					match[tmp] = true
					tmp = ac.last[tmp]
				}
			}
		}
	}

	return uint(len(match))
}