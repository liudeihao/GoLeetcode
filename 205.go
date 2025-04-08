package main

func isIsomorphic(s string, t string) bool {
	s2t := make(map[uint8]uint8)
	t2s := make(map[uint8]uint8)
	for i := range s {
		sch, tch := s[i], t[i]
		if s2t[sch] == 0 && t2s[tch] == 0 {
			s2t[sch] = tch
			t2s[tch] = sch
		} else if s2t[sch] != tch || t2s[tch] != sch {
			return false
		}
	}
	return true
}
