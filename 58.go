package main

func lengthOfLastWord(s string) int {
	idx, ans := len(s)-1, 0
	for s[idx] == ' ' {
		idx--
	}
	for idx >= 0 && s[idx] != ' ' {
		ans++
		idx--
	}
	return ans
}
