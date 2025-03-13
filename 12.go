package main

var sm = map[int]string{
	4:   "IV", //4
	9:   "IX", //9
	40:  "XL", // 40
	90:  "XC", // 90
	400: "CD", // 400
	900: "CM", // 900
}

var cs = []string{"I", "V", "X", "L", "C", "D", "M"}
var ns = []int{1, 5, 10, 50, 100, 500, 1000}

func getBiggest(num int) (int, string) {
	if num == 0 {
		return 0, ""
	}
	for i := range cs {
		if num < ns[i] {
			return ns[i-1], cs[i-1]
		}
	}
	return 1000, "M"
}

func intToRoman(num int) string {
	if num == 0 {
		return ""
	}
	cnt, x := 1, num
	for x > 10 {
		x = x / 10
		cnt *= 10
	}
	if x == 4 || x == 9 {
		return sm[x*cnt] + intToRoman(num-x*cnt)
	} else {
		big, roman := getBiggest(num)
		return roman + intToRoman(num-big)
	}
}

/* 如果只是要解决这个问题，用这个吧。。
var (
    thousands = []string{"", "M", "MM", "MMM"}
    hundreds  = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
    tens      = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
    ones      = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
)

func intToRoman(num int) string {
    return thousands[num/1000] + hundreds[num%1000/100] + tens[num%100/10] + ones[num%10]
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/integer-to-roman/solutions/774611/zheng-shu-zhuan-luo-ma-shu-zi-by-leetcod-75rs/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
