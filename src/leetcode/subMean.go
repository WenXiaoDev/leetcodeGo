package leetcode

import (
	"fmt"
	"sort"
	"math"
)

func subtractMean(arr []int) ([]int,[]int) {
	sort.Ints(arr)
	sum,curr,N := 0,0,len(arr)
	var temp float64
	result := N-1
	for _,v := range arr {
		sum += v
	}
	for i := N-1; i >= 0; i-- {
		curr += arr[i]
		sub := float64(curr)/float64(N-1) - float64(sum-curr)/float64(i)
		if math.Abs(sub) > temp {
			result = i
		}
	}
	return arr[result:],arr[:result]
}

func TestUnknown() {
	fmt.Println("testing examples:")
	exp := []int{1,99,100}
	pileA,pileB := subtractMean(exp)
	fmt.Println(pileA)
	fmt.Println(pileB)
}
