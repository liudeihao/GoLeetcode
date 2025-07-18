package main

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	d2s := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	n := len(digits)
	var ans []string
	var path []rune
	var backtrack func(idx int)
	backtrack = func(idx int) {
		if len(path) == n {
			ans = append(ans, string(path))
			return
		}
		for _, r := range d2s[digits[idx]] {
			path = append(path, r)
			backtrack(idx + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return ans
}
