package leetcode
/*

思路(time complexity = n)：
	从头开始遍历数组，每到达一个新位置，检查之前的子段和是否小于0
*/
func maxSubArray(nums []int) int {
	result,curr := nums[0],0
	for i := 0; i < len(nums); i++ {
		if curr < 0 {
			curr = nums[i]
		} else {
			curr += nums[i]
		}
		if result < curr {
			result = curr
		}
	}
	return result
}