package leetcode
/*
LeetCode No 63
Unique Paths with obstacles in the grid
*/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    n,m := len(obstacleGrid),len(obstacleGrid[0])
	dp := make([][]int,n)
	for i := 0; i < n; i++{
		dp[i] = make([]int,m)
	}
    if obstacleGrid[n-1][m-1] == 1 {
        dp[n-1][m-1] = 0
    } else {
        dp[n-1][m-1] = 1
    }
	for i := n-1; i >= 0; i-- {
		for j := m-1; j >= 0; j-- {
			var down,right int
			if obstacleGrid[i][j] == 1 {
				continue
			}
			if i == n-1 || obstacleGrid[i+1][j] == 1 {
				down = 0
			} else {
				down = dp[i+1][j]
			}
			if j == m-1 || obstacleGrid[i][j+1] == 1 {
				right = 0
			} else {
				right = dp[i][j+1]
			}
			dp[i][j] = right+down+dp[i][j]
		}
	}
	return dp[0][0]
}