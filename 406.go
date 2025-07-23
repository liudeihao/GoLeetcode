package main

import (
    "container/list"
    "slices"
)

func reconstructQueue(people [][]int) [][]int {
    // 优先排队的人是：h大、k小的！
    slices.SortFunc(people, func(a, b []int) int {
        if a[0] == b[0] {
            return a[1] - b[1]
        }
        return b[0] - a[0]
    })

    ls := list.New()

    for _, person := range people {
        pos := person[1] // k就是要插入的位置...
        mark := ls.PushBack(person)
        e := ls.Front()
        for pos > 0 {
            pos--
            e = e.Next()
        }
        ls.MoveBefore(mark, e)
    }
    var ans [][]int
    for e := ls.Front(); e != nil; e = e.Next() {
        ans = append(ans, e.Value.([]int))
    }
    return ans
}
