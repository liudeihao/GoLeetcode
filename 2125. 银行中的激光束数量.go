package main

func numberOfBeams(bank []string) int {
	cnt := make([]int, 0)
	for _, s := range bank {
		curr := 0
		for _, v := range s {
			if v == '1' {
				curr++
			}
		}
		if curr > 0 {
			cnt = append(cnt, curr)
		}
	}
	ans := 0
	for i := 0; i < len(cnt)-1; i++ {
		ans += cnt[i] * cnt[i+1]
	}
	return ans
}
