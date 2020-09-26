package designAddAndSearchWordsDataStructure

type Node struct {
	c string
	d map[string]*Node
	e bool
}

type WordDictionary struct {
	d map[string]*Node
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	d := make(map[string]*Node)
	return WordDictionary{d}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	if len(word) > 0 {
		if _, ok := this.d[string(word[0])]; !ok {
			this.d[string(word[0])] = &Node{string(word[0]), make(map[string]*Node), false}
		}

		n := this.d[string(word[0])]
		for i := 1; i < len(word); i++ {
			if _, ok := n.d[string(word[i])]; !ok {
				n.d[string(word[i])] = &Node{string(word[i]), make(map[string]*Node), false}
			}
			n = n.d[string(word[i])]
		}
		n.e = true
	}
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	if string(word[0]) == "." {
		for _, v := range this.d {
			if this.search(*v, word) {
				return true
			}
		}
		return false
	} else if v, ok := this.d[string(word[0])]; ok {
		return this.search(*v, word)
	} else {
		return false
	}
}

func (this *WordDictionary) search(n Node, word string) bool {
	if n.c == string(word[0]) || string(word[0]) == "." {
		if len(word) == 1 {
			return n.e
		} else {
			if string(word[1]) == "." {
				for _, v := range n.d {
					if this.search(*v, string(word[1:])) {
						return true
					}
				}
				return false
			} else if v, ok := n.d[string(word[1])]; ok {
				return this.search(*v, string(word[1:]))
			} else {
				return false
			}
		}
	} else {
		return false
	}
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
