package main

func groupAnagrams(strs []string) [][]string {
    ansMap := make(map[[26]int][]string)
    for _, v := range strs {
        cnt := [26]int{}
        for _, c := range v {
            cnt[c-'a']++
        }
        ansMap[cnt] = append(ansMap[cnt], v)
    }
    var ans [][]string
    for _, v := range ansMap {
        ans = append(ans, v)
    }
    return ans
}
