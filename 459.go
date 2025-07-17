package main

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	var factors []int
	for i := 1; i <= n/2; i++ {
		if n == n/i*i {
			factors = append(factors, i)
		}
	}

	for _, sublen := range factors {
		ans := true
		s0 := s[0:sublen]
		for i := sublen; i < n; i += sublen {
			si := s[i : i+sublen]
			if s0 != si {
				ans = false
				break
			}
		}
		if ans {
			return true
		}
	}
	return false
}
