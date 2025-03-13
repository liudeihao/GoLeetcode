package main

func romanToInt(s string) int {
	m := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	n := len(s)
	result := m[s[n-1]]
	for i := n - 2; i >= 0; i-- {
		if m[s[i]] < m[s[i+1]] {
			result -= m[s[i]]
		} else {
			result += m[s[i]]
		}
	}
	return result
}
