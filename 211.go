package main

type WordDictionary struct {
	ms map[string]bool
}

func Constructor() WordDictionary {
	return WordDictionary{ms: map[string]bool{}}
}

func (wd *WordDictionary) AddWord(word string) {
	wd.ms[word] = true
}

func (wd *WordDictionary) Search(word string) bool {
	return wd.SearchHelper(word, 0)
}

func (wd *WordDictionary) SearchHelper(word string, cur int) bool {
	p := cur
	for p < len(word) && word[p] != '.' {
		p++
	}
	if p == len(word) {
		_, ok := wd.ms[word]
		return ok
	}
	for i := 'a'; i <= 'z'; i++ {
		sb := []byte(word)
		sb[p] = byte(i)
		ok := wd.SearchHelper(string(sb), p+1)
		if ok {
			wd.ms[string(sb)] = true
			return true
		}
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
