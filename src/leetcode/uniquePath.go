package leetcode
/*
LeetCode No 62
Unique Paths
*/
func uniquePaths(m int, n int) int {
    dp := make([][]int,n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int,m)
		dp[i][m-1] = 1
	}
	for i := 0; i < m-1; i++ {
		dp[n-1][i] = 1
	}
	for i := n-2; i >= 0; i-- {
		for j  := m-2; j >= 0; j-- {
			dp[i][j] = dp[i+1][j]+dp[i][j+1]
		}
	}
	return dp[0][0]
}