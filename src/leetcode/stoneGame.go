package leetcode
/*
LeetCode No 877
Alex: Let's play Stone Game!
Lee : Fuck you!!!
*/

func max(a,b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
func min(a,b int) int {
	if a<= b {
		return a
	} else {
		return b
	}
}
func stoneGame(piles []int) bool {
	N := len(piles)
	dp := make([][]int,N+1)
	for i := 1; i < N+1; i++ {
		dp[i] = make([]int,N+1)
		dp[i][i] = -piles[i]
	}
	for size := 2; size <= N; size++ {
		for i := 1; i <= N-size+1; i++ {
			j := i+size-1
			parity := (size)%2
			if parity == 1 {
				dp[i][j] = min(-piles[i]+dp[i+1][j],-piles[j]+dp[i][j-1])
			} else {
				dp[i][j] = max(piles[i]+dp[i+1][j],piles[j]+dp[i][j-1])
			}
		}
	}
	return dp[1][N] > 0
}