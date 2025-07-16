package main

func isAnagram(s string, t string) bool {
	ms := make(map[rune]int)
	for _, c := range s {
		ms[c]++
	}
	for _, c := range t {
		ms[c]--
	}
	for _, v := range ms {
		if v != 0 {
			return false
		}
	}
	return true
}
