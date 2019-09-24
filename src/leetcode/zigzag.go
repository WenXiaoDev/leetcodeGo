package leetcode
/*
LeetCode No 6
		                    L     D     R
	LEETCODEISHIRING ---->	E   O E   I I
		                    E C   I H   N
		                    T     S     G
*/

import "fmt"

func convert(s string, numRows int) string {
	size := len(s)
	if numRows == 1 || numRows == size {
		return s
	}
	if numRows == 2 {
		var left,right string
		for i,j := 0,1; i < size; i,j = i+2,j+2 {
			left += string(s[i])
			if j < size {
				right += string(s[j]) 
			}
		}
		return left+right
	}
	itvl := 2*numRows-2
	result := []uint8{}
	for i := 0; i < size; i+=itvl {
		result = append(result,s[i])
	}
	for i := 1; i < numRows-1; i++ {
		for start := i; start < size; start+=itvl {
			result = append(result,s[start])
			if start+itvl-2*i < size {
				result = append(result,s[start+itvl-2*i])
			}
		}
	}
	for i := numRows-1; i < size; i+=itvl {
		result = append(result,s[i])
	}
	return string(result)
}

func Test6() {
	str := "LEETCODEISHIRING"
	result := convert(str,4)
	fmt.Println(result)
}