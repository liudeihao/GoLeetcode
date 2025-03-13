package main

func longestCommonPrefix(strs []string) string {
	s0 := strs[0]
	for i := range s0 {
		for _, s := range strs {
			if i >= len(s) {
				return s0[:i]
			}
			if s[i] != s0[i] {
				return s0[:i]
			}
		}
	}
	return s0
}
