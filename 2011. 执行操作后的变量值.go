package main

func finalValueAfterOperations(operations []string) int {
    mp := map[string]int{
        "++X": 1,
        "X++": 1,
        "--X": -1,
        "X--": -1,
    }
    ans := 0
    for _, op := range operations {
        ans += mp[op]
    }
    return ans
}
