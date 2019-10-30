package leetcode
/*
LeetCode No 980
1 ：起点 2.终点 0：可以经过 -1：障碍
*/

type node struct {
	x,y int
}

func uniquePathIII(grid [][]int) int {
	rows,cols := len(grid),len(grid[0])
	var startX,startY int
	var ans,pathNum int
	dr := []int{0,1,0,-1}
	dc := []int{1,0,-1,0}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 {
				pathNum++
			} else if grid[i][j] == 1 {
				startX,startY = i,j
			}
		}
	}
	// dfs recursive
	var dfs func(r,c int)
	dfs = func(r,c int) {
		pathNum--
		if pathNum < 0 {
			return
		}
		if grid[r][c] == 2 {
			if pathNum == 0 {
				ans++
			}
			return
		}
		grid[r][c] = 3
		for i := 0; i < 4; i++ {
			nr,nc := r+dr[i],c+dc[i]
			if nc >=0 && nc < cols && nr >=0 && nr < rows {
				if grid[nr][nc]%2 == 0 {
					dfs(nr,nc)
				}
			}
		}
		grid[r][c] = 0
	}
	dfs(startX,startY)
	return ans
}