package leetcode

/*
LeetCode 260
只出现一次的数字III：
数组中恰好有两个数字只出现1次，其它数字均出现了2次，找出这两个数组。
使用时间复杂度为O(n)的方法，空间复杂度尽量为O(1)

1.位运算 稍慢 space:O(1)
    xor&(-xor) 得到xor中最低的非0位。
     5： 0000,0101
    -5： 1111,1011 = 1111，1010 + 1
	1）两个数字a和b按位异或的值（a^b）的二进制表示中，若第i位的值为1，则说明这a和b在这
	   一位上一个为0一个为1；找到异或值xor最低的非0位。
	2）只需要根据异或值的这一位，将数组进行分组，那么a和b必然被分到不同的组，而其他
	   有重复的数字也会因此被分到不同的组，且相同的数字会被分到同一组。
	   为了得到能将数组进行分组的那一位，我们可以将那一位保存在数字mask中，在mask中，只
	   有那一位上面的值为1，其它都为0；
	3）将两个分组进行连续的异或运算，最后得到的两个值就是要求的两个值。也就是把原问题拆
	   分为两个“singleNumber I”问题。
2.哈希 稍快 space:0=O(n)
	用map存储出现的数字，若出现2次就删掉。
	最终map中只剩下要求的数字。
*/

func singleNumberIII(nums []int) []int {
	// first: calculate xor
	xor := 0
	ans := make([]int, 2)
	for _, v := range nums {
		xor ^= v
	}
	// second: calculate mask
	mask := xor & (-xor)
	// third: solve two sub-problem
	for _, v := range nums {
		if v&mask == 0 {
			ans[0] ^= v
		} else {
			ans[1] ^= v
		}
	}
	return ans
}

func singleNumberIII2(nums []int) []int {
	dict := make(map[int]int)
	ans := make([]int, 0)
	for _, v := range nums {
		if _, ok := dict[v]; !ok {
			dict[v] = 1
		} else {
			delete(dict, v)
		}
	}
	for k, _ := range dict {
		ans = append(ans, k)
	}
	return ans
}
