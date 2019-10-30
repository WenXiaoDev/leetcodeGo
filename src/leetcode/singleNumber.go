package leetcode

/*
LeetCode No 136
*/

//直接用异或运算

func singleNumber(nums []int) int {
	a := 0
	for _, v := range nums {
		a ^= v
	}
	return a
}
