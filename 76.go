package main

func minWindow(s string, t string) string {
	m, n := len(s), len(t)
	if m < n {
		return ""
	}

	var (
		ans      string
		minLen   = m + 1
		tMap     = map[byte]int{}
		curMap   = map[byte]int{}
		tCount   int
		curCount int
	)
	for i := range t {
		tMap[t[i]]++
	}
	tCount = len(tMap)
	for left, right := 0, 0; right < m; right++ {
		ch := s[right]
		curMap[ch]++
		if curMap[ch] == tMap[ch] {
			curCount++
		}
		for curCount == tCount {
			ch := s[left]
			v, ok := tMap[ch]
			// 是t里的，而且正好满足条件
			if ok && curMap[ch] >= v {
				if right-left+1 < minLen {
					ans = s[left : right+1]
					minLen = right - left + 1
				}
				// 因为left要++了
				curMap[ch]--
				if curMap[ch] < v {
					curCount--
				}
			}
			left++
		}
	}
	return ans
}
