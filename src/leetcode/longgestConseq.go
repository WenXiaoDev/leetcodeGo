package leetcode
/*
LeetCode No 128
*/
import "fmt"


// use union-find set(并查集) slower
func longestConsecutive(nums []int) int {
	acc := make([]int,len(nums))
	set := make(map[int]int)
	itered := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		acc[i] = i
	}
	for i := 0;i < len(nums); i++ {
		current := nums[i]
		if _,ok := set[current]; !ok {
			set[current] = i
		} else {
			continue
		}
		if index,ok := set[current-1]; ok {
			acc[i] = index
		}
		if index,ok := set[current+1]; ok {
			acc[index] = i
		}
	}
	var maxLen int
	for i := 0; i < len(acc); i++ {
		if _,ok := itered[nums[i]]; acc[i] != i || ok {
			continue
		}
		itered[nums[i]] = true
		len := 1
		for {
			if _,ok := set[nums[i]+len]; ok {
				len++
			} else {
				break
			}
		}
		if len > maxLen {
			maxLen = len
		}
	}
	return maxLen
}

//use Hash Table -faster
func longestConsecutive2(nums []int) int {
    maxLen := 0
    set := make(map[int]bool)
    for _,v := range nums {
        if _,ok := set[v]; !ok {
            set[v] = true
        }
    }
    for val,_ := range set {
        if _,ok := set[val-1]; !ok {
            len := 1
            for {
                if _,exist := set[val+len]; !exist {
                    break
                }
                len++
            }
            if len > maxLen {
                maxLen = len
            }
        }
    }
    return maxLen
}

func Test128() {
	array := []int{100,4,200,1,3,2,5,7,6,20,9,8,10}
	acc := longestConsecutive(array)
	fmt.Println(len(array))
	fmt.Println(acc)
}