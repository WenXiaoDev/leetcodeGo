package leetcode
/*
LeetCode No 13
输入: "MCMXCIV"
输出: 1994
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
I 可以放在 V (5)   和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50)  和 C (100) 的左边，来表示 40 和 90。 
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
*/
import "fmt"

var dict = map[uint8]int {
	73:1,86:5,88:10,76:50,67:100,68:500,77:1000,
}
func romaToInt(s string) int {
	total := 0
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			total += dict[s[i]]
			continue
		}
		if dict[s[i]] >= dict[s[i+1]] {
			total += dict[s[i]]
			continue
		}
		var current,j int
		for j = i; j < len(s); j++ {
			if dict[s[j]] < dict[s[j+1]] && j+1 < len(s) {
				current += dict[s[j]]
				continue
			}
		}
		total = total+dict[s[j]]-current
		i = j
	}
	return total
}
func romaToInt2(s string) int {
	total := 0
	for i := 0; i < len(s)-1; i++ {
		if dict[s[i]] >= dict[s[i+1]] {
			total += dict[s[i]]
			continue
		}
		total -= dict[s[i]]
	}
	return total+dict[s[len(s)-1]]
}

func Test13() {
	str := "MCMXCIV"
	num := romaToInt2(str)
	fmt.Println(num)
	for i := 0; i < len(str); i++ {
		fmt.Printf("unit8:%v char:%v\n",str[i],string(str[i]))
	}
}