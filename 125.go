package main

func keepAlphas(s string) string {
	var sb []byte
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			sb = append(sb, byte(c+32))
		} else if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			sb = append(sb, byte(c))
		}
	}
	return string(sb)
}

func isPalindrome(s string) bool {
	s = keepAlphas(s)
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
