package data_structures

// 单调栈就是栈内元素单调递增或者单调递减的栈，单调栈只能在栈顶操作。
// 单调栈可以用来解决在一个数组中找出每个元素对应左右两部分比自己小的值并且最近的值的情况。

// 1. 滑动窗口问题：快速找到滑动窗口内的最大值。
// 2. 单调队列问题：维护一个单调递增或递减的序列。
// 3. 股票价格问题：找到股票价格的下一个更高或更低的价格。

type MonotonyStack struct {
	stack []int
	n     int
}

func (s *MonotonyStack) Top() (int, bool) {
	if s.n > 0 {
		return s.stack[s.n-1], true
	}
	return 0, false
}

func (s *MonotonyStack) Push(i int) {
	if s.n == 0 {
		s.stack = append(s.stack, i)
		s.n++
		return
	}
	for t, _ := s.Top(); t >= i && s.n > 0; {
		s.pop()
		t, _ = s.Top()
	}
	s.stack = append(s.stack, i)
	s.n++
	return
}

func (s *MonotonyStack) pop() {
	s.n--
	s.stack = s.stack[:s.n]
}

func (s *MonotonyStack) Clear() {
	s.n = 0
	s.stack = []int{}
}
