package leetcode

/*
Leetcode No 64
*/
import "fmt"

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[m-1][n-1] = grid[m-1][n-1]
	for i := n - 2; i >= 0; i-- {
		dp[m-1][i] = grid[m-1][i] + dp[m-1][i+1]
	}
	for i := m - 2; i >= 0; i-- {
		dp[i][n-1] = grid[i][n-1] + dp[i+1][n-1]
	}
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			right := grid[i][j] + dp[i][j+1]
			down := grid[i][j] + dp[i+1][j]
			if right < down {
				dp[i][j] = right
			} else {
				dp[i][j] = down
			}
		}
	}
	fmt.Println(dp)
	return dp[0][0]
}

// 无需额外存储空间的版本
func minPathSum2() {

}

func Test64() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	ans := minPathSum(grid)
	fmt.Println(ans)
}
