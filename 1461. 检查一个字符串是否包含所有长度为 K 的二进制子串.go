package main

func hasAllCodes(s string, k int) bool {
	subs := make(map[string]bool) // 可以存int来优化，从O(n*k)降到O(k)
	for i := 0; i+k <= len(s); i++ {
		subs[s[i:i+k]] = true
	}
	return len(subs) == 1<<k
}
