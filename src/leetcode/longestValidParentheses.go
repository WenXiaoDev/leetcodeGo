package leetcode
/*
LeetCode No 32

*/
//1. 
func longestValidParentheses(s string) int {
	tab := make([]int,len(s))
	var maxLen int
	for i := 1; i < len(s); i++ {
		if s[i] == '(' {
			continue
		}
		if s[i-1] == '(' {
			if i > 1 {
				tab[i] = tab[i-2]+2
			} else {
				tab[i] = 2
			}
		}
		if i-tab[i-1] > 0 && s[i-tab[i-1]-1] == '(' {
			if i-tab[i-1]-2 >= 0 {
				tab[i] = tab[i-1]+tab[i-tab[i-1]-2]+2
			} else {
				tab[i] = tab[i-1]+2
			}
		}
		if tab[i] > maxLen {
			maxLen = tab[i]
		}
	}
	return maxLen
}

// 2. 用栈来保存各个括号字符的下标。并计算长度，并记录 时间复杂度 n
func stackSolution(s string) int {
	stack := []int{-1}
	maxLen := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack,i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack,i)
			} else if top := stack[len(stack)-1]; i-top > maxLen {
				maxLen = i-top
			}
		}
	}
	return maxLen
}

