package main

type Trie_208 struct {
	children [26]*Trie_208
	isEnd    bool // 这是和一般多叉树不一样的地方（他们都是val int)
}

func Constructor_208() Trie_208 {
	return Trie_208{}
}

func (t *Trie_208) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie_208{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie_208) SearchPrefix(prefix string) *Trie_208 {
	node := t
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func (t *Trie_208) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isEnd
}

func (t *Trie_208) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
