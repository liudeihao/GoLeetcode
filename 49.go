package main

func groupAnagrams(strs []string) [][]string {
	ansMap := make(map[string][]string)
	for _, str := range strs {
		chs := make([]int, 26)
		for i := range str {
			ch := str[i] - 'a'
			chs[ch]++
		}
		var bs []byte
		for i, v := range chs {
			for ; v > 0; v-- {
				bs = append(bs, uint8(i)+'a')
			}
		}
		s := string(bs)
		ansMap[s] = append(ansMap[s], str)
	}
	var ans [][]string
	for _, v := range ansMap {
		ans = append(ans, v)
	}
	return ans
}
