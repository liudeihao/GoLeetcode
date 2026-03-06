package main

func findAnagrams(s string, p string) []int {
    slen, plen := len(s), len(p)
    if slen < plen {
        return []int{}
    }
    diff := [26]int{} // diff[i] 记录 cur_s 第i个字符数量 - p中第i个字符数量
    for i := range p {
        diff[s[i]-'a']++
        diff[p[i]-'a']--
    }

    diffCnt := 0
    for i := range diff {
        if diff[i] != 0 {
            diffCnt++
        }
    }

    var ans []int
    if diffCnt == 0 {
        ans = append(ans, 0)
    }

    for i := 1; i+plen-1 < slen; i++ {
        // i 是左边界, 则右边界是 i+plen-1
        oldC, newC := s[i-1]-'a', s[i+plen-1]-'a'
        // 若 diff[oldC]--  diff[newC]++一起执行，则当oldC==newC时，会重复计算diffCnt。
        // 或者直接oldC==newC额外处理。

        // 1. 先处理移出窗口的旧字符
        diff[oldC]--
        if diff[oldC] == 0 { // 滑动后左边去掉的字符对应数量变为一致。
            diffCnt--
        } else if diff[oldC] == -1 { // 变为不一致。
            diffCnt++
        }
        // 2. 再处理进入窗口的新字符
        diff[newC]++
        if diff[newC] == 0 { // 滑动后右边增加的字符对应数量变为一致。
            diffCnt--
        } else if diff[newC] == 1 { // 变为不一致。
            diffCnt++
        }

        if diffCnt == 0 {
            ans = append(ans, i)
        }
    }
    return ans
}
