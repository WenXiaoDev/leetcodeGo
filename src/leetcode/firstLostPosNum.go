package leetcode

import (
	"fmt"
	"math"
)

/*
LeetCode No 41
方法1：
一个重要认知：在本题中，长度为n的数组里面，缺失的最小正整数x，其最小值为1，最大值为n+1。例如：
		12345678 -> 9
		12356 -> 4
		若x=n+2，意味着1到n+1都在数组中，这与长度为n矛盾

方法2：桶排序
基于方法1中的认知，输入的nums数组中，如果元素的值大于等于1而小于等于n，那么在nums中有唯一的
位置，使得下标加1的值等于它。所以，只要把属于区间[1,n]的数字换到它们对应的位置上面，那么其它
位置的下标加1的值就是缺失的那些数字的值，它们中的最小值就是所求的值。
*/
func firstMissingPositive2(nums []int) int {
	index, size := 0, len(nums)
	for index < size {
		temp := nums[index]
		if temp > 0 && temp <= size && temp != index+1 && nums[index] != nums[temp-1] {
			nums[index], nums[temp-1] = nums[temp-1], nums[index]
		} else {
			index++
		}
	}
	var ans int
	for ans = 0; ans < size; ans++ {
		if nums[ans] != ans+1 {
			break
		}
	}
	return ans + 1
}

func firstMissingPositive1(nums []int) int {
	size, ans := len(nums), 1
	exist_1 := false
	for i := 0; i < size; i++ {
		if nums[i] == 1 {
			exist_1 = true
		} else if nums[i] <= 0 || nums[i] > size {
			nums[i] = 1
		}
	}
	if !exist_1 {
		return 1
	}
	for i := 0; i < size; i++ {
		temp := math.Abs(float64(nums[i])) - 1.0
		if nums[int(temp)] > 0 {
			nums[int(temp)] *= -1
		}
	}
	for i, v := range nums {
		if v > 0 {
			ans = i + 1
			break
		} else if i == size-1 {
			ans = size + 1
		}
	}
	return ans
}

func Test41() {
	arrs := [][]int{
		{3, 4, -1, 1},
		{7, 8, 9, 11, 12, 1},
		{2, 1},
		{1, 1},
	}
	for i, v := range arrs {
		ans := firstMissingPositive2(v)
		fmt.Printf("arr%v:%v ans:%v\n", i, v, ans)
	}
}
