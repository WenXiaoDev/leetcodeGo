package leetcode
import "fmt"
/*
LeetCode No 292
Nim Game
*/
func canWinNim(n int) bool {
	if n == 0 {
		return false
	} else if n < 4 {
		return true
	}
	dp := make([]bool,n+1,n+1)
	dp[0],dp[1],dp[2],dp[3] = false,true,true,true
	for i := 4; i <= n; i++ {
		dp[i] = !(dp[i-1] && dp[i-2] && dp[i-3])
	}
	return dp[n]
}

func Test292() {
	for i := 0; i <= 30; i++ {
		fmt.Print(i,": ",canWinNim(i)," ")
	}
}