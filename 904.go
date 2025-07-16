package main

func totalFruit(fruits []int) int {
	ans := 0
	mp := make(map[int]int)
	l := 0
	for r, x := range fruits { // 优化为for-range
		mp[x]++
		for len(mp) > 2 { // 直接用len即可！
			mp[fruits[l]]--
			if mp[fruits[l]] == 0 {
				delete(mp, fruits[l])
			}
			l++
		}
		ans = max(ans, r-l+1)
	}
	return ans
}
