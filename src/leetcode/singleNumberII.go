package leetcode

/*
LeetCode 137
*/
import "fmt"

func singleNumberII(nums []int) int {
	dict := make(map[int]int, len(nums)/3+1)
	for _, v := range nums {
		if i, ok := dict[v]; !ok {
			dict[v] = 1
		} else if i < 3 {
			dict[v]++
		}
	}
	for key, val := range dict {
		if val != 3 {
			return key
		}
	}
	return -1
}

func singleNumberII2(nums []int) int {
	var ans uint
	tab := make([]uint, 64)
	for j := 0; j < 64; j++ {
		for i := 0; i < len(nums); i++ {
			temp := uint(nums[i]) >> uint(j)
			tab[j] += temp & 1
		}
	}
	for i := 63; i >= 0; i-- {
		ans <<= uint(1)
		ans += tab[i] % uint(3)
	}
	return int(ans)
}

func Test137() {
	testExmps := [][]int{
		{2, 2, 3, 2},
		{-1, -1, -1, 2},
	}
	for _, v := range testExmps {
		ans := singleNumberII2(v)
		fmt.Println(v, ans)
	}
}
