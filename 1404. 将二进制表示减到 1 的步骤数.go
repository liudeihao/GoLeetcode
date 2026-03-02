package main

func numSteps(s string) int {
    ans := 0
    bs := []byte(s)
    i := len(s) - 1
    for i > 0 {
        if bs[i] == '0' {
            i--
            ans++
        } else {
            break
        }
    }
    for i > 0 {
        k := i
        for ; k >= 0; k-- {
            if bs[k] == '0' {
                bs[k] = '1'
                ans += i - k + 1
                i = k
            }
        }
        if k == -1 {
            ans += i - k + 1
            return ans
        }
    }
    if bs[0] == '0' {
        ans++
    }
    return ans
}
