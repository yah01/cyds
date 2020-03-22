package cyds

type ACAM struct {
	Trie

	next []Sizet
	last []Sizet
}

//after all Insert finishing, call Build
func (ac *ACAM) Build() {
	var queue Queue

	for _, val := range ac.at(0).ch {
		queue.Push(val)
	}

	for !queue.Empty() {
		u := queue.Pop().(Sizet)

		for key, val := range ac.at(u).ch {
			queue.Push(val)
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

func (ac *ACAM) Match(str string) Sizet {
	s := []rune(str)
	var cur Sizet = 0

	var match map[Sizet]bool = make(map[Sizet]bool)

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

	return Sizet(len(match))
}
