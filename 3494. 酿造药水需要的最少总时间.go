package main

func minTime(skill []int, mana []int) int64 {
    // 巫师可以等药水，但药水绝对不能等巫师。
    n, m := len(skill), len(mana)
    endTime := make([]int, n)
    for j := 0; j < m; j++ {
        var curTime int // 这瓶药水最快只能在 curTime 这个时间点彻底酿造完。
        // 正向循环是为了妥协最慢的人（找延迟）
        for i := 0; i < n; i++ {
            curTime = max(curTime, endTime[i]) // 一旦卡住，就把整瓶药水的“排期”往后挪。
            curTime += skill[i] * mana[j]
        }
        endTime[n-1] = curTime
        // 逆向循环是为了保证药水的连贯性（定时间）
        // 既然最终出厂时间确定了，而且药水的流转是“无缝衔接”不可压缩的，我们就必须把前面 n-1 个巫师的 endTime 强行更新。
        // 为什么要倒推？ 因为如果你只正向算，有些巫师的 endTime 可能是假的！
        // 比如巫师 0 很快就做完了，但巫师 1 堵车了导致整瓶药水推迟发车。此时巫师 0 真实的完工时间，必须是巫师 1 开始的时间。
        for i := n - 2; i >= 0; i-- {
            endTime[i] = endTime[i+1] - skill[i+1]*mana[j]
        }
    }
    return int64(endTime[n-1])
}
