package leetcode

/*
LeetCode No 268

1.桶排序
2.异或运算
*/
import "fmt"

func missingNumber(nums []int) int {
	index, size := 0, len(nums)
	for index < size {
		temp := nums[index]
		if temp < size && temp != index {
			nums[index], nums[temp] = nums[temp], nums[index]
		} else {
			index++
		}
	}
	var ans int
	for ans = 0; ans < size; ans++ {
		if nums[ans] != ans {
			break
		}
	}
	return ans
}

// method 2: 异或运算
// 1.若n在nums中，则有一个属于[0,n-1]的数被排除在外，
// n 在一个循环中，每次与下标和对应元素的异或值进行异或运算
// 一个数字在进行两次相同的异或运算之后，仍然是它本身，如 a^b^b = a
func missingNumber2(nums []int) int {
	missing := len(nums)
	for i := 0; i < len(nums); i++ {
		missing ^= i ^ nums[i]
	}
	return missing
}

func Test268() {
	testExmp := [][]int{
		{9, 6, 4, 2, 3, 5, 7, 0, 1},
		{3, 0, 1},
	}
	for i, v := range testExmp {
		ans := missingNumber2(v)
		fmt.Println(i, v, ans)
	}
}
