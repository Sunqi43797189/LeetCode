package leetcode

import (
	"strings"
)

// https://leetcode-cn.com/problems/tic-tac-toe-lcci/

// o 先下棋
// 1 o 赢了 o-x = 1
// 2 x 赢了 o-x = 0
// 3 没分胜负 x-o = 0 或 o-x = 1
func tictactoe(board []string) string {
	n := len(board)
	xCount, oCount := 0, 0
	for _, value := range board {
		for _, key := range value {
			if string(key) == "X" {
				xCount++
			}
			if string(key) == "O" {
				oCount++
			}
		}
	}

	if oCount-xCount == 1 || oCount == xCount {
		if win(board, getFlag("X", n)) {
			return "X"
		}
		if win(board, getFlag("O", n)) {
			return "O"
		}
		if oCount+xCount == n*n {
			return "Draw"
		}
	}

	return "Pending"
}

func win(board []string, flag string) bool {
	n := len(board)
	// 横向N连
	for _, value := range board {
		if value == flag {
			return true
		}
	}
	// 纵向N连
	for i := 0; i < n; i++ {
		prefix := ""
		for _, value := range board {
			prefix += string(value[i])
		}
		if flag == prefix {
			return true
		}
	}
	// /向对角线
	{
		prefix := ""
		for index, value := range board {
			prefix += string(value[n-1-index])
		}
		if prefix == flag {
			return true
		}
	}

	// \向对角线
	{
		prefix := ""
		for index, value := range board {
			prefix += string(value[index])
		}
		if prefix == flag {
			return true
		}
	}
	return false
}

func getFlag(str string, size int) string {
	arr := make([]string, size)
	for index, _ := range arr {
		arr[index] = str
	}
	return strings.Join(arr, "")
}
