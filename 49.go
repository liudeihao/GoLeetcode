package main

func groupAnagrams(strs []string) [][]string {
	ansMap := make(map[[26]int][]string)
	for _, str := range strs {
		chs := [26]int{}
		for _, v := range str {
			chs[v]++
		}
		ansMap[chs] = append(ansMap[chs], str)
	}
	ans := make([][]string, 0, len(ansMap))
	for _, v := range ansMap {
		ans = append(ans, v)
	}
	return ans
}
