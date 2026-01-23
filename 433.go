package main

type Queue_433 struct {
	items []string
}

func (q *Queue_433) Push(x string) {
	q.items = append(q.items, x)
}

func (q *Queue_433) Pop() string {
	ret := q.items[0]
	q.items = q.items[1:]
	return ret
}

func (q *Queue_433) Empty() bool {
	return len(q.items) == 0
}

func (q *Queue_433) Len() int {
	return len(q.items)
}

func diffOneChar(s1, s2 string) bool {
	cnt := 0
	for i := range s1 {
		if s1[i] != s2[i] {
			cnt++
		}
	}
	return cnt == 1
}

func minMutation(startGene string, endGene string, bank []string) int {
	var si = map[string]int{}
	si[startGene] = 0
	q := Queue_433{}
	q.Push(startGene)
	depth := 0
	for !q.Empty() {
		for n := q.Len(); n > 0; {
			front := q.Pop()
			for _, s := range bank {
				canTransform := diffOneChar(front, s)
				if s == endGene && canTransform {
					return depth + 1
				}
				if _, ok := si[s]; !ok && canTransform {
					si[s] = depth + 1
					q.Push(s)
				}
			}
			n--
		}
		depth++
	}
	d, ok := si[endGene]
	if !ok {
		return -1
	}
	return d
}
