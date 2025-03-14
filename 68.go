package main

func fullJustify(words []string, maxWidth int) []string {
	n := len(words)
	var (
		wordNums       = make([]int, n)
		extraSpaceNums = make([]int, n)
	)
	i, line := 0, 0
	for i < n {
		// 先塞一个
		lineLength := len(words[i])
		i++
		wordNums[line]++
		// 直到下一个词塞不下
		for i < n && lineLength+len(words[i])+1 <= maxWidth {
			lineLength += len(words[i]) + 1
			wordNums[line]++
			i++
		}
		extraSpaceNums[line] = maxWidth - lineLength
		line++
	}

	i = 0
	extraSpaceNums[line-1] = 0
	ans := make([]string, line)
	for l := range line {
		s := make([]byte, maxWidth)

		// 为了避免后面太麻烦，先都给你填上空
		for idx := range s {
			s[idx] = ' '
		}

		wordNum := wordNums[l]
		extraSpaceNum := extraSpaceNums[l]
		cursor := 0

		// 均匀分配额外多出来的空格
		extraSpacePerWord := 0
		splitNum := wordNum - 1
		for splitNum != 0 && extraSpaceNum-splitNum >= 0 {
			extraSpaceNum -= splitNum
			extraSpacePerWord++
		}

		for j := 0; j < wordNum; j++ {
			// 复制一个词
			for _, c := range words[i] {
				s[cursor] = byte(c)
				cursor++
			}
			i++
			// 添加空格
			nSpaces := 1 + extraSpacePerWord
			if extraSpaceNum > 0 && j < extraSpaceNum {
				nSpaces++
			}
			if j < splitNum {
				for _ = range nSpaces {
					s[cursor] = ' '
					cursor++
				}
			}
		}

		ans[l] = string(s)
	}
	return ans
}
