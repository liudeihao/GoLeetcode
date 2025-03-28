package main

type Trie struct {
	children [26]*Trie
	word     string // 基于前缀树，学习一下。
}

func (t *Trie) Insert(s string) {
	// 还是用string吧
	node := t
	for _, ch := range s {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
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
	t := &Trie{}
	for _, word := range words {
		t.Insert(word)
	}
	m, n := len(board), len(board[0])
	seen := map[string]bool{}

	var dfs func(*Trie, int, int)
	dfs = func(node *Trie, x, y int) {
		// 确实形参取名x，y更好一些
		ch := board[x][y]
		node = node.children[ch-'a']
		if node == nil {
			return
		}
		// 你看，这样就不用给Trie写Search了
		if node.word != "" {
			seen[node.word] = true
		}
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
			dfs(node, nx, ny)
		}
		// 这就是回溯
		board[x][y] = ch
	}

	for i := range m {
		for j := range n {
			dfs(t, i, j)
		}
	}
	ans := make([]string, 0, len(seen))
	for s := range seen {
		ans = append(ans, s)
	}
	return ans
}
