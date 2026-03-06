package main

func minWindow(s string, t string) string {
    slen := len(s)
    diffMap := make(map[byte]int)
    for i := range t {
        diffMap[t[i]]--
    }
    diffCnt := 0
    for _, v := range diffMap {
        if v < 0 {
            diffCnt++
        }
    }

    var ans string

    l, r := 0, 0
    for r < slen {
        // 定义 [l, r) 为当前窗口
        for r < slen && diffCnt > 0 {
            newC := s[r]
            diffMap[newC]++
            if diffMap[newC] == 0 {
                diffCnt--
            }
            r++
        }
        for l < r && diffCnt == 0 {
            if len(ans) == 0 || r-l < len(ans) {
                ans = s[l:r]
            }
            oldC := s[l]
            diffMap[oldC]--
            if diffMap[oldC] == -1 {
                diffCnt++
            }
            l++
        }
    }
    return ans
}
