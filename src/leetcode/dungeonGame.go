package leetcode

/*
LeetCode No 174
*/
import "fmt"

func lowwer(val int) int {
	// 注意，骑士的血量在任何时候至少为 1
	if val < 1 {
		return 1
	}
	return val
}
func minium(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 1
	}
	m, n := len(dungeon), len(dungeon[0])
	dp := make([][]int, len(dungeon))
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}
	// hp + delta = nextHp
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j != n-1 {
				dp[i][j] = lowwer(dp[i][j+1] - dungeon[i][j])
			} else if i != m-1 && j == n-1 {
				dp[i][j] = lowwer(dp[i+1][j] - dungeon[i][j])
			} else if i == m-1 && j == n-1 {
				dp[i][j] = lowwer(1 - dungeon[i][j])
			} else {
				hpRight := lowwer(dp[i][j+1] - dungeon[i][j])
				hpDown := lowwer(dp[i+1][j] - dungeon[i][j])
				dp[i][j] = minium(hpRight, hpDown)
			}
		}
	}
	return dp[0][0]
}

func Test174() {
	grid := [][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}
	ans := calculateMinimumHP(grid)
	fmt.Println(ans)
}
