package main

func minOperations_1758(s string) int {
    one0 := 0
    zero0 := 0
    for i, c := range s {
        if (i&1 == 0 && c != '1') || (i&1 == 1 && c != '0') {
            one0++
        }
        if (i&1 == 0 && c != '0') || (i&1 == 1 && c != '1') {
            zero0++
        }
    }
    return min(one0, zero0)
}
