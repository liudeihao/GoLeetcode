package main

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Construct() Trie {
	return Trie{}
}

func (t *Trie) Insert(s []byte) {
	node := t
	for _, ch := range s {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) SearchPrefix(s []byte) *Trie {
	node := t
	for _, ch := range s {
		ch -= 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func (t *Trie) Search(s []byte) (isPrefix, isEnd bool) {
	node := t.SearchPrefix(s)
	return node != nil, node != nil && node.isEnd
}

func findWords(board [][]byte, words []string) []string {
	t := Construct()
	for _, word := range words {
		t.Insert([]byte(word))
	}
	m, n := len(board), len(board[0])
	all := map[string]bool{}
	dx := [...]int{-1, 0, 0, 1}
	dy := [...]int{0, -1, 1, 0}
	var v [12][12]bool
	for i := range m {
		for j := range n {
			s := []byte{board[i][j]}
			v[i][j] = true
			found := map[string]bool{}
			var dfs func(int, int, []byte, [12][12]bool)
			dfs = func(i, j int, s []byte, v [12][12]bool) {
				if found[string(s)] {
					return
				}
				isPrefix, isEnd := t.Search(s)
				if !isPrefix {
					return
				}
				if isEnd {
					found[string(s)] = true
				}
				for k := range 4 {
					x, y := i+dx[k], j+dy[k]
					if x < 0 || y < 0 || x >= m || y >= n || v[x][y] {
						continue
					}
					s0 := append(s, board[x][y])
					v[x][y] = true
					dfs(x, y, s0, v)
					v[x][y] = false
				}
			}
			dfs(i, j, s, v)
			v[i][j] = false
			for k := range found {
				all[k] = true
			}
		}
	}
	var ans []string
	for k := range all {
		ans = append(ans, k)
	}
	return ans
}
