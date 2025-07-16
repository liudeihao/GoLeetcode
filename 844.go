package main

func backspaceCompare(s string, t string) bool {
	sb, tb := []byte(s), []byte(t)
	raw2sting := func(bs []byte) string {
		slow, fast := 0, 0
		for fast < len(bs) {
			if bs[fast] == '#' {
				fast++
				if slow == 0 {
					continue
				} else {
					slow--
				}
			} else {
				bs[slow] = bs[fast]
				fast++
				slow++
			}
		}
		return string(bs[:slow])
	}
	return raw2sting(sb) == raw2sting(tb)
}
