package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	var ans []int
	// 邻接表
	graph := make([][]int, numCourses)
	inDegree := make([]int, numCourses)
	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
		inDegree[p[0]]++
	}
	// 队列
	var q []int
	for i := range numCourses {
		if inDegree[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		front := q[0]
		q = q[1:]
		ans = append(ans, front)
		for _, v := range graph[front] {
			inDegree[v]--
			if inDegree[v] == 0 {
				q = append(q, v)
			}
		}
	}
	fmt.Println(ans)
	if len(ans) != numCourses {
		return nil
	}
	return ans
}
