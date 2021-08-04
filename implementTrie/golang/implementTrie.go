package implementTrie

type Trie struct {
	b byte
	s []*Trie
	t bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{'.', make([]*Trie, 26), false}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {

	cur := this
	for i := 0; i < len(word); i++ {
		if cur.s[word[i]-97] == nil {
			cur.s[word[i]-97] = &Trie{word[i], make([]*Trie, 26), false}
		}

		cur = cur.s[word[i]-97]
	}

	cur.t = true
}

func (this *Trie) Search(word string) bool {
	cur := this
	for i := 0; i < len(word); i++ {
		//fmt.Printf("11111\n")
		if cur.s[word[i]-97] == nil {
			//fmt.Printf("22222 %d\n", i)
			return false
		}
		cur = cur.s[word[i]-97]
	}

	return cur.t
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for i := 0; i < len(prefix); i++ {
		//fmt.Printf("11111\n")
		if cur.s[prefix[i]-97] == nil {
			return false
		}
		cur = cur.s[prefix[i]-97]
	}

	//fmt.Printf("22222\n")
	return true
}
