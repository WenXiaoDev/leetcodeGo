package leetcode

/*
LeetCode No. 10	Hard
realize the match of reguluar expression with simbol '*' and '.'
直接看题解了
	动态规划法是基于在回溯法的基础上，把需要重复求解的子问题先求解并记录下来。
	回溯法的过程是：
	从头开始依次对比各个字符，看是否匹配，并观察其下一个字符是否是*，如果是，就
	需要考虑*位置对应的原字符串的字符是否匹配，或者考虑*修饰的字符为0个，而*之后的
	pattern后缀与当前子串是否匹配。例子如下：
e.g s: missisippi
	p: mis*is*p*.
	因为有*存在，它修饰前面1个字符，如第一个*出现时，s与s匹配，接下来只要判断sisippi与
	s*is*p*. 是否匹配，或者ssisippi与is*p*.是否匹配，有一个符合即可。
*/
import "fmt"

//回溯法 recall
func isMatch1(s, p string) bool {
	if p == "" {
		return s == ""
	}
	firstMatch := s != "" && (s[0] == p[0] || p[0] == '.')
	if len(p) >= 2 && p[1] == '*' {
		return isMatch1(s, p[2:]) || (firstMatch && isMatch1(s[1:], p))
	} else {
		return firstMatch && isMatch1(s[1:], p[1:])
	}
}

//动态规划 dynamic programming
func isMatch(s, p string) bool {
	tab := make([][]bool, len(s)+1)
	for i := 0; i < len(p)+1; i++ {
		tab[i] = make([]bool, len(p)+1)
	}
	tab[len(s)][len(p)] = true
	for i := len(s); i >= 0; i-- {
		for j := len(p) - 1; j >= 0; j-- {
			firstMatch := s[i:] != "" && (s[i] == p[j] || p[j] == '.')
			if len(p)-j >= 2 && p[j+1] == '*' {
				tab[i][j] = tab[i][j+2] || (firstMatch && tab[i+1][j])
			} else {
				tab[i][j] = firstMatch && tab[i+1][j+1]
			}
		}
	}
	return tab[0][0]
}

func Test10() {
	fmt.Println(isMatch1("sa", "s*"))
	fmt.Println(isMatch("missisippi", "mis*is*p*."))
}
