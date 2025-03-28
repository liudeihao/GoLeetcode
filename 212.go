package main

type Trie struct {
	// 这里被改成了map
	children map[byte]*Trie
	word     string // 基于前缀树，学习一下。
}

func (t *Trie) Insert(s string) {
	// 还是用string吧
	node := t
	for i := range s {
		// 用ch:=s[i]而不是i,ch:=range s，因为range s得到的是rune
		ch := s[i]
		if node.children[ch] == nil {
			// 要注意初始化map了，不然panic
			node.children[ch] = &Trie{children: map[byte]*Trie{}}
		}
		node = node.children[ch]
	}
	node.word = s
}

// 这个我会
var dirs = []struct {
	x, y int
}{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}

func findWords(board [][]byte, words []string) []string {
	// 要注意初始化map了，不然panic
	t := &Trie{children: map[byte]*Trie{}}
	for _, word := range words {
		t.Insert(word)
	}
	m, n := len(board), len(board[0])
	var ans []string

	var dfs func(*Trie, int, int)
	dfs = func(node *Trie, x, y int) {
		// 确实形参取名x，y更好一些
		ch := board[x][y]
		nxt := node.children[ch]
		if nxt == nil {
			return
		}
		// 你看，这样就不用给Trie写Search了
		if nxt.word != "" {
			ans = append(ans, nxt.word)
			// 在前缀树里删掉已经找到过的单词，这很好！
			nxt.word = ""
		}

		// 也就是说没有孩子的话就不用再找了
		if len(nxt.children) > 0 {
			// 你看，这样就不用写visited了
			board[x][y] = '#'
			for _, d := range dirs {
				// 这个命名为nx，ny也很好
				nx, ny := x+d.x, y+d.y
				if nx < 0 || ny < 0 || nx >= m || ny >= n || board[nx][ny] == '#' {
					continue
				}
				// 而且这样不用每次从树顶遍历，还省去了你写的stack，这一点很重要！
				// 因为很重要所以我这里加了两行注释！
				dfs(nxt, nx, ny)
			}
			// 这就是回溯
			board[x][y] = ch
		}

		if len(nxt.children) == 0 {
			// 直接把nxt删了
			delete(node.children, ch)
		}
	}

	for i := range m {
		for j := range n {
			dfs(t, i, j)
		}
	}
	return ans
}
