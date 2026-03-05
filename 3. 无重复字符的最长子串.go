package main

func lengthOfLongestSubstring(s string) int {
    n := len(s)
    mp := make(map[byte]bool)
    l, r := 0, 0
    ans := 0
    for r < n {
        // [l, r) 中没有重复的字符
        for r < n && !mp[s[r]] {
            mp[s[r]] = true
            r++
        }
        ans = max(ans, r-l) // len( s[l, r) ) == r-l
        if r == n {
            break
        }
        for l < r && mp[s[r]] {
            mp[s[l]] = false
            l++
        }
    }
    return ans
}
