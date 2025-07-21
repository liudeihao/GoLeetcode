package main

/*
回溯超时了。


// 这个是考虑可能有多张同一个票的
func findItinerary(tickets [][]string) []string {
	n := len(tickets)
	graph := make(map[string][]string)
	left := make(map[string]map[string]int)
	for _, t := range tickets {
		graph[t[0]] = append(graph[t[0]], t[1])
	}
	for k, strs := range graph {
		// 需要字典序，map无序（悲）
		slices.Sort(graph[k])
		left[k] = make(map[string]int)
		for i := 0; i < len(strs); i++ {
			left[k][strs[i]]++
		}
		graph[k] = slices.Compact(graph[k])
	}

	path := []string{"JFK"}

	var backtrack func(string) []string
	backtrack = func(src string) []string {
		if len(path) >= n+1 {
			return path
		}
		for _, dst := range graph[src] {
			if left[src][dst] == 0 {
				continue
			}
			left[src][dst]--
			path = append(path, dst)
			if backtrack(dst) != nil {
				return path
			}
			left[src][dst]++
			path = path[:len(path)-1]
		}
		return nil
	}

	return backtrack("JFK")
}

// 这个是每种票惟一的：
package main

import "slices"

func findItinerary(tickets [][]string) []string {
	n := len(tickets)
	graph := make(map[string][]string)
	used := make(map[string][]bool)
	for _, t := range tickets {
		graph[t[0]] = append(graph[t[0]], t[1])
	}
	for k := range graph {
		slices.Sort(graph[k])
		used[k] = make([]bool, len(graph[k]))
	}
	path := []string{"JFK"}

	var backtrack func(string) []string
	backtrack = func(src string) []string {
		if len(path) == n+1 {
			return path
		}
		for i, dst := range graph[src] {
			if used[src][i] {
				continue
			}
			used[src][i] = true
			path = append(path, dst)
			if backtrack(dst) != nil {
				return path
			}
			used[src][i] = false
			path = path[:len(path)-1]
		}
		return nil
	}

	return backtrack("JFK")
}

*/
