package leetcode

func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}

	numIslands := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				numIslands++
				dfs(grid, i, j)
			}
		}
	}

	return numIslands
}

func dfs(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != '1' {
		return
	}

	grid[i][j] = '0'  // mark as visited
	dfs(grid, i+1, j) // down
	dfs(grid, i-1, j) // up
	dfs(grid, i, j+1) // right
	dfs(grid, i, j-1) // left
}
