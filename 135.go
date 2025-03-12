package main

func candy(ratings []int) int {
	n := len(ratings)
	asc := make([]int, n)
	asc[0] = 1
	for cnt, i := 1, 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			cnt++
		} else {
			cnt = 1
		}
		asc[i] = cnt
	}

	desc := make([]int, n)
	desc[n-1] = 1
	for cnt, i := 1, n-2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			cnt++
		} else {
			cnt = 1
		}
		desc[i] = cnt
	}

	result := 0
	for i := range asc {
		result += max(asc[i], desc[i])
	}
	return result
}
