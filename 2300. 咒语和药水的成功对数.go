package main

import (
    "slices"
    "sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
    ans := make([]int, len(spells))
    slices.Sort(potions)
    for i, s := range spells {
        l := sort.Search(len(potions), func(j int) bool {
            return s*potions[j] >= int(success)
        })
        ans[i] = len(potions) - l
    }
    return ans
}
