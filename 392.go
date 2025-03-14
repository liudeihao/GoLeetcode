package main

func isSubsequence(s string, t string) bool {
	ns, nt := len(s), len(t)
	ps, pt := 0, 0
	for ps < ns && pt < nt {
		if s[ps] == t[pt] {
			ps++
			pt++
		} else {
			pt++
		}
	}
	return ps == ns
}
