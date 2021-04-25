package main

// LeetCode 200 https://leetcode-cn.com/problems/number-of-islands/

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	res := 0
	var boolArr = make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		boolArr[i] = make([]bool, len(grid[i]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' && !boolArr[i][j] {
				DfsNumIslands(grid, boolArr, i, j)
				res++
			}
		}
	}
	return res
}

func DfsNumIslands(grid [][]byte, boolArr [][]bool, i, j int) {
	for _, near := range getNear(i,j, len(grid), len(grid[0])) {
		if grid[near[0]][near[1]] == '1' && !boolArr[near[0]][near[1]] {
			boolArr[near[0]][near[1]] = true
			DfsNumIslands(grid, boolArr, near[0], near[1])
		}
	}
}

func getNear(i, j, x, y int) [][]int {
	var nears [][]int
	if i-1 > 0 {
		nears = append(nears, []int{i - 1, j})
	}

	if j+1 < y {
		nears = append(nears, []int{i, j + 1})
	}

	if i+1 < x {
		nears = append(nears, []int{i + 1, j})
	}

	if j-1 > 0 {
		nears = append(nears, []int{i, j - 1})
	}
	return nears
}
